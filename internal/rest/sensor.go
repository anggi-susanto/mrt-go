package rest

import (
	"context"

	"github.com/anggi-susanto/mrt-go/domain"
	"github.com/gofiber/fiber/v2"
)

// ResponseError represents an error response

// SensorService is the interface that wraps the Create, GetAll, GetByID, Update, and Delete methods.

type SensorService interface {
	Create(ctx context.Context, w *domain.SensorRequest) error
	GetAll(ctx context.Context, page int, limit int) ([]domain.Sensor, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, w *domain.Sensor) error
	GetByID(ctx context.Context, id string) (*domain.Sensor, error)
}

// SensorHandler is the handler for SensorService
type SensorHandler struct {
	service SensorService
}

// SensorIDEndpoint is the endpoint for SensorService
const SensorIDEndpoint = "/sensor/:id"

// NewSensorHandler initializes a new SensorHandler with the provided Fiber app and SensorService.
//
// Parameters:
// - app: The Fiber app instance.
// - service: The SensorService instance.
//
// Return type: None.
func NewSensorHandler(app *fiber.App, service SensorService) {
	handler := &SensorHandler{service: service}
	app.Post("/sensor", handler.Create)
	app.Get("/sensor", handler.GetAll)
	app.Get(SensorIDEndpoint, handler.GetByID)
	app.Put(SensorIDEndpoint, handler.Update)
	app.Delete(SensorIDEndpoint, handler.Delete)
}

// Create handles the creation of sensor data.
//
// @Summary create sensor data
// @Description create sensor data
// @Tags sensor
// @Accept json
// @Produce json
// @Param waste_water body domain.Sensor true "sensor data"
// @Success 201 {object} domain.Sensor
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /sensor [post]
func (h *SensorHandler) Create(ctx *fiber.Ctx) error {
	w := &domain.SensorRequest{}
	if err := ctx.BodyParser(w); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ResponseError{Message: err.Error()})
	}
	if err := h.service.Create(ctx.Context(), w); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(w)
}

// GetAll retrieves all sensor data.
//
// It takes a fiber context as a parameter and returns an error.
// @Summary get all sensor data
// @Description get all sensor data
// @Tags sensor
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Success 200 {array} domain.Sensor "Sensor data"
// @Failure 500 {object} ResponseError
// @Router /sensor [get]
func (h *SensorHandler) GetAll(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
	wastes, err := h.service.GetAll(ctx.Context(), page, limit)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: "Failed to get all sensor data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(wastes)
}

// GetByID retrieves a WasteWater object by ID.
//
// ctx *fiber.Ctx - Context object containing the request information.
// error - Returns an error if one occurs.
// @Summary get sensor data by id
// @Description get sensor data by id
// @Tags sensor
// @Accept json
// @Produce json
// @Param id path string true "Sensor data ID"
// @Success 200 {object} domain.Sensor
// @Failure 500 {object} ResponseError
// @Router /sensor/{id} [get]
// @Failure 404 {object} ResponseError
func (h *SensorHandler) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	w, err := h.service.GetByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: err.Error()})
	}

	if w == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(ResponseError{Message: "Data not found"})
	}
	return ctx.Status(fiber.StatusOK).JSON(w)
}

// Update updates the SensorHandler.
//
// It takes a fiber context as a parameter and returns an error.
// @Summary update sensor data
// @Description update sensor data
// @Tags sensor
// @Accept json
// @Produce json
// @Param waste_water body domain.Sensor true "sensor data"
// @Success 200 {object} domain.Sensor
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /sensor/{id} [put]
// @Failure 404 {object} ResponseError
func (h *SensorHandler) Update(ctx *fiber.Ctx) error {
	w := &domain.Sensor{}
	if err := ctx.BodyParser(w); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ResponseError{Message: err.Error()})
	}
	if err := h.service.Update(ctx.Context(), w); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(w)
}

// Delete deletes a SensorHandler item.
//
// ctx *fiber.Ctx parameter. Returns an error.
// @Summary delete sensor data
// @Description delete sensor data
// @Tags sensor
// @Accept json
// @Produce json
// @Param id path string true "Sensor data ID"
// @Success 204
// @Failure 500 {object} ResponseError
// @Router /sensor/{id} [delete]
// @Failure 404 {object} ResponseError
func (h *SensorHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := h.service.Delete(ctx.Context(), id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: err.Error()})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
