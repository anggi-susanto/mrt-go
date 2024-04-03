package sensor

import (
	"context"

	"github.com/anggi-susanto/mrt-go/domain"
)

// SensorRepositoryInterface is an autogenerated interface for SensorRepository
type SensorRepositoryInterface interface {
	Create(ctx context.Context, w *domain.SensorRequest) error
	GetAll(ctx context.Context, page int, limit int) ([]domain.Sensor, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, w *domain.Sensor) error
	GetByID(ctx context.Context, id string) (*domain.Sensor, error)
}

// Service is the interface that wraps the Create, GetAll, GetByID, Update, and Delete methods.
type Service struct {
	sensorRepository SensorRepositoryInterface
}

// NewService creates a new instance of the Service struct, initializing it with the provided SensorRepositoryInterface.
//
// Parameters:
// - sensorRepository: The SensorRepositoryInterface implementation used by the Service.
//
// Returns:
// - A pointer to the newly created Service instance.
func NewService(sensorRepository SensorRepositoryInterface) *Service {
	return &Service{
		sensorRepository: sensorRepository,
	}
}

// Create creates a new waste water data record in the service.
//
// ctx: The context.Context object for the request.
// w: The waste water data to be created.
// Returns an error if there was a problem creating the record.
func (s *Service) Create(ctx context.Context, w *domain.SensorRequest) error {
	return s.sensorRepository.Create(ctx, w)
}

// GetAll retrieves all waste water data with pagination.
//
// ctx context.Context, page int, limit int
// []domain.SensorData, error
func (s *Service) GetAll(ctx context.Context, page, limit int) ([]domain.Sensor, error) {
	return s.sensorRepository.GetAll(ctx, page, limit)
}

// GetByID retrieves a SensorData by ID.
//
// ctx - context.Context for the operation.
// id - string representing the ID of the data.
// Returns a pointer to domain.SensorData and an error.
func (s *Service) GetByID(ctx context.Context, id string) (*domain.Sensor, error) {
	return s.sensorRepository.GetByID(ctx, id)
}

// Delete deletes a SensorData by ID.
//
// ctx - context.Context for the operation.
// id - string representing the ID of the data to be deleted.
// Returns an error if there was a problem deleting the data.
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.sensorRepository.Delete(ctx, id)
}

// Update updates a SensorData.
//
// ctx - context.Context for the operation.
// w - pointer to domain.SensorData representing the data to be updated.
// Returns an error if there was a problem updating the data.
func (s *Service) Update(ctx context.Context, w *domain.Sensor) error {
	return s.sensorRepository.Update(ctx, w)
}