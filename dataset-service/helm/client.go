package helm

import (
	"errors"
	"fmt"
	"time"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/storage/driver"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/caraml-dev/timber/common/log"
)

// Client is interface for a helm client
type Client interface {
	// ReadChart read a helm chart given by the chart path.
	ReadChart(chartPath string) (*chart.Chart, error)
	// InstallOrUpgrade upgrades an existing helm release or install a new one.
	InstallOrUpgrade(release string, ns string, chart *chart.Chart, values map[string]any, actionConfig *action.Configuration) (*release.Release, error)
	// Uninstall uninstalls an existing helm release.
	Uninstall(release string, ns string, actionConfig *action.Configuration) error
}

const (
	helmDriver = "secret"
)

// NewClient create a helm client that's connected to a cluster specified by the kubeConfig
func NewClient(kubeConfig string) Client {
	envSettings := cli.New()
	envSettings.KubeConfig = kubeConfig

	return &helmClient{
		clientGetter: envSettings.RESTClientGetter(),
		chartCache:   make(map[string]*chart.Chart),
	}
}

type helmClient struct {
	clientGetter genericclioptions.RESTClientGetter
	chartCache   map[string]*chart.Chart
}

// ReadChart read a helm chart located in chartPath.
// in-memory caching is implemented to avoid repeatedly re-loading same chart
func (h *helmClient) ReadChart(chartPath string) (*chart.Chart, error) {
	c, ok := h.chartCache[chartPath]
	if ok {
		log.Debugf("using cached chart for %s", chartPath)
		return c, nil
	}

	settings := cli.New()
	chartPathOption := action.ChartPathOptions{}
	chartPath, err := chartPathOption.LocateChart(chartPath, settings)
	if err != nil {
		return nil, fmt.Errorf("error locating chart %s, %w", chartPath, err)
	}

	c, err = loader.Load(chartPath)
	if err != nil {
		return nil, fmt.Errorf("error loading chart %s, %w", chartPath, err)
	}

	// add to cache
	h.chartCache[chartPath] = c

	return c, nil
}

// InstallOrUpgrade installs or upgrades an existing helm release
func (h *helmClient) InstallOrUpgrade(release string,
	namespace string,
	chart *chart.Chart,
	values map[string]any,
	actionConfig *action.Configuration) (*release.Release, error) {

	actionConfig, err := h.initializeConfig(actionConfig, namespace)
	if err != nil {
		return nil, fmt.Errorf("error initializeConfig: %w", err)
	}

	_, err = h.getRelease(release, namespace, actionConfig)
	if err != nil {
		if !errors.Is(err, driver.ErrReleaseNotFound) {
			return nil, err
		}

		// install
		return h.install(release, namespace, chart, values, actionConfig)
	}

	upgrade := h.newUpgradeAction(actionConfig, namespace)

	log.Debugf("upgrading helm release: %s, namespace: %s, chart: %s, chart version: %s", release, namespace, chart.Name(), chart.Metadata.Version)
	r, err := upgrade.Run(release, chart, values)
	if err != nil {
		return nil, err
	}

	log.Debugf("release manifest: %v", r.Manifest)
	return r, nil
}

// install a new helm release and block until completion
func (h *helmClient) install(release string,
	namespace string,
	chart *chart.Chart,
	values map[string]any,
	actionConfig *action.Configuration) (*release.Release, error) {

	actionConfig, err := h.initializeConfig(actionConfig, namespace)
	if err != nil {
		return nil, fmt.Errorf("error initializeConfig: %w", err)
	}

	installation := h.newInstallAction(actionConfig, release, namespace)

	log.Debugf("installing helm release: %s, namespace: %s, chart: %s, chart version: %s", release, namespace, chart.Name(), chart.Metadata.Version)
	r, err := installation.Run(chart, values)
	if err != nil {
		return nil, err
	}

	log.Debugf("release manifest: %v", r.Manifest)
	return r, nil
}

// Uninstall a helm release
func (h *helmClient) Uninstall(release string, namespace string, actionConfig *action.Configuration) error {
	actionConfig, err := h.initializeConfig(actionConfig, namespace)
	if err != nil {
		return fmt.Errorf("error initializeConfig: %w", err)
	}

	// check if release exists
	_, err = h.getRelease(release, namespace, actionConfig)
	if err != nil {
		if !errors.Is(err, driver.ErrReleaseNotFound) {
			return err
		}

		log.Debugf("helm release: %s, namespace: %s does not exists", release, namespace)
		return nil
	}

	uninstall := h.newUninstallAction(actionConfig)
	log.Debugf("uninstalling helm release: %s, namespace: %s", release, namespace)
	_, err = uninstall.Run(release)
	return err
}

// create new installation action
func (h *helmClient) newInstallAction(actionConfig *action.Configuration, release string, namespace string) *action.Install {
	installation := action.NewInstall(actionConfig)
	installation.ReleaseName = release
	installation.Namespace = namespace
	installation.CreateNamespace = true
	installation.Wait = true
	installation.Timeout = time.Minute * 10 // TODO: make it configurable

	return installation
}

// create new action.InstallOrUpgrade
func (h *helmClient) newUpgradeAction(actionConfig *action.Configuration, namespace string) *action.Upgrade {
	upgrade := action.NewUpgrade(actionConfig)
	upgrade.Namespace = namespace
	upgrade.Wait = true
	upgrade.Timeout = time.Minute * 10 // TODO: make it configurable

	return upgrade
}

func (h *helmClient) newUninstallAction(config *action.Configuration) *action.Uninstall {
	uninstall := action.NewUninstall(config)
	uninstall.Wait = true
	uninstall.Timeout = time.Minute * 10 // TODO: make it configurable

	return uninstall
}

// GetRelease get release name in the given namespace
func (h *helmClient) getRelease(releaseName string, namespace string, actionConfig *action.Configuration) (*release.Release, error) {
	actionConfig, err := h.initializeConfig(actionConfig, namespace)
	if err != nil {
		return nil, fmt.Errorf("error initializeConfig: %w", err)
	}
	get := action.NewStatus(actionConfig)

	return get.Run(releaseName)
}

// initializeConfig initialize action config
func (h *helmClient) initializeConfig(actionConfig *action.Configuration, namespace string) (*action.Configuration, error) {
	if actionConfig == nil {
		actionConfig = new(action.Configuration)
		err := actionConfig.Init(h.clientGetter, namespace, helmDriver, log.Debugf)
		if err != nil {
			return nil, err
		}
	}

	return actionConfig, nil
}
