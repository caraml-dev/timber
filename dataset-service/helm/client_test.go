package helm

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chartutil"
	kubefake "helm.sh/helm/v3/pkg/kube/fake"
	"helm.sh/helm/v3/pkg/registry"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/storage"
	"helm.sh/helm/v3/pkg/storage/driver"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/caraml-dev/timber/common/log"
)

func actionConfigFixture(t *testing.T) *action.Configuration {
	t.Helper()

	registryClient, err := registry.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	return &action.Configuration{
		Releases:       storage.Init(driver.NewMemory()),
		KubeClient:     &kubefake.FailingKubeClient{PrintingKubeClient: kubefake.PrintingKubeClient{Out: ioutil.Discard}},
		Capabilities:   chartutil.DefaultCapabilities,
		RegistryClient: registryClient,
		Log:            log.Debugf,
	}
}

func Test_helmClient_Install(t *testing.T) {
	type fields struct {
		clientGetter genericclioptions.RESTClientGetter
		chartCache   map[string]*chart.Chart
	}
	type args struct {
		releaseName   string
		namespaceName string
		chart         *chart.Chart
		values        map[string]any
		actionConfig  *action.Configuration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test install",
			fields: fields{
				clientGetter: nil,
				chartCache:   map[string]*chart.Chart{},
			},
			args: args{
				releaseName:   "my-release",
				namespaceName: "my-namespace",
				chart:         buildChart(),
				values: map[string]any{
					"name": any("debug-container"),
				},
				actionConfig: actionConfigFixture(t),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &helmClient{
				clientGetter: tt.fields.clientGetter,
				chartCache:   tt.fields.chartCache,
			}
			got, err := h.Install(tt.args.releaseName, tt.args.namespaceName, tt.args.chart, tt.args.values, tt.args.actionConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("Install() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.args.releaseName, got.Name)
			assert.Equal(t, tt.args.namespaceName, got.Namespace)
			assert.Equal(t, release.StatusDeployed, got.Info.Status)
			assert.Equal(t, "Install complete", got.Info.Description)
			assert.NotEmpty(t, got.Manifest)
			assert.Contains(t, got.Manifest, "name: debug-container")
		})
	}
}

func Test_helmClient_Upgrade(t *testing.T) {
	type fields struct {
		clientGetter genericclioptions.RESTClientGetter
		chartCache   map[string]*chart.Chart
	}
	type args struct {
		releaseName   string
		namespaceName string
		chart         *chart.Chart
		oldValues     map[string]any
		values        map[string]any
		actionConfig  *action.Configuration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test upgrade",
			fields: fields{
				clientGetter: nil,
				chartCache:   map[string]*chart.Chart{},
			},

			args: args{
				releaseName:   "my-release",
				namespaceName: "my-namespace",
				chart:         buildChart(),
				values: map[string]any{
					"name": any("debug-container-updated"),
				},
				oldValues: map[string]any{
					"name": any("debug-container"),
				},
				actionConfig: actionConfigFixture(t),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &helmClient{
				clientGetter: tt.fields.clientGetter,
				chartCache:   tt.fields.chartCache,
			}

			_, err := h.Install(tt.args.releaseName, tt.args.namespaceName, tt.args.chart, tt.args.oldValues, tt.args.actionConfig)
			assert.NoError(t, err)

			got, err := h.Upgrade(tt.args.releaseName, tt.args.namespaceName, tt.args.chart, tt.args.values, tt.args.actionConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("Upgrade() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.args.releaseName, got.Name)
			assert.Equal(t, tt.args.namespaceName, got.Namespace)
			assert.Equal(t, release.StatusDeployed, got.Info.Status)
			assert.Equal(t, "Upgrade complete", got.Info.Description)
			assert.NotEmpty(t, got.Manifest)
			assert.Contains(t, got.Manifest, "name: debug-container-updated")
		})
	}
}

func buildChart() *chart.Chart {
	template := `apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: debug
  name: debug
spec:
  replicas: 1
  selector:
    matchLabels:
      app: debug
  strategy: {}
  template:
    metadata:
      labels:
        app: debug
    spec:
      containers:
      - image: ubuntu
        name: {{ .Values.name }}
`

	return &chart.Chart{
		Metadata: &chart.Metadata{
			APIVersion: "v1",
			Name:       "test-chart",
			Version:    "0.1.0",
		},
		Templates: []*chart.File{
			{Name: "templates/deployment", Data: []byte(template)},
		},
	}
}
