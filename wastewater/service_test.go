package wastewater_test

import (
	"context"
	"errors"
	"testing"

	"github.com/anggi-susanto/mrt-go/domain"
	"github.com/anggi-susanto/mrt-go/wastewater"
	"github.com/anggi-susanto/mrt-go/wastewater/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceCreate(t *testing.T) {
	mockWasteWater := domain.WastewaterDataRequest{
		BOD: 10,
	}
	t.Run("Success", func(t *testing.T) {
		mockWasteWaterRepo := new(mocks.WasteWaterRepositoryInterface)
		mockWasteWaterRepo.On("Create", mock.Anything, mock.Anything).Return(nil)
		s := wastewater.NewService(mockWasteWaterRepo)
		err := s.Create(context.Background(), &mockWasteWater)
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockWasteWaterRepo := new(mocks.WasteWaterRepositoryInterface)
		mockWasteWaterRepo.On("Create", mock.Anything, mock.Anything).Return(errors.New("error")).Once()
		s := wastewater.NewService(mockWasteWaterRepo)
		err := s.Create(context.Background(), &mockWasteWater)
		assert.Error(t, err)
	})
}

func TestServiceGetAll(t *testing.T) {
	mockWasteWater := []domain.WasteWaterData{
		{
			BOD: 10,
		},
	}
	t.Run("Success", func(t *testing.T) {
		mockWasteWaterRepo := new(mocks.WasteWaterRepositoryInterface)
		mockWasteWaterRepo.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(mockWasteWater, nil)
		s := wastewater.NewService(mockWasteWaterRepo)
		data, err := s.GetAll(context.Background(), 1, 10)
		assert.Len(t, data, len(mockWasteWater))
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockWasteWaterRepo := new(mocks.WasteWaterRepositoryInterface)
		mockWasteWaterRepo.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
		s := wastewater.NewService(mockWasteWaterRepo)
		data, err := s.GetAll(context.Background(), 1, 10)
		assert.Nil(t, data)
		assert.Error(t, err)
	})
}

func TestServiceUpdate(t *testing.T) {
	mockWasteWater := domain.WasteWaterData{
		BOD: 10,
	}
	t.Run("Success", func(t *testing.T) {
		mockWasteWaterRepo := new(mocks.WasteWaterRepositoryInterface)
		mockWasteWaterRepo.On("Update", mock.Anything, mock.Anything).Return(nil)
		s := wastewater.NewService(mockWasteWaterRepo)
		err := s.Update(context.Background(), &mockWasteWater)
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockWasteWaterRepo := new(mocks.WasteWaterRepositoryInterface)
		mockWasteWaterRepo.On("Update", mock.Anything, mock.Anything).Return(errors.New("error")).Once()
		s := wastewater.NewService(mockWasteWaterRepo)
		err := s.Update(context.Background(), &mockWasteWater)
		assert.Error(t, err)
	})
}

func TestServiceGetByID(t *testing.T) {
	mockWasteWater := domain.WasteWaterData{
		BOD: 10,
	}
	t.Run("Success", func(t *testing.T) {
		mockWasteWaterRepo := new(mocks.WasteWaterRepositoryInterface)
		mockWasteWaterRepo.On("GetByID", mock.Anything, mock.Anything, mock.Anything).Return(&mockWasteWater, nil)
		s := wastewater.NewService(mockWasteWaterRepo)
		data, err := s.GetByID(context.Background(), "1")
		assert.Equal(t, data.BOD, mockWasteWater.BOD)
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockWasteWaterRepo := new(mocks.WasteWaterRepositoryInterface)
		mockWasteWaterRepo.On("GetByID", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
		s := wastewater.NewService(mockWasteWaterRepo)
		data, err := s.GetByID(context.Background(), "1")
		assert.Nil(t, data)
		assert.Error(t, err)
	})
}

func TestServiceDelete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockWasteWaterRepo := new(mocks.WasteWaterRepositoryInterface)
		mockWasteWaterRepo.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		s := wastewater.NewService(mockWasteWaterRepo)
		err := s.Delete(context.Background(), "1")
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockWasteWaterRepo := new(mocks.WasteWaterRepositoryInterface)
		mockWasteWaterRepo.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()
		s := wastewater.NewService(mockWasteWaterRepo)
		err := s.Delete(context.Background(), "1")
		assert.Error(t, err)
	})
}
