// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package deployer

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/juju/errors"
	"github.com/juju/gnuflag"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/api/base"
	"github.com/juju/juju/api/client/resources"
	commoncharm "github.com/juju/juju/api/common/charm"
	"github.com/juju/juju/api/common/charms"
	"github.com/juju/juju/cmd/juju/application/deployer/mocks"
	"github.com/juju/juju/cmd/modelcmd"
	corebase "github.com/juju/juju/core/base"
	corecharm "github.com/juju/juju/core/charm"
	"github.com/juju/juju/core/model"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/internal/charm"
	charmresource "github.com/juju/juju/internal/charm/resource"
	"github.com/juju/juju/testcharms"
	coretesting "github.com/juju/juju/testing"
)

type deployerSuite struct {
	testing.IsolationSuite

	consumeDetails *mocks.MockConsumeDetails
	resolver       *mocks.MockResolver
	deployerAPI    *mocks.MockDeployerAPI
	modelCommand   *mocks.MockModelCommand
	filesystem     *mocks.MockFilesystem
	bundle         *mocks.MockBundle
	charmDeployAPI *mocks.MockCharmDeployAPI
	charmReader    *mocks.MockCharmReader
	charm          *mocks.MockCharm

	deployResourceIDs map[string]string
}

var _ = gc.Suite(&deployerSuite{})

func (s *deployerSuite) SetUpTest(_ *gc.C) {
	s.deployResourceIDs = make(map[string]string)
}

func (s *deployerSuite) TestGetDeployerPredeployedLocalCharm(c *gc.C) {
	defer s.setupMocks(c).Finish()
	s.expectModelGet(c)
	s.expectModelType()

	cfg := s.basicDeployerConfig()
	ch := "local:test-charm"
	s.expectStat(ch, errors.NotFoundf("file"))
	cfg.CharmOrBundle = ch

	s.charmDeployAPI.EXPECT().CharmInfo("local:test-charm").Return(&charms.CharmInfo{
		Meta: &charm.Meta{
			Name: "wordpress",
		},
		Manifest: &charm.Manifest{
			Bases: []charm.Base{{Name: "ubuntu", Channel: charm.Channel{Track: "20.04", Risk: "stable"}}},
		},
	}, nil)

	factory := s.newDeployerFactory()
	deployer, err := factory.GetDeployer(context.Background(), cfg, s.charmDeployAPI, s.resolver)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(deployer.String(), gc.Equals, fmt.Sprintf("deploy predeployed local charm: %s", ch))
}

func (s *deployerSuite) TestGetDeployerLocalCharm(c *gc.C) {
	defer s.setupMocks(c).Finish()
	s.expectModelGet(c)
	cfg := s.basicDeployerConfig(corebase.MustParseBaseFromString("ubuntu@22.04"))
	s.expectModelType()

	dir := c.MkDir()
	charmPath := testcharms.RepoWithSeries("bionic").ClonedDirPath(dir, "multi-series")
	s.expectStat(charmPath, nil)
	cfg.CharmOrBundle = charmPath

	factory := s.newDeployerFactory()
	deployer, err := factory.GetDeployer(context.Background(), cfg, s.charmDeployAPI, s.resolver)
	c.Assert(err, jc.ErrorIsNil)
	ch := "local:multi-series-1"
	c.Assert(deployer.String(), gc.Equals, fmt.Sprintf("deploy local charm: %s", ch))
}

func (s *deployerSuite) TestGetDeployerLocalCharmPathWithSchema(c *gc.C) {
	defer s.setupMocks(c).Finish()
	s.expectModelGet(c)
	cfg := s.basicDeployerConfig(corebase.MustParseBaseFromString("ubuntu@22.04"))
	s.expectModelType()

	dir := c.MkDir()
	charmPath := testcharms.RepoWithSeries("bionic").ClonedDirPath(dir, "multi-series")
	charmPath = "local:" + charmPath
	s.expectStat(charmPath, errors.NotFoundf("file"))
	cfg.CharmOrBundle = charmPath

	factory := s.newDeployerFactory()
	deployer, err := factory.GetDeployer(context.Background(), cfg, s.charmDeployAPI, s.resolver)
	c.Assert(err, jc.ErrorIsNil)
	ch := "local:multi-series-1"
	c.Assert(deployer.String(), gc.Equals, fmt.Sprintf("deploy local charm: %s", ch))
}

