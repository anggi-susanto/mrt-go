package rest

import (
	"context"

	"github.com/anggi-susanto/mrt-go/domain"
	"github.com/gofiber/fiber/v2"
)

// ResponseError represents an error response

// DeviceService is the interface that wraps the Create, GetAll, GetByID, Update, and Delete methods.

type DeviceService interface {
	Create(ctx context.Context, w *domain.DeviceRequest) error
	GetAll(ctx context.Context, page int, limit int) ([]domain.Device, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, w *domain.Device) error
	GetByID(ctx context.Context, id string) (*domain.Device, error)
}

// DeviceHandler is the handler for DeviceService
type DeviceHandler struct {
	service DeviceService
}

// DeviceIDEndpoint is the endpoint for DeviceService
const DeviceIDEndpoint = "/device/:id"

// NewDeviceHandler initializes a new DeviceHandler with the provided Fiber app and DeviceService.
//
// Parameters:
// - app: The Fiber app instance.
// - service: The DeviceService instance.
//
// Return type: None.
func NewDeviceHandler(app *fiber.App, service DeviceService) {
	handler := &DeviceHandler{service: service}
	app.Post("/device", handler.Create)
	app.Get("/device", handler.GetAll)
	app.Get(DeviceIDEndpoint, handler.GetByID)
	app.Put(DeviceIDEndpoint, handler.Update)
	app.Delete(DeviceIDEndpoint, handler.Delete)
}

// Create handles the creation of device data.
//
// @Summary create device data
// @Description create device data
// @Tags device
// @Accept json
// @Produce json
// @Param waste_water body domain.Device true "device data"
// @Success 201 {object} domain.Device
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /device [post]
func (h *DeviceHandler) Create(ctx *fiber.Ctx) error {
	w := &domain.DeviceRequest{}
	if err := ctx.BodyParser(w); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ResponseError{Message: err.Error()})
	}
	if err := h.service.Create(ctx.Context(), w); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(w)
}

// GetAll retrieves all device data.
//
// It takes a fiber context as a parameter and returns an error.
// @Summary get all device data
// @Description get all device data
// @Tags device
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Success 200 {array} domain.Device "Device data"
// @Failure 500 {object} ResponseError
// @Router /device [get]
func (h *DeviceHandler) GetAll(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
	wastes, err := h.service.GetAll(ctx.Context(), page, limit)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: "Failed to get all device data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(wastes)
}

// GetByID retrieves a WasteWater object by ID.
//
// ctx *fiber.Ctx - Context object containing the request information.
// error - Returns an error if one occurs.
// @Summary get device data by id
// @Description get device data by id
// @Tags device
// @Accept json
// @Produce json
// @Param id path string true "Device data ID"
// @Success 200 {object} domain.Device
// @Failure 500 {object} ResponseError
// @Router /device/{id} [get]
// @Failure 404 {object} ResponseError
func (h *DeviceHandler) GetByID(ctx *fiber.Ctx) error {
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

// Update updates the DeviceHandler.
//
// It takes a fiber context as a parameter and returns an error.
// @Summary update device data
// @Description update device data
// @Tags device
// @Accept json
// @Produce json
// @Param waste_water body domain.Device true "device data"
// @Success 200 {object} domain.Device
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /device/{id} [put]
// @Failure 404 {object} ResponseError
func (h *DeviceHandler) Update(ctx *fiber.Ctx) error {
	w := &domain.Device{}
	if err := ctx.BodyParser(w); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ResponseError{Message: err.Error()})
	}
	if err := h.service.Update(ctx.Context(), w); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(w)
}

// Delete deletes a DeviceHandler item.
//
// ctx *fiber.Ctx parameter. Returns an error.
// @Summary delete device data
// @Description delete device data
// @Tags device
// @Accept json
// @Produce json
// @Param id path string true "Device data ID"
// @Success 204
// @Failure 500 {object} ResponseError
// @Router /device/{id} [delete]
// @Failure 404 {object} ResponseError
func (h *DeviceHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := h.service.Delete(ctx.Context(), id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: err.Error()})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
