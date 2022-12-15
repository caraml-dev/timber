package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	tu "github.com/caraml-dev/timber/common/testutils"
)

func TestDefaultConfigs(t *testing.T) {
	host := "localhost"
	port := 3000
	sentryURL := "https://xx.xx.xx"
	sentryLabels := map[string]string{}

	cfg := tu.Config{}
	var filePaths []string
	err := ParseConfig(&cfg, filePaths)

	require.NoError(t, err)
	require.Equal(t, host, cfg.Host)
	require.Equal(t, port, cfg.Port)
	require.Equal(t, sentryURL, cfg.Sentry.URL)
	require.Equal(t, sentryLabels, cfg.Sentry.Labels)
}

func TestFileConfigs(t *testing.T) {
	host := "localhost"
	port := 3030
	sentryURL := "https://yy.yy.yy"
	sentryLabels := map[string]string{"env": "local"}

	cfg := tu.Config{}
	filePaths := []string{"../testdata/config1.yaml"}

	err := ParseConfig(&cfg, filePaths)
	require.NoError(t, err)
	require.Equal(t, host, cfg.Host)
	require.Equal(t, port, cfg.Port)
	require.Equal(t, sentryURL, cfg.Sentry.URL)
	require.Equal(t, sentryLabels, cfg.Sentry.Labels)
}

func TestEnvConfigs(t *testing.T) {
	host := "envhost"
	port := "9999"
	sentryURL := "https://zz.zz.zz"
	sentryLabels := map[string]string{}

	os.Setenv("HOST", host)
	os.Setenv("PORT", port)
	os.Setenv("SENTRY::URL", sentryURL)

	cfg := tu.Config{}
	var filePaths []string

	err := ParseConfig(&cfg, filePaths)
	require.NoError(t, err)
	require.Equal(t, host, cfg.Host)
	require.Equal(t, port, fmt.Sprint(cfg.Port))
	require.Equal(t, sentryURL, cfg.Sentry.URL)
	require.Equal(t, sentryLabels, cfg.Sentry.Labels)
}