func (s *deployerSuite) TestGetDeployerLocalCharmError(c *gc.C) {
	defer s.setupMocks(c).Finish()
	path := "./bad.charm"
	factory := s.newDeployerFactory().(*factory)
	factory.charmOrBundle = path

	s.expectStat(path, os.ErrNotExist)

	_, err := factory.GetDeployer(context.Background(), DeployerConfig{CharmOrBundle: path}, s.charmDeployAPI, nil)
	c.Assert(err, gc.ErrorMatches, `no charm was found at "./bad.charm"`)
}

func (s *deployerSuite) TestGetDeployerCharmHubCharm(c *gc.C) {
	defer s.setupMocks(c).Finish()
	ch := "ch:test-charm"

	s.expectModelType()
	// NotValid ensures that maybeReadRepositoryBundle won't find
	// charmOrBundle is a bundle.
	s.expectResolveBundleURL(errors.NotValidf("not a bundle"), 1)

	cfg := s.basicDeployerConfig()
	s.expectStat(ch, errors.NotFoundf("file"))
	cfg.CharmOrBundle = ch

	factory := s.newDeployerFactory()
	deployer, err := factory.GetDeployer(context.Background(), cfg, s.charmDeployAPI, s.resolver)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(deployer.String(), gc.Equals, fmt.Sprintf("deploy charm: %s", ch))
}

func (s *deployerSuite) TestGetDeployerCharmHubCharmWithRevision(c *gc.C) {
	cfg := s.channelDeployerConfig()
	cfg.Revision = 42
	cfg.Channel, _ = charm.ParseChannel("stable")
	curl := "ch:test-charm"
	deployer, err := s.testGetDeployerRepositoryCharmWithRevision(c, curl, cfg)
	c.Assert(err, jc.ErrorIsNil)
	str := fmt.Sprintf("deploy charm: %s with revision %d will refresh from channel %s", curl, cfg.Revision, cfg.Channel.String())
	c.Assert(deployer.String(), gc.Equals, str)
}

func (s *deployerSuite) TestGetDeployerCharmHubCharmWithRevisionFail(c *gc.C) {
	cfg := s.basicDeployerConfig()
	cfg.Revision = 42
	curl := "ch:test-charm"
	_, err := s.testGetDeployerRepositoryCharmWithRevision(c, curl, cfg)
	c.Assert(err, gc.ErrorMatches, "specifying a revision requires a channel for future upgrades. Please use --channel")
}

func (s *deployerSuite) testGetDeployerRepositoryCharmWithRevision(c *gc.C, curl string, cfg DeployerConfig) (Deployer, error) {
	defer s.setupMocks(c).Finish()
	s.expectModelType()
	// NotValid ensures that maybeReadRepositoryBundle won't find
	// charmOrBundle is a bundle.
	s.expectResolveBundleURL(errors.NotValidf("not a bundle"), 1)

	cfg.CharmOrBundle = curl
	s.expectStat(cfg.CharmOrBundle, errors.NotFoundf("file"))

	factory := s.newDeployerFactory()
	return factory.GetDeployer(context.Background(), cfg, s.charmDeployAPI, s.resolver)
}

func (s *deployerSuite) TestBaseOverride(c *gc.C) {
	defer s.setupMocks(c).Finish()
	s.expectModelType()
	s.expectResolveBundleURL(errors.NotValidf("not a bundle"), 1)

	cfg := s.basicDeployerConfig(corebase.MustParseBaseFromString("ubuntu@21.04"))
	curl := "ch:test-charm"
	s.expectStat(curl, errors.NotFoundf("file"))
	cfg.CharmOrBundle = curl

	factory := s.newDeployerFactory()
	deployer, err := factory.GetDeployer(context.Background(), cfg, s.charmDeployAPI, s.resolver)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(deployer.String(), gc.Equals, fmt.Sprintf("deploy charm: %s", curl))

	charmDeployer := deployer.(*repositoryCharm)
	c.Assert(charmDeployer.id.Origin.Base.OS, gc.Equals, "ubuntu")
	c.Assert(charmDeployer.id.Origin.Base.Channel.String(), gc.Equals, "21.04/stable")
}

