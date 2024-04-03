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

const wasteWaterEnpoint = "/waste-water"
const contentType = "Content-Type"
const applicationJson = "application/json"

func TestCreateWasteWaterHandlerSuccess(t *testing.T) {
	app := fiber.New()
	mockService := new(mocks.WasteWaterServices)
	rest.NewWasteWaterHandler(app, mockService)

	mockService.On("Create", mock.Anything, mock.Anything).Return(nil)

	// Create a mock request body
	waterData := domain.WastewaterDataRequest{
		BOD: 10,
	}
	body, _ := json.Marshal(waterData)

	req := httptest.NewRequest(http.MethodPost, wasteWaterEnpoint, bytes.NewReader(body))
	req.Header.Set(contentType, applicationJson)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	respData := domain.WasteWaterData{}
	_ = json.Unmarshal(data, &respData)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	assert.Equal(t, respData.BOD, waterData.BOD)
}

func TestCreateWasteWaterHandlerErrorParsingRequestBody(t *testing.T) {
	app := fiber.New()
	mockService := new(mocks.WasteWaterServices)
	rest.NewWasteWaterHandler(app, mockService)

	mockService.On("Create", mock.Anything, mock.Anything).Return(nil)

	// Create a request with invalid JSON body
	req := httptest.NewRequest(http.MethodPost, wasteWaterEnpoint, bytes.NewReader([]byte("{")))
	req.Header.Set(contentType, applicationJson)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, "{\"message\":\"unexpected end of JSON input\"}", string(data))
}

func TestCreateWasteWaterHandlerErrorCreatingData(t *testing.T) {
	app := fiber.New()
	mockService := new(mocks.WasteWaterServices)
	rest.NewWasteWaterHandler(app, mockService)

	// Create a mock request body
	waterData := domain.WasteWaterData{
		// Your mock data here
	}
	body, _ := json.Marshal(waterData)
	mockService.On("Create", mock.Anything, mock.Anything).Return(errors.New("error"))

	req := httptest.NewRequest(http.MethodPost, wasteWaterEnpoint, bytes.NewReader(body))
	req.Header.Set(contentType, applicationJson)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, "{\"message\":\"error\"}", string(data))
}

func TestWasteWaterHandlerGetAll(t *testing.T) {

	waterData := []domain.WasteWaterData{
		{
			BOD: 10,
		},
	}
	t.Run("Success with default values", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.WasteWaterServices) // Implement a mock service for testing purposes
		rest.NewWasteWaterHandler(app, mockService)
		mockService.On("GetAll", mock.Anything, 1, 10).Return(waterData, nil)
		req := httptest.NewRequest(http.MethodGet, "/waste-water", nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response []domain.WasteWaterData
		data, _ := io.ReadAll(resp.Body)

		err = json.Unmarshal(data, &response)
		assert.Nil(t, err)
		assert.Equal(t, response[0].BOD, waterData[0].BOD)
	})

	t.Run("Error case", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.WasteWaterServices) // Implement a mock service for testing purposes
		rest.NewWasteWaterHandler(app, mockService)
		mockService.On("GetAll", mock.Anything, 1, 10).Return(nil, errors.New("error"))
		req := httptest.NewRequest(http.MethodGet, "/waste-water", nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		data, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "{\"message\":\"Failed to get all waste water data\"}", string(data))
	})

}

func TestWasteWaterHandlerGetByID(t *testing.T) {
	waterData := domain.WasteWaterData{
		ID:  primitive.NewObjectID(),
		BOD: 10,
	}
	// Test for retrieving waste water data by a valid ID
	t.Run("Valid ID", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.WasteWaterServices) // Implement a mock service for testing purposes
		rest.NewWasteWaterHandler(app, mockService)
		mockService.On("GetByID", mock.Anything, waterData.ID.String()).Return(&waterData, nil)
		req := httptest.NewRequest(http.MethodGet, "/waste-water/"+waterData.ID.String(), nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response domain.WasteWaterData
		data, _ := io.ReadAll(resp.Body)

		err = json.Unmarshal(data, &response)
		assert.Nil(t, err)
		assert.Equal(t, response.BOD, waterData.BOD)
	})

	// Test for retrieving waste water data by an invalid ID
	t.Run("Invalid ID", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.WasteWaterServices) // Implement a mock service for testing purposes
		rest.NewWasteWaterHandler(app, mockService)
		mockService.On("GetByID", mock.Anything, waterData.ID.String()).Return(nil, nil)
		req := httptest.NewRequest(http.MethodGet, "/waste-water/"+waterData.ID.String(), nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		var response domain.WasteWaterData
		data, _ := io.ReadAll(resp.Body)

		err = json.Unmarshal(data, &response)
		assert.Nil(t, err)
		assert.NotEqual(t, response.BOD, waterData.BOD)
	})

	// Test for handling an error while retrieving waste water data by ID
	t.Run("Error handling", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mocks.WasteWaterServices) // Implement a mock service for testing purposes
		rest.NewWasteWaterHandler(app, mockService)
		mockService.On("GetByID", mock.Anything, waterData.ID.String()).Return(nil, errors.New("error"))
		req := httptest.NewRequest(http.MethodGet, "/waste-water/"+waterData.ID.String(), nil)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		data, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "{\"message\":\"error\"}", string(data))
	})
}
