package device_test

import (
	"context"
	"errors"
	"testing"

	"github.com/anggi-susanto/mrt-go/device"
	"github.com/anggi-susanto/mrt-go/device/mocks"
	"github.com/anggi-susanto/mrt-go/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceCreate(t *testing.T) {
	mockDevice := domain.DeviceRequest{
		Name: "device",
	}
	t.Run("Success", func(t *testing.T) {
		mockDeviceRepo := new(mocks.DeviceRepositoryInterface)
		mockDeviceRepo.On("Create", mock.Anything, mock.Anything).Return(nil)
		s := device.NewService(mockDeviceRepo)
		err := s.Create(context.Background(), &mockDevice)
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockDeviceRepo := new(mocks.DeviceRepositoryInterface)
		mockDeviceRepo.On("Create", mock.Anything, mock.Anything).Return(errors.New("error")).Once()
		s := device.NewService(mockDeviceRepo)
		err := s.Create(context.Background(), &mockDevice)
		assert.Error(t, err)
	})
}

func TestServiceGetAll(t *testing.T) {
	mockDevice := []domain.Device{
		{
			Name: "device",
		},
	}
	t.Run("Success", func(t *testing.T) {
		mockDeviceRepo := new(mocks.DeviceRepositoryInterface)
		mockDeviceRepo.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(mockDevice, nil)
		s := device.NewService(mockDeviceRepo)
		data, err := s.GetAll(context.Background(), 1, 10)
		assert.Len(t, data, len(mockDevice))
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockDeviceRepo := new(mocks.DeviceRepositoryInterface)
		mockDeviceRepo.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
		s := device.NewService(mockDeviceRepo)
		data, err := s.GetAll(context.Background(), 1, 10)
		assert.Nil(t, data)
		assert.Error(t, err)
	})
}

func TestServiceUpdate(t *testing.T) {
	mockDevice := domain.Device{
		Name: "device",
	}
	t.Run("Success", func(t *testing.T) {
		mockDeviceRepo := new(mocks.DeviceRepositoryInterface)
		mockDeviceRepo.On("Update", mock.Anything, mock.Anything).Return(nil)
		s := device.NewService(mockDeviceRepo)
		err := s.Update(context.Background(), &mockDevice)
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockDeviceRepo := new(mocks.DeviceRepositoryInterface)
		mockDeviceRepo.On("Update", mock.Anything, mock.Anything).Return(errors.New("error")).Once()
		s := device.NewService(mockDeviceRepo)
		err := s.Update(context.Background(), &mockDevice)
		assert.Error(t, err)
	})
}

func TestServiceGetByID(t *testing.T) {
	mockDevice := domain.Device{
		Name: "device",
	}
	t.Run("Success", func(t *testing.T) {
		mockDeviceRepo := new(mocks.DeviceRepositoryInterface)
		mockDeviceRepo.On("GetByID", mock.Anything, mock.Anything, mock.Anything).Return(&mockDevice, nil)
		s := device.NewService(mockDeviceRepo)
		data, err := s.GetByID(context.Background(), "1")
		assert.Equal(t, data.Name, mockDevice.Name)
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockDeviceRepo := new(mocks.DeviceRepositoryInterface)
		mockDeviceRepo.On("GetByID", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
		s := device.NewService(mockDeviceRepo)
		data, err := s.GetByID(context.Background(), "1")
		assert.Nil(t, data)
		assert.Error(t, err)
	})
}

func TestServiceDelete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockDeviceRepo := new(mocks.DeviceRepositoryInterface)
		mockDeviceRepo.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		s := device.NewService(mockDeviceRepo)
		err := s.Delete(context.Background(), "1")
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockDeviceRepo := new(mocks.DeviceRepositoryInterface)
		mockDeviceRepo.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()
		s := device.NewService(mockDeviceRepo)
		err := s.Delete(context.Background(), "1")
		assert.Error(t, err)
	})
}