func (s *deployerSuite) TestGetDeployerLocalBundle(c *gc.C) {
	defer s.setupMocks(c).Finish()

	cfg := s.basicDeployerConfig(corebase.MustParseBaseFromString("ubuntu@18.04"))
	cfg.FlagSet = &gnuflag.FlagSet{}
	s.expectModelType()

	content := `
      series: xenial
      applications:
          wordpress:
              charm: wordpress
              num_units: 1
          mysql:
              charm: mysql
              num_units: 2
      relations:
          - ["wordpress:db", "mysql:server"]
`
	bundlePath := s.makeBundleDir(c, content)
	s.expectStat(bundlePath, nil)
	cfg.CharmOrBundle = bundlePath

	factory := s.newDeployerFactory()
	deployer, err := factory.GetDeployer(context.Background(), cfg, s.charmDeployAPI, s.resolver)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(deployer.String(), gc.Equals, fmt.Sprintf("deploy local bundle from: %s", bundlePath))
}

func (s *deployerSuite) TestGetDeployerCharmHubBundleWithChannel(c *gc.C) {
	defer s.setupMocks(c).Finish()

	bundle := "ch:test-bundle"
	cfg := s.channelDeployerConfig()
	cfg.CharmOrBundle = bundle

	s.expectResolveBundleURL(nil, 1)

	cfg.Base = corebase.MustParseBaseFromString("ubuntu@18.04")
	cfg.FlagSet = &gnuflag.FlagSet{}
	s.expectStat(cfg.CharmOrBundle, errors.NotFoundf("file"))
	s.expectModelType()
	s.expectGetBundle(nil)
	s.expectData()
	s.expectBundleBytes()

	factory := s.newDeployerFactory()
	deployer, err := factory.GetDeployer(context.Background(), cfg, s.charmDeployAPI, s.resolver)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(deployer.String(), gc.Equals, fmt.Sprintf("deploy bundle: %s from channel edge", bundle))
}

func (s *deployerSuite) TestGetDeployerCharmHubBundleWithRevision(c *gc.C) {
	defer s.setupMocks(c).Finish()

	bundle := "ch:test-bundle"
	cfg := s.basicDeployerConfig()
	cfg.Revision = 8
	cfg.CharmOrBundle = bundle
	cfg.Base = corebase.MustParseBaseFromString("ubuntu@18.04")
	cfg.FlagSet = &gnuflag.FlagSet{}
	s.expectStat(cfg.CharmOrBundle, errors.NotFoundf("file"))
	s.expectModelType()
	s.expectGetBundle(nil)
	s.expectData()
	s.expectBundleBytes()

	s.expectResolveBundleURL(nil, 1)
	factory := s.newDeployerFactory()
	deployer, err := factory.GetDeployer(context.Background(), cfg, s.charmDeployAPI, s.resolver)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(deployer.String(), gc.Equals, fmt.Sprintf("deploy bundle: %s with revision 8", bundle))
}

func (s *deployerSuite) TestGetDeployerCharmHubBundleWithRevisionURL(c *gc.C) {
	defer s.setupMocks(c).Finish()

	bundle := "ch:test-bundle-8"
	cfg := s.basicDeployerConfig(corebase.MustParseBaseFromString("ubuntu@18.04"))
	cfg.CharmOrBundle = bundle
	cfg.FlagSet = &gnuflag.FlagSet{}
	s.expectStat(cfg.CharmOrBundle, errors.NotFoundf("file"))
	s.expectModelType()

	factory := s.newDeployerFactory()
	_, err := factory.GetDeployer(context.Background(), cfg, s.charmDeployAPI, s.resolver)
	c.Assert(err, gc.ErrorMatches, "cannot specify revision in a charm or bundle name. Please use --revision.")
}

