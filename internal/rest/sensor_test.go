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

const sensorEnpoint = "/sensor"

func TestCreateSensorHandlerSuccess(t *testing.T) {
	app := fiber.New()
	mockService := new(mocks.SensorService)
	rest.NewSensorHandler(app, mockService)

	mockService.On("Create", mock.Anything, mock.Anything).Return(nil)

	// Create a mock request body
	sensor := domain.SensorRequest{
		Name: "Sensor",
	}
	body, _ := json.Marshal(sensor)

	req := httptest.NewRequest(http.MethodPost, sensorEnpoint, bytes.NewReader(body))
	req.Header.Set(contentType, applicationJson)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	respData := domain.Sensor{}
	_ = json.Unmarshal(data, &respData)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	assert.Equal(t, respData.Name, sensor.Name)
}

func TestCreateSensorHandlerErrorParsingRequestBody(t *testing.T) {
	app := fiber.New()
	mockService := new(mocks.SensorService)
	rest.NewSensorHandler(app, mockService)

	mockService.On("Create", mock.Anything, mock.Anything).Return(nil)

	// Create a request with invalid JSON body
	req := httptest.NewRequest(http.MethodPost, sensorEnpoint, bytes.NewReader([]byte("{")))
	req.Header.Set(contentType, applicationJson)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, "{\"message\":\"unexpected end of JSON input\"}", string(data))
}

func TestCreateSensorHandlerErrorCreatingData(t *testing.T) {
	app := fiber.New()
	mockService := new(mocks.SensorService)
	rest.NewSensorHandler(app, mockService)

	// Create a mock request body
	sensor := domain.Sensor{
		// Your mock data here
	}
	body, _ := json.Marshal(sensor)
	mockService.On("Create", mock.Anything, mock.Anything).Return(errors.New("error"))

	req := httptest.NewRequest(http.MethodPost, sensorEnpoint, bytes.NewReader(body))
	req.Header.Set(contentType, applicationJson)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, "{\"message\":\"error\"}", string(data))
}

func TestSensorHandlerGetAll(t *testing.T) {

	sensor := []domain.Sensor{
		{
			Name: "sensor",
		},
	}
	t.Run("Success with default values", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.SensorService) // Implement a mock service for testing purposes
		rest.NewSensorHandler(app, mockService)
		mockService.On("GetAll", mock.Anything, 1, 10).Return(sensor, nil)
		req := httptest.NewRequest(http.MethodGet, "/sensor", nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response []domain.Sensor
		data, _ := io.ReadAll(resp.Body)

		err = json.Unmarshal(data, &response)
		assert.Nil(t, err)
		assert.Equal(t, response[0].Name, sensor[0].Name)
	})

	t.Run("Error case", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.SensorService) // Implement a mock service for testing purposes
		rest.NewSensorHandler(app, mockService)
		mockService.On("GetAll", mock.Anything, 1, 10).Return(nil, errors.New("error"))
		req := httptest.NewRequest(http.MethodGet, "/sensor", nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		data, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "{\"message\":\"Failed to get all sensor data\"}", string(data))
	})

}

func TestSensorHandlerGetByID(t *testing.T) {
	sensor := domain.Sensor{
		ID:   primitive.NewObjectID(),
		Name: "sensor",
	}
	// Test for retrieving sensor data by a valid ID
	t.Run("Valid ID", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.SensorService) // Implement a mock service for testing purposes
		rest.NewSensorHandler(app, mockService)
		mockService.On("GetByID", mock.Anything, sensor.ID.String()).Return(&sensor, nil)
		req := httptest.NewRequest(http.MethodGet, "/sensor/"+sensor.ID.String(), nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response domain.Sensor
		data, _ := io.ReadAll(resp.Body)

		err = json.Unmarshal(data, &response)
		assert.Nil(t, err)
		assert.Equal(t, response.Name, sensor.Name)
	})

	// Test for retrieving sensor data by an invalid ID
	t.Run("Invalid ID", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.SensorService) // Implement a mock service for testing purposes
		rest.NewSensorHandler(app, mockService)
		mockService.On("GetByID", mock.Anything, sensor.ID.String()).Return(nil, nil)
		req := httptest.NewRequest(http.MethodGet, "/sensor/"+sensor.ID.String(), nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		var response domain.Sensor
		data, _ := io.ReadAll(resp.Body)

		err = json.Unmarshal(data, &response)
		assert.Nil(t, err)
		assert.NotEqual(t, response.Name, sensor.Name)
	})

	// Test for handling an error while retrieving sensor data by ID
	t.Run("Error handling", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.SensorService) // Implement a mock service for testing purposes
		rest.NewSensorHandler(app, mockService)
		mockService.On("GetByID", mock.Anything, sensor.ID.String()).Return(nil, errors.New("error"))
		req := httptest.NewRequest(http.MethodGet, "/sensor/"+sensor.ID.String(), nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		data, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "{\"message\":\"error\"}", string(data))
	})
}
