// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package provider

import (
	"context"
	"net/url"
	osexec "os/exec"

	jujuclock "github.com/juju/clock"
	"github.com/juju/errors"
	"github.com/juju/jsonschema"
	"github.com/juju/utils/v4/exec"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	k8slabels "k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/juju/juju/caas"
	k8s "github.com/juju/juju/caas/kubernetes"
	k8scloud "github.com/juju/juju/caas/kubernetes/cloud"
	"github.com/juju/juju/caas/kubernetes/provider/utils"
	k8swatcher "github.com/juju/juju/caas/kubernetes/provider/watcher"
	"github.com/juju/juju/cloud"
	"github.com/juju/juju/environs"
	environsbootstrap "github.com/juju/juju/environs/bootstrap"
	environscloudspec "github.com/juju/juju/environs/cloudspec"
	"github.com/juju/juju/environs/envcontext"
)

// ClusterMetadataStorageChecker provides functionalities for checking k8s cluster storage and pods details.
type ClusterMetadataStorageChecker interface {
	k8s.ClusterMetadataChecker
	ListStorageClasses(ctx context.Context, selector k8slabels.Selector) ([]storagev1.StorageClass, error)
	ListPods(ctx context.Context, namespace string, selector k8slabels.Selector) ([]corev1.Pod, error)
}

type kubernetesEnvironProvider struct {
	environProviderCredentials
	cmdRunner          CommandRunner
	builtinCloudGetter func(CommandRunner) (cloud.Cloud, error)
	brokerGetter       func(context.Context, environs.OpenParams, environs.CredentialInvalidator) (ClusterMetadataStorageChecker, error)
}

var _ environs.EnvironProvider = (*kubernetesEnvironProvider)(nil)
var providerInstance = kubernetesEnvironProvider{
	environProviderCredentials: environProviderCredentials{
		cmdRunner: defaultRunner{},
		builtinCredentialGetter: func(ctx context.Context, cmdRunner CommandRunner) (cloud.Credential, error) {
			return attemptMicroK8sCredential(ctx, cmdRunner, decideKubeConfigDir)
		},
	},
	cmdRunner: defaultRunner{},
	builtinCloudGetter: func(cmdRunner CommandRunner) (cloud.Cloud, error) {
		return attemptMicroK8sCloud(cmdRunner, decideKubeConfigDir)
	},
	brokerGetter: func(ctx context.Context, args environs.OpenParams, invalidator environs.CredentialInvalidator) (ClusterMetadataStorageChecker, error) {
		broker, err := caas.New(ctx, args, invalidator)
		if err != nil {
			return nil, errors.Trace(err)
		}

		metaChecker, supported := broker.(ClusterMetadataStorageChecker)
		if !supported {
			return nil, errors.NotSupportedf("cluster metadata ")
		}
		return metaChecker, nil
	},
}

// Version is part of the EnvironProvider interface.
func (kubernetesEnvironProvider) Version() int {
	return 0
}

// CommandRunner allows to run commands on the underlying system
type CommandRunner interface {
	RunCommands(run exec.RunParams) (*exec.ExecResponse, error)
	LookPath(string) (string, error)
}

type defaultRunner struct{}

func (defaultRunner) RunCommands(run exec.RunParams) (*exec.ExecResponse, error) {
	return exec.RunCommands(run)
}

func (defaultRunner) LookPath(file string) (string, error) {
	return osexec.LookPath(file)
}

// NewK8sClients returns the k8s clients to access a cluster.
// Override for testing.
var NewK8sClients = func(c *rest.Config) (
	k8sClient kubernetes.Interface,
	apiextensionsclient apiextensionsclientset.Interface,
	dynamicClient dynamic.Interface,
	err error,
) {
	k8sClient, err = kubernetes.NewForConfig(c)
	if err != nil {
		return nil, nil, nil, err
	}
	apiextensionsclient, err = apiextensionsclientset.NewForConfig(c)
	if err != nil {
		return nil, nil, nil, err
	}
	dynamicClient, err = dynamic.NewForConfig(c)
	if err != nil {
		return nil, nil, nil, err
	}
	return k8sClient, apiextensionsclient, dynamicClient, nil
}