func (s *deployerSuite) TestGetDeployerCharmHubBundleError(c *gc.C) {
	defer s.setupMocks(c).Finish()

	bundle := "ch:test-bundle"
	cfg := s.channelDeployerConfig(corebase.MustParseBaseFromString("ubuntu@18.04"))
	cfg.Revision = 42
	cfg.CharmOrBundle = bundle
	cfg.FlagSet = &gnuflag.FlagSet{}
	s.expectStat(cfg.CharmOrBundle, errors.NotFoundf("file"))
	s.expectModelType()
	s.expectResolveBundleURL(nil, 1)

	factory := s.newDeployerFactory()
	_, err := factory.GetDeployer(context.Background(), cfg, s.charmDeployAPI, s.resolver)
	c.Assert(err, gc.ErrorMatches, "revision and channel are mutually exclusive when deploying a bundle. Please choose one.")
}

func (s *deployerSuite) TestResolveCharmURL(c *gc.C) {
	tests := []struct {
		defaultSchema charm.Schema
		path          string
		url           *charm.URL
		err           error
	}{{
		defaultSchema: charm.CharmHub,
		path:          "wordpress",
		url:           &charm.URL{Schema: "ch", Name: "wordpress", Revision: -1},
	}, {
		defaultSchema: charm.CharmHub,
		path:          "ch:wordpress",
		url:           &charm.URL{Schema: "ch", Name: "wordpress", Revision: -1},
	}, {
		defaultSchema: charm.CharmHub,
		path:          "local:wordpress",
		url:           &charm.URL{Schema: "local", Name: "wordpress", Revision: -1},
	}}
	for i, test := range tests {
		c.Logf("%d %s", i, test.path)
		url, err := resolveCharmURL(test.path, test.defaultSchema)
		if test.err != nil {
			c.Assert(err, gc.ErrorMatches, test.err.Error())
		} else {
			c.Assert(err, jc.ErrorIsNil)
			c.Assert(url, gc.DeepEquals, test.url)
		}
	}
}

func (s *deployerSuite) TestValidateResourcesNeededForLocalDeployCAAS(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.modelCommand.EXPECT().ModelType().Return(model.CAAS, nil).AnyTimes()

	f := &factory{
		model: s.modelCommand,
	}

	err := f.validateResourcesNeededForLocalDeploy(&charm.Meta{})
	c.Assert(err, jc.ErrorIsNil)
}

func (s *deployerSuite) TestValidateResourcesNeededForLocalDeployIAAS(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.modelCommand.EXPECT().ModelType().Return(model.IAAS, nil).AnyTimes()

	f := &factory{
		model: s.modelCommand,
	}

	err := f.validateResourcesNeededForLocalDeploy(&charm.Meta{})
	c.Assert(err, jc.ErrorIsNil)
}

func (s *deployerSuite) makeBundleDir(c *gc.C, content string) string {
	bundlePath := filepath.Join(c.MkDir(), "example")
	c.Assert(os.Mkdir(bundlePath, 0777), jc.ErrorIsNil)
	err := os.WriteFile(filepath.Join(bundlePath, "bundle.yaml"), []byte(content), 0644)
	c.Assert(err, jc.ErrorIsNil)
	err = os.WriteFile(filepath.Join(bundlePath, "README.md"), []byte("README"), 0644)
	c.Assert(err, jc.ErrorIsNil)
	return bundlePath
}

func (s *deployerSuite) newDeployerFactory() DeployerFactory {
	dep := DeployerDependencies{
		DeployResources: func(
			context.Context,
			string,
			resources.CharmID,
			map[string]string,
			map[string]charmresource.Meta,
			base.APICallCloser,
			modelcmd.Filesystem,
		) (ids map[string]string, err error) {
			return s.deployResourceIDs, nil
		},
		Model:                s.modelCommand,
		NewConsumeDetailsAPI: func(url *charm.OfferURL) (ConsumeDetails, error) { return s.consumeDetails, nil },
		FileSystem:           s.filesystem,
		CharmReader:          fsCharmReader{},
	}
	return NewDeployerFactory(dep)
}

