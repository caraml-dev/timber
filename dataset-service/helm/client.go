package helm

import (
	"fmt"

	"github.com/caraml-dev/timber/common/log"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// Helm client
type Client interface {
	// ReadChart read a helm chart given by the chart path.
	ReadChart(chartPath string) (*chart.Chart, error)
	// Install installs a new helm release. Failed if there is an existing release.
	Install(releaseName string, namespaceName string, chart *chart.Chart, values map[string]any, actionConfig *action.Configuration) (*release.Release, error)
	// Upgrade upgrades an existing helm release. Failed if there are no existing release.
	Upgrade(releaseName string, namespaceName string, chart *chart.Chart, values map[string]any, actionConfig *action.Configuration) (*release.Release, error)
	// GetRelease get a helm release.
	GetRelease(releaseName string, namespaceName string, actionConfig *action.Configuration) (*release.Release, error)
}

const (
	k8sUser    = "user"
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

// Install a new helm release
func (h *helmClient) Install(releaseName string, namespaceName string, chart *chart.Chart, values map[string]any, actionConfig *action.Configuration) (*release.Release, error) {
	actionConfig, err := h.initializeConfig(actionConfig, namespaceName)
	if err != nil {
		return nil, fmt.Errorf("error initializeConfig: %w", err)
	}

	installation := action.NewInstall(actionConfig)
	installation.ReleaseName = releaseName
	installation.Namespace = namespaceName
	installation.CreateNamespace = true

	log.Debugf("installing helm release: %s, namespace: %s, chart: %s, chart version: %s", releaseName, namespaceName, chart.Name(), chart.Metadata.Version)
	r, err := installation.Run(chart, values)
	if err != nil {
		return nil, err
	}

	log.Debugf("release manifest: %v", r.Manifest)
	return r, nil
}

// Updgrade an existing helm release
func (h *helmClient) Upgrade(releaseName string, namespaceName string, chart *chart.Chart, values map[string]any, actionConfig *action.Configuration) (*release.Release, error) {
	actionConfig, err := h.initializeConfig(actionConfig, namespaceName)
	if err != nil {
		return nil, fmt.Errorf("error initializeConfig: %w", err)
	}

	upgrade := action.NewInstall(actionConfig)
	upgrade.ReleaseName = releaseName
	upgrade.Namespace = namespaceName

	log.Debugf("upgrading helm release: %s, namespace: %s, chart: %s, chart version: %s", releaseName, namespaceName, chart.Name(), chart.Metadata.Version)
	r, err := upgrade.Run(chart, values)
	if err != nil {
		return nil, err
	}

	log.Debugf("release manifest: %v", r.Manifest)
	return r, nil
}

// GetRelease get release name in the given namespace
func (h helmClient) GetRelease(releaseName string, namespaceName string, actionConfig *action.Configuration) (*release.Release, error) {
	actionConfig, err := h.initializeConfig(actionConfig, namespaceName)
	if err != nil {
		return nil, fmt.Errorf("error initializeConfig: %w", err)
	}
	get := action.NewStatus(actionConfig)

	return get.Run(releaseName)
}

// initializeConfig initialize action config
func (h *helmClient) initializeConfig(actionConfig *action.Configuration, namespaceName string) (*action.Configuration, error) {
	if actionConfig == nil {
		actionConfig = new(action.Configuration)
	}
	err := actionConfig.Init(h.clientGetter, namespaceName, helmDriver, log.Debugf)
	if err != nil {
		return nil, err
	}
	return actionConfig, nil
}