// CloudSpecToK8sRestConfig translates cloudspec to k8s rest config.
func CloudSpecToK8sRestConfig(cloudSpec environscloudspec.CloudSpec) (*rest.Config, error) {
	if cloudSpec.IsControllerCloud {
		rc, err := rest.InClusterConfig()
		if err != nil && err != rest.ErrNotInCluster {
			return nil, errors.Trace(err)
		}
		if rc != nil {
			logger.Tracef(context.TODO(), "using in-cluster config")
			return rc, nil
		}
	}

	if cloudSpec.Credential == nil {
		return nil, errors.Errorf("cloud %v has no credential", cloudSpec.Name)
	}

	var caData []byte
	for _, cacert := range cloudSpec.CACertificates {
		caData = append(caData, cacert...)
	}

	credentialAttrs := cloudSpec.Credential.Attributes()
	return &rest.Config{
		Host:        cloudSpec.Endpoint,
		Username:    credentialAttrs[k8scloud.CredAttrUsername],
		Password:    credentialAttrs[k8scloud.CredAttrPassword],
		BearerToken: credentialAttrs[k8scloud.CredAttrToken],
		TLSClientConfig: rest.TLSClientConfig{
			CertData: []byte(credentialAttrs[k8scloud.CredAttrClientCertificateData]),
			KeyData:  []byte(credentialAttrs[k8scloud.CredAttrClientKeyData]),
			CAData:   caData,
			Insecure: cloudSpec.SkipTLSVerify,
		},
	}, nil
}

func newRestClient(cfg *rest.Config) (rest.Interface, error) {
	return rest.RESTClientFor(cfg)
}

// Open is part of the ContainerEnvironProvider interface.
func (p kubernetesEnvironProvider) Open(ctx context.Context, args environs.OpenParams, invalidator environs.CredentialInvalidator) (caas.Broker, error) {
	logger.Debugf(context.TODO(), "opening model %q.", args.Config.Name())
	if err := p.validateCloudSpec(args.Cloud); err != nil {
		return nil, errors.Annotate(err, "validating cloud spec")
	}
	k8sRestConfig, err := CloudSpecToK8sRestConfig(args.Cloud)
	if err != nil {
		return nil, errors.Trace(err)
	}

	namespace, err := NamespaceForModel(ctx, args.Config.Name(), args.ControllerUUID, k8sRestConfig)
	if err != nil && !errors.Is(err, errors.NotFound) {
		return nil, err
	}

	return newK8sBroker(ctx,
		args.ControllerUUID, k8sRestConfig, args.Config, namespace,
		NewK8sClients, newRestClient, k8swatcher.NewKubernetesNotifyWatcher, k8swatcher.NewKubernetesStringsWatcher,
		utils.RandomPrefix, jujuclock.WallClock)
}

// NamespaceForModel returns the namespace which is associated with the specified model.
func NamespaceForModel(ctx context.Context, modelName string, controllerUUID string, k8sRestConfig *rest.Config) (string, error) {
	if modelName != environsbootstrap.ControllerModelName {
		return modelName, nil
	}
	k8sClient, _, _, err := NewK8sClients(k8sRestConfig)
	if err != nil {
		return "", errors.Trace(err)
	}

	ns, err := findControllerNamespace(ctx, k8sClient, controllerUUID)
	if err != nil {
		return "", errors.Trace(err)
	}
	return ns.Name, nil
}

// CloudSchema returns the schema for adding new clouds of this type.
func (p kubernetesEnvironProvider) CloudSchema() *jsonschema.Schema {
	return nil
}

// Ping tests the connection to the cloud, to verify the endpoint is valid.
func (p kubernetesEnvironProvider) Ping(ctx envcontext.ProviderCallContext, endpoint string) error {
	return errors.NotImplementedf("Ping")
}

// ValidateCloud is specified in the EnvironProvider interface.
func (p kubernetesEnvironProvider) ValidateCloud(ctx context.Context, spec environscloudspec.CloudSpec) error {
	return errors.Annotate(p.validateCloudSpec(spec), "validating cloud spec")
}

// DetectRegions is specified in the environs.CloudRegionDetector interface.
func (p kubernetesEnvironProvider) DetectRegions() ([]cloud.Region, error) {
	return nil, errors.NotFoundf("regions")
}

func (p kubernetesEnvironProvider) validateCloudSpec(spec environscloudspec.CloudSpec) error {
	if err := spec.Validate(); err != nil {
		return errors.Trace(err)
	}
	if _, err := url.Parse(spec.Endpoint); err != nil {
		return errors.NotValidf("endpoint %q", spec.Endpoint)
	}
	if spec.Credential == nil {
		return errors.NotValidf("missing credential")
	}

	if authType := spec.Credential.AuthType(); !k8scloud.SupportedAuthTypes().Contains(authType) {
		return errors.NotSupportedf("%q auth-type", authType)
	}
	return nil
}
