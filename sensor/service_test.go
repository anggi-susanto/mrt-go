package sensor_test

import (
	"context"
	"errors"
	"testing"

	"github.com/anggi-susanto/mrt-go/domain"
	"github.com/anggi-susanto/mrt-go/sensor"
	"github.com/anggi-susanto/mrt-go/sensor/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceCreate(t *testing.T) {
	mockSensor := domain.SensorRequest{
		Name: "sensor",
	}
	t.Run("Success", func(t *testing.T) {
		mockSensorRepo := new(mocks.SensorRepositoryInterface)
		mockSensorRepo.On("Create", mock.Anything, mock.Anything).Return(nil)
		s := sensor.NewService(mockSensorRepo)
		err := s.Create(context.Background(), &mockSensor)
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockSensorRepo := new(mocks.SensorRepositoryInterface)
		mockSensorRepo.On("Create", mock.Anything, mock.Anything).Return(errors.New("error")).Once()
		s := sensor.NewService(mockSensorRepo)
		err := s.Create(context.Background(), &mockSensor)
		assert.Error(t, err)
	})
}

func TestServiceGetAll(t *testing.T) {
	mockSensor := []domain.Sensor{
		{
			Name: "sensor",
		},
	}
	t.Run("Success", func(t *testing.T) {
		mockSensorRepo := new(mocks.SensorRepositoryInterface)
		mockSensorRepo.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(mockSensor, nil)
		s := sensor.NewService(mockSensorRepo)
		data, err := s.GetAll(context.Background(), 1, 10)
		assert.Len(t, data, len(mockSensor))
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockSensorRepo := new(mocks.SensorRepositoryInterface)
		mockSensorRepo.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
		s := sensor.NewService(mockSensorRepo)
		data, err := s.GetAll(context.Background(), 1, 10)
		assert.Nil(t, data)
		assert.Error(t, err)
	})
}

func TestServiceUpdate(t *testing.T) {
	mockSensor := domain.Sensor{
		Name: "sensor",
	}
	t.Run("Success", func(t *testing.T) {
		mockSensorRepo := new(mocks.SensorRepositoryInterface)
		mockSensorRepo.On("Update", mock.Anything, mock.Anything).Return(nil)
		s := sensor.NewService(mockSensorRepo)
		err := s.Update(context.Background(), &mockSensor)
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockSensorRepo := new(mocks.SensorRepositoryInterface)
		mockSensorRepo.On("Update", mock.Anything, mock.Anything).Return(errors.New("error")).Once()
		s := sensor.NewService(mockSensorRepo)
		err := s.Update(context.Background(), &mockSensor)
		assert.Error(t, err)
	})
}

func TestServiceGetByID(t *testing.T) {
	mockSensor := domain.Sensor{
		Name: "sensor",
	}
	t.Run("Success", func(t *testing.T) {
		mockSensorRepo := new(mocks.SensorRepositoryInterface)
		mockSensorRepo.On("GetByID", mock.Anything, mock.Anything, mock.Anything).Return(&mockSensor, nil)
		s := sensor.NewService(mockSensorRepo)
		data, err := s.GetByID(context.Background(), "1")
		assert.Equal(t, data.Name, mockSensor.Name)
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockSensorRepo := new(mocks.SensorRepositoryInterface)
		mockSensorRepo.On("GetByID", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
		s := sensor.NewService(mockSensorRepo)
		data, err := s.GetByID(context.Background(), "1")
		assert.Nil(t, data)
		assert.Error(t, err)
	})
}

func TestServiceDelete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockSensorRepo := new(mocks.SensorRepositoryInterface)
		mockSensorRepo.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		s := sensor.NewService(mockSensorRepo)
		err := s.Delete(context.Background(), "1")
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockSensorRepo := new(mocks.SensorRepositoryInterface)
		mockSensorRepo.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()
		s := sensor.NewService(mockSensorRepo)
		err := s.Delete(context.Background(), "1")
		assert.Error(t, err)
	})
}
