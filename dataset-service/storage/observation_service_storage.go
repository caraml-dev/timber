package storage

import (
	"context"
	"errors"

	dserrors "github.com/caraml-dev/timber/dataset-service/errors"
	"github.com/caraml-dev/timber/dataset-service/model"
	"gorm.io/gorm"
)

// observationServiceEntityName entity name for observation service
const observationServiceEntityName = "observation_service"

// ObservationService interface providing access for ObservationService storage
type ObservationService interface {
	// Get an observation service given its identifier
	Get(ctx context.Context, input GetInput) (*model.ObservationService, error)
	// Create a new observation service and return the stored observation service with ID populated or error
	Create(ctx context.Context, lw *model.ObservationService) (*model.ObservationService, error)
	// Update an existing observation service and return the stored observation service or error
	Update(ctx context.Context, lw *model.ObservationService) (*model.ObservationService, error)
	// List all observation service given the list input
	List(ctx context.Context, listInput ListInput) ([]*model.ObservationService, error)
}

type observationService struct {
	db *gorm.DB
}

// NewObservationService creates new observation service storage
func NewObservationService(db *gorm.DB) ObservationService {
	return &observationService{db: db}
}

// Get an observation service given its identifier
func (o *observationService) Get(ctx context.Context, input GetInput) (*model.ObservationService, error) {
	observationService := &model.ObservationService{}
	tx := o.db.WithContext(ctx).Where(&model.ObservationService{
		Base: model.Base{
			ID:        input.ID,
			ProjectID: input.ProjectID,
		},
		Name: input.Name,
	}).Take(&observationService)

	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return observationService, dserrors.NewNotFoundError(observationServiceEntityName, input)
	}

	return observationService, tx.Error
}

// Create a new observation service and return the stored observation service with ID populated or error
func (o *observationService) Create(ctx context.Context, obs *model.ObservationService) (*model.ObservationService, error) {
	tx := o.db.WithContext(ctx).Create(obs)
	if tx.Error != nil && errors.As(tx.Error, &duplicateEntryError) {
		// handle duplicate
		return nil, dserrors.NewConflictError(observationServiceEntityName, tx.Error)
	}

	return obs, tx.Error
}

// Update an existing observation service and return the stored observation service or error
func (o *observationService) Update(ctx context.Context, obs *model.ObservationService) (*model.ObservationService, error) {
	tx := o.db.WithContext(ctx).Updates(obs)
	return obs, tx.Error
}

// List all observation services given the list input
func (o *observationService) List(ctx context.Context, listInput ListInput) ([]*model.ObservationService, error) {
	var observationServices []*model.ObservationService
	tx := o.db.WithContext(ctx).Where(&model.ObservationService{
		Base: model.Base{
			ProjectID: listInput.ProjectID,
		},
	}).Limit(listInput.Limit).
		Offset(listInput.Offset).
		// TODO: allow users to provide the ordering option
		Order("id").
		Find(&observationServices)
	return observationServices, tx.Error
}