func (s *deployerSuite) basicDeployerConfig(bases ...corebase.Base) DeployerConfig {
	var base corebase.Base
	if len(bases) == 0 {
		base = corebase.MustParseBaseFromString("ubuntu@20.04")
	} else {
		base = bases[0]
	}
	return DeployerConfig{
		Base:     base,
		Revision: -1,
	}
}

func (s *deployerSuite) channelDeployerConfig(bases ...corebase.Base) DeployerConfig {
	var base corebase.Base
	if len(bases) == 0 {
		base = corebase.MustParseBaseFromString("ubuntu@20.04")
	} else {
		base = bases[0]
	}
	return DeployerConfig{
		Base: base,
		Channel: charm.Channel{
			Risk: "edge",
		},
		Revision: -1,
	}
}

func (s *deployerSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)
	s.consumeDetails = mocks.NewMockConsumeDetails(ctrl)
	s.resolver = mocks.NewMockResolver(ctrl)
	s.bundle = mocks.NewMockBundle(ctrl)
	s.deployerAPI = mocks.NewMockDeployerAPI(ctrl)
	s.modelCommand = mocks.NewMockModelCommand(ctrl)
	s.filesystem = mocks.NewMockFilesystem(ctrl)
	s.charmDeployAPI = mocks.NewMockCharmDeployAPI(ctrl)
	s.charmReader = mocks.NewMockCharmReader(ctrl)
	s.charm = mocks.NewMockCharm(ctrl)
	return ctrl
}

func (s *deployerSuite) expectResolveBundleURL(err error, times int) {
	s.resolver.EXPECT().ResolveBundleURL(
		gomock.AssignableToTypeOf(&charm.URL{}),
		gomock.AssignableToTypeOf(commoncharm.Origin{}),
	).DoAndReturn(
		// Ensure the same curl that is provided, is returned.
		func(curl *charm.URL, origin commoncharm.Origin) (*charm.URL, commoncharm.Origin, error) {
			return curl, origin, err
		}).Times(times)
}

func (s *deployerSuite) expectStat(name string, err error) {
	s.filesystem.EXPECT().Stat(name).Return(nil, err)
}

func (s *deployerSuite) expectModelGet(c *gc.C) {
	minimal := map[string]interface{}{
		"name":            "test",
		"type":            "manual",
		"uuid":            coretesting.ModelTag.Id(),
		"controller-uuid": coretesting.ControllerTag.Id(),
		"firewall-mode":   "instance",
		"secret-backend":  "auto",
		// While the ca-cert bits aren't entirely minimal, they avoid the need
		// to set up a fake home.
		"ca-cert":        coretesting.CACert,
		"ca-private-key": coretesting.CAKey,
	}
	cfg, err := config.New(true, minimal)
	c.Assert(err, jc.ErrorIsNil)
	s.charmDeployAPI.EXPECT().ModelGet().Return(cfg.AllAttrs(), nil)
}

func (s *deployerSuite) expectModelType() {
	s.modelCommand.EXPECT().ModelType().Return(model.IAAS, nil).AnyTimes()
}

func (s *deployerSuite) expectGetBundle(err error) {
	s.resolver.EXPECT().GetBundle(gomock.Any(), gomock.AssignableToTypeOf(&charm.URL{}), gomock.Any(), gomock.Any()).Return(s.bundle, err)
}

func (s *deployerSuite) expectData() {
	s.bundle.EXPECT().Data().Return(&charm.BundleData{})
}

func (s *deployerSuite) expectBundleBytes() {
	s.bundle.EXPECT().BundleBytes().Return([]byte{})
}

// TODO (stickupkid): Remove this in favour of better unit tests with mocks.
// Currently most of the tests are integration tests, pretending to be unit
// tests.
type fsCharmReader struct{}

// NewCharmAtPath attempts to read a charm from a path on the filesystem.
func (fsCharmReader) NewCharmAtPath(path string) (charm.Charm, *charm.URL, error) {
	return corecharm.NewCharmAtPath(path)
}
