package rest_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/anggi-susanto/mrt-go/domain"
	"github.com/anggi-susanto/mrt-go/internal/rest"
	"github.com/anggi-susanto/mrt-go/internal/rest/mocks"
)

const deviceEnpoint = "/device"

func TestCreateDeviceHandlerSuccess(t *testing.T) {
	app := fiber.New()
	mockService := new(mocks.DeviceService)
	rest.NewDeviceHandler(app, mockService)

	mockService.On("Create", mock.Anything, mock.Anything).Return(nil)

	// Create a mock request body
	device := domain.DeviceRequest{
		Name: "Device",
	}
	body, _ := json.Marshal(device)

	req := httptest.NewRequest(http.MethodPost, deviceEnpoint, bytes.NewReader(body))
	req.Header.Set(contentType, applicationJson)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	respData := domain.Device{}
	_ = json.Unmarshal(data, &respData)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	assert.Equal(t, respData.Name, device.Name)
}

func TestCreateDeviceHandlerErrorParsingRequestBody(t *testing.T) {
	app := fiber.New()
	mockService := new(mocks.DeviceService)
	rest.NewDeviceHandler(app, mockService)

	mockService.On("Create", mock.Anything, mock.Anything).Return(nil)

	// Create a request with invalid JSON body
	req := httptest.NewRequest(http.MethodPost, deviceEnpoint, bytes.NewReader([]byte("{")))
	req.Header.Set(contentType, applicationJson)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, "{\"message\":\"unexpected end of JSON input\"}", string(data))
}

func TestCreateDeviceHandlerErrorCreatingData(t *testing.T) {
	app := fiber.New()
	mockService := new(mocks.DeviceService)
	rest.NewDeviceHandler(app, mockService)

	// Create a mock request body
	device := domain.Device{
		// Your mock data here
	}
	body, _ := json.Marshal(device)
	mockService.On("Create", mock.Anything, mock.Anything).Return(errors.New("error"))

	req := httptest.NewRequest(http.MethodPost, deviceEnpoint, bytes.NewReader(body))
	req.Header.Set(contentType, applicationJson)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, "{\"message\":\"error\"}", string(data))
}

func TestDeviceHandlerGetAll(t *testing.T) {

	device := []domain.Device{
		{
			Name: "device",
		},
	}
	t.Run("Success with default values", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.DeviceService) // Implement a mock service for testing purposes
		rest.NewDeviceHandler(app, mockService)
		mockService.On("GetAll", mock.Anything, 1, 10).Return(device, nil)
		req := httptest.NewRequest(http.MethodGet, "/device", nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response []domain.Device
		data, _ := io.ReadAll(resp.Body)

		err = json.Unmarshal(data, &response)
		assert.Nil(t, err)
		assert.Equal(t, response[0].Name, device[0].Name)
	})

	t.Run("Error case", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.DeviceService) // Implement a mock service for testing purposes
		rest.NewDeviceHandler(app, mockService)
		mockService.On("GetAll", mock.Anything, 1, 10).Return(nil, errors.New("error"))
		req := httptest.NewRequest(http.MethodGet, "/device", nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		data, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "{\"message\":\"Failed to get all device data\"}", string(data))
	})

}

func TestDeviceHandlerGetByID(t *testing.T) {
	device := domain.Device{
		ID:   primitive.NewObjectID(),
		Name: "device",
	}
	// Test for retrieving device data by a valid ID
	t.Run("Valid ID", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.DeviceService) // Implement a mock service for testing purposes
		rest.NewDeviceHandler(app, mockService)
		mockService.On("GetByID", mock.Anything, device.ID.String()).Return(&device, nil)
		req := httptest.NewRequest(http.MethodGet, "/device/"+device.ID.String(), nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response domain.Device
		data, _ := io.ReadAll(resp.Body)

		err = json.Unmarshal(data, &response)
		assert.Nil(t, err)
		assert.Equal(t, response.Name, device.Name)
	})

	// Test for retrieving device data by an invalid ID
	t.Run("Invalid ID", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.DeviceService) // Implement a mock service for testing purposes
		rest.NewDeviceHandler(app, mockService)
		mockService.On("GetByID", mock.Anything, device.ID.String()).Return(nil, nil)
		req := httptest.NewRequest(http.MethodGet, "/device/"+device.ID.String(), nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		var response domain.Device
		data, _ := io.ReadAll(resp.Body)

		err = json.Unmarshal(data, &response)
		assert.Nil(t, err)
		assert.NotEqual(t, response.Name, device.Name)
	})

	// Test for handling an error while retrieving device data by ID
	t.Run("Error handling", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.DeviceService) // Implement a mock service for testing purposes
		rest.NewDeviceHandler(app, mockService)
		mockService.On("GetByID", mock.Anything, device.ID.String()).Return(nil, errors.New("error"))
		req := httptest.NewRequest(http.MethodGet, "/device/"+device.ID.String(), nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		data, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "{\"message\":\"error\"}", string(data))
	})
}
