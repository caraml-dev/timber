package model

import timberv1 "github.com/caraml-dev/timber/dataset-service/api"

type Status string

const (
	// StatusUnspecified status is not initialized
	StatusUnspecified Status = "STATUS_UNSPECIFIED"
	// StatusDeployed successfully deployed
	StatusDeployed Status = "STATUS_DEPLOYED"
	// StatusUninstalled successfully uninstalled
	StatusUninstalled Status = "STATUS_UNINSTALLED"
	// StatusFailed failed deployment
	StatusFailed Status = "STATUS_FAILED"
	// StatusPending waiting for deployment to complete
	StatusPending Status = "STATUS_PENDING"
)

func (s Status) ToStatusProto() timberv1.Status {
	switch s {
	case StatusDeployed:
		return timberv1.Status_STATUS_DEPLOYED
	case StatusPending:
		return timberv1.Status_STATUS_PENDING
	case StatusFailed:
		return timberv1.Status_STATUS_FAILED
	case StatusUninstalled:
		return timberv1.Status_STATUS_UNINSTALLED
	default:
		return timberv1.Status_STATUS_UNSPECIFIED
	}
}

func StatusFromProto(statusProto timberv1.Status) Status {
	switch statusProto {
	case timberv1.Status_STATUS_DEPLOYED:
		return StatusDeployed
	case timberv1.Status_STATUS_PENDING:
		return StatusPending
	case timberv1.Status_STATUS_FAILED:
		return StatusFailed
	case timberv1.Status_STATUS_UNINSTALLED:
		return StatusUninstalled
	default:
		return StatusUnspecified
	}
}
