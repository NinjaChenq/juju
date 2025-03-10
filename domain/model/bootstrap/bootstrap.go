// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package bootstrap

import (
	"context"
	"fmt"

	"github.com/canonical/sqlair"
	"github.com/juju/version/v2"

	"github.com/juju/juju/core/constraints"
	"github.com/juju/juju/core/database"
	coremodel "github.com/juju/juju/core/model"
	jujuversion "github.com/juju/juju/core/version"
	"github.com/juju/juju/domain/model"
	modelerrors "github.com/juju/juju/domain/model/errors"
	"github.com/juju/juju/domain/model/service"
	"github.com/juju/juju/domain/model/state"
	secretbackenderrors "github.com/juju/juju/domain/secretbackend/errors"
	internaldatabase "github.com/juju/juju/internal/database"
	jujusecrets "github.com/juju/juju/internal/secrets/provider/juju"
	kubernetessecrets "github.com/juju/juju/internal/secrets/provider/kubernetes"
	"github.com/juju/juju/internal/uuid"
)

type modelTypeStateFunc func(context.Context, string) (string, error)

func (m modelTypeStateFunc) CloudType(c context.Context, n string) (string, error) {
	return m(c, n)
}

// CreateGlobalModelRecord is responsible for making a new model with all of its
// associated metadata during the bootstrap process.
// If the GlobalModelCreationArgs does not have a credential name set then no
// cloud credential will be associated with the model.
//
// The only supported agent version during bootstrap is that of the current
// controller. This will be the default if no agent version is supplied.
//
// The following error types can be expected to be returned:
// - modelerrors.AlreadyExists: When the model uuid is already in use or a model
// with the same name and owner already exists.
// - errors.NotFound: When the cloud, cloud region, or credential do not exist.
// - errors.NotValid: When the model uuid is not valid.
// - [modelerrors.AgentVersionNotSupported]
// - [usererrors.NotFound] When the model owner does not exist.
// - [secretbackenderrors.NotFound] When the secret backend for the model
// cannot be found.
//
// CreateGlobalModelRecord expects the caller to generate their own model id and
// pass it to this function. In an ideal world we want to have this stopped and
// make this function generate a new id and return the value. This can only be
// achieved once we have the Juju client stop generating id's for controller
// models in the bootstrap process.
func CreateGlobalModelRecord(
	modelID coremodel.UUID,
	args model.GlobalModelCreationArgs,
) internaldatabase.BootstrapOpt {
	return func(ctx context.Context, controller, model database.TxnRunner) error {
		if err := args.Validate(); err != nil {
			return fmt.Errorf("cannot create model when validating args: %w", err)
		}

		if err := modelID.Validate(); err != nil {
			return fmt.Errorf(
				"cannot create model %q when validating id: %w", args.Name, err,
			)
		}

		agentVersion := args.AgentVersion
		if args.AgentVersion == version.Zero {
			agentVersion = jujuversion.Current
		}

		if agentVersion.Major != jujuversion.Current.Major || agentVersion.Minor != jujuversion.Current.Minor {
			return fmt.Errorf("%w %q during bootstrap", modelerrors.AgentVersionNotSupported, agentVersion)
		}
		args.AgentVersion = agentVersion

		activator := state.GetActivator()
		return controller.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
			modelTypeState := modelTypeStateFunc(
				func(ctx context.Context, cloudName string) (string, error) {
					return state.CloudType()(ctx, preparer{}, tx, cloudName)
				})
			modelType, err := service.ModelTypeForCloud(ctx, modelTypeState, args.Cloud)
			if err != nil {
				return fmt.Errorf("determining cloud type for model %q: %w", args.Name, err)
			}

			if args.SecretBackend == "" && modelType == coremodel.CAAS {
				args.SecretBackend = kubernetessecrets.BackendName
			} else if args.SecretBackend == "" && modelType == coremodel.IAAS {
				args.SecretBackend = jujusecrets.BackendName
			} else if args.SecretBackend == "" {
				return fmt.Errorf(
					"%w for model type %q when creating model with name %q",
					secretbackenderrors.NotFound,
					modelType,
					args.Name,
				)
			}

			if err := state.Create(ctx, preparer{}, tx, modelID, modelType, args); err != nil {
				return fmt.Errorf("create bootstrap model %q with uuid %q: %w", args.Name, modelID, err)
			}

			if err := activator(ctx, preparer{}, tx, modelID); err != nil {
				return fmt.Errorf("activating bootstrap model %q with uuid %q: %w", args.Name, modelID, err)
			}
			return nil
		})
	}
}

// CreateReadOnlyModel creates a new model within the model database with all of
// its associated metadata. The data will be read-only and cannot be modified
// once created.
func CreateReadOnlyModel(
	id coremodel.UUID,
	controllerUUID uuid.UUID,
) internaldatabase.BootstrapOpt {
	return func(ctx context.Context, controllerDB, modelDB database.TxnRunner) error {
		if err := id.Validate(); err != nil {
			return fmt.Errorf("creating read only model, id %q: %w", id, err)
		}

		var m coremodel.Model
		err := controllerDB.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
			var err error
			m, err = state.GetModel(ctx, tx, id)
			return err
		})
		if err != nil {
			return fmt.Errorf("getting model for id %q: %w", id, err)
		}

		args := model.ModelDetailArgs{
			UUID:              m.UUID,
			AgentVersion:      m.AgentVersion,
			ControllerUUID:    controllerUUID,
			Name:              m.Name,
			Type:              m.ModelType,
			Cloud:             m.Cloud,
			CloudRegion:       m.CloudRegion,
			CredentialOwner:   m.Credential.Owner,
			CredentialName:    m.Credential.Name,
			IsControllerModel: true,
		}

		return modelDB.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
			return state.CreateReadOnlyModel(ctx, args, preparer{}, tx)
		})
	}
}

// SetModelConstraints sets the constraints for the controller model during bootstrap.
// The following error types can be expected:
// - [networkerrors.SpaceNotFound]: when a space constraint is set but the
// space does not exist.
// - [machineerrors.InvalidContainerType]: when the container type set on the
// constraints is invalid.
// - [modelerrors.NotFound]: when no model exists to set constraints for.
func SetModelConstraints(constraints constraints.Value) internaldatabase.BootstrapOpt {
	return func(ctx context.Context, controller, modelDB database.TxnRunner) error {
		return modelDB.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
			modelCons := model.FromCoreConstraints(constraints)
			return state.SetModelConstraints(ctx, preparer{}, tx, modelCons)
		})
	}
}

type preparer struct{}

func (p preparer) Prepare(query string, typeSamples ...any) (*sqlair.Statement, error) {
	return sqlair.Prepare(query, typeSamples...)
}
