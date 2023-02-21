package helm

import (
	"helm.sh/helm/v3/pkg/release"

	"github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/model"
)

// ConvertStatus converts helm status to dataset service api status
func ConvertStatus(status release.Status) model.Status {
	switch status {
	case release.StatusDeployed, release.StatusSuperseded:
		return model.StatusDeployed
	case release.StatusUninstalled:
		return model.StatusUninstalled
	case release.StatusFailed:
		return model.StatusFailed
	case release.StatusUninstalling, release.StatusPendingInstall, release.StatusPendingUpgrade, release.StatusPendingRollback:
		return model.StatusPending
	default:
		return model.StatusUnspecified
	}
}

// ConvertStatusToProto converts helm status to dataset service api status
// TODO: remove
func ConvertStatusToProto(status release.Status) api.Status {
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
