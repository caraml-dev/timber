package model

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
	// StatusFailed waiting for deployment to complete
	StatusPending Status = "STATUS_PENDING"
)
