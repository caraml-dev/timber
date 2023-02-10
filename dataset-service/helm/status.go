package helm

import (
	"github.com/caraml-dev/timber/dataset-service/api"
	"helm.sh/helm/v3/pkg/release"
)

// ConvertStatus converts helm status to dataset service api status
func ConvertStatus(status release.Status) api.Status {
	switch status {
	case release.StatusDeployed, release.StatusSuperseded:
		return api.Status_STATUS_DEPLOYED
	case release.StatusUninstalled:
		return api.Status_STATUS_UNINSTALLED
	case release.StatusFailed:
		return api.Status_STATUS_FAILED
	case release.StatusUninstalling, release.StatusPendingInstall, release.StatusPendingUpgrade, release.StatusPendingRollback:
		return api.Status_STATUS_PENDING
	default:
		return api.Status_STATUS_UNSPECIFIED
	}
}
