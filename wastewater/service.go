package wastewater

import (
	"context"

	"github.org/anggi-susanto/mrt-go/domain"
)

// WasteWaterRepositoryInterface is the interface that wraps the Create, GetAll, GetByID, Update, and Delete methods.
type WasteWaterRepositoryInterface interface {
	Create(ctx context.Context, w *domain.WasteWaterData) error
	GetAll(ctx context.Context, page, limit int) ([]domain.WasteWaterData, error)
	GetByID(ctx context.Context, id string) (*domain.WasteWaterData, error)
	Update(ctx context.Context, w *domain.WasteWaterData) error
	Delete(ctx context.Context, id string) error
}

// Service is the interface that wraps the Create, GetAll, GetByID, Update, and Delete methods.
type Service struct {
	wasteWaterRepository WasteWaterRepositoryInterface
}

// NewService creates a new instance of the Service struct, initializing it with the provided WasteWaterRepositoryInterface.
//
// Parameters:
// - wasteWaterRepository: The WasteWaterRepositoryInterface implementation used by the Service.
//
// Returns:
// - A pointer to the newly created Service instance.
func NewService(wasteWaterRepository WasteWaterRepositoryInterface) *Service {
	return &Service{
		wasteWaterRepository: wasteWaterRepository,
	}
}

// Create creates a new waste water data record in the service.
//
// ctx: The context.Context object for the request.
// w: The waste water data to be created.
// Returns an error if there was a problem creating the record.
func (s *Service) Create(ctx context.Context, w *domain.WasteWaterData) error {
	return s.wasteWaterRepository.Create(ctx, w)
}

// GetAll retrieves all waste water data with pagination.
//
// ctx context.Context, page int, limit int
// []domain.WasteWaterData, error
func (s *Service) GetAll(ctx context.Context, page, limit int) ([]domain.WasteWaterData, error) {
	return s.wasteWaterRepository.GetAll(ctx, page, limit)
}

// GetByID retrieves a WasteWaterData by ID.
//
// ctx - context.Context for the operation.
// id - string representing the ID of the data.
// Returns a pointer to domain.WasteWaterData and an error.
func (s *Service) GetByID(ctx context.Context, id string) (*domain.WasteWaterData, error) {
	return s.wasteWaterRepository.GetByID(ctx, id)
}

// Delete deletes a WasteWaterData by ID.
//
// ctx - context.Context for the operation.
// id - string representing the ID of the data to be deleted.
// Returns an error if there was a problem deleting the data.
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.wasteWaterRepository.Delete(ctx, id)
}
