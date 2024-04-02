package rest

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.org/anggi-susanto/mrt-go/domain"
)

// ResponseError represents an error response
type ResponseError struct {
	Message string `json:"message"`
}

// WasteWaterServices is the interface that wraps the Create, GetAll, GetByID, Update, and Delete methods.
type WasteWaterServices interface {
	Create(ctx context.Context, w *domain.WasteWaterData) error
	GetAll(ctx context.Context, page, limit int) ([]domain.WasteWaterData, error)
	GetByID(ctx context.Context, id string) (*domain.WasteWaterData, error)
	Update(ctx context.Context, w *domain.WasteWaterData) error
	Delete(ctx context.Context, id string) error
}

// WasteWaterHandler is the handler for WasteWaterServices
type WasteWaterHandler struct {
	service WasteWaterServices
}

// WasteWaterIDEndpoint is the endpoint for WasteWaterServices
const WasteWaterIDEndpoint = "/waste-water/:id"

// NewWasteWaterHandler initializes a new WasteWaterHandler with the provided Fiber app and WasteWaterServices.
//
// Parameters:
// - app: The Fiber app instance.
// - service: The WasteWaterServices instance.
//
// Return type: None.
func NewWasteWaterHandler(app *fiber.App, service WasteWaterServices) {
	handler := &WasteWaterHandler{service: service}
	app.Post("/waste-water", handler.Create)
	app.Get("/waste-water", handler.GetAll)
	app.Get(WasteWaterIDEndpoint, handler.GetByID)
	app.Put(WasteWaterIDEndpoint, handler.Update)
	app.Delete(WasteWaterIDEndpoint, handler.Delete)
}

// Create handles the creation of waste water data.
//
// @Summary create waste water data
// @Description create waste water data
// @Tags waste water
// @Accept json
// @Produce json
// @Param waste_water body domain.WasteWaterData true "waste water data"
// @Success 201 {object} domain.WasteWaterData
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /waste-water [post]
func (h *WasteWaterHandler) Create(ctx *fiber.Ctx) error {
	w := &domain.WasteWaterData{}
	if err := ctx.BodyParser(w); err != nil {
		logrus.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(ResponseError{Message: err.Error()})
	}
	if err := h.service.Create(ctx.Context(), w); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(w)
}

// GetAll retrieves all waste water data.
//
// It takes a fiber context as a parameter and returns an error.
// @Summary get all waste water data
// @Description get all waste water data
// @Tags waste water
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Success 200 {array} domain.WasteWaterData "Waste water data"
// @Failure 500 {object} ResponseError
// @Router /waste-water [get]
func (h *WasteWaterHandler) GetAll(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
	wastes, err := h.service.GetAll(ctx.Context(), page, limit)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: "Failed to get all waste water data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(wastes)
}

// GetByID retrieves a WasteWater object by ID.
//
// ctx *fiber.Ctx - Context object containing the request information.
// error - Returns an error if one occurs.
// @Summary get waste water data by id
// @Description get waste water data by id
// @Tags waste water
// @Accept json
// @Produce json
// @Param id path string true "Waste water data ID"
// @Success 200 {object} domain.WasteWaterData
// @Failure 500 {object} ResponseError
// @Router /waste-water/{id} [get]
// @Failure 404 {object} ResponseError
func (h *WasteWaterHandler) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	w, err := h.service.GetByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(w)
}

// Update updates the WasteWaterHandler.
//
// It takes a fiber context as a parameter and returns an error.
// @Summary update waste water data
// @Description update waste water data
// @Tags waste water
// @Accept json
// @Produce json
// @Param waste_water body domain.WasteWaterData true "waste water data"
// @Success 200 {object} domain.WasteWaterData
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /waste-water/{id} [put]
// @Failure 404 {object} ResponseError
func (h *WasteWaterHandler) Update(ctx *fiber.Ctx) error {
	w := &domain.WasteWaterData{}
	if err := ctx.BodyParser(w); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ResponseError{Message: err.Error()})
	}
	if err := h.service.Update(ctx.Context(), w); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(w)
}

// Delete deletes a WasteWaterHandler item.
//
// ctx *fiber.Ctx parameter. Returns an error.
// @Summary delete waste water data
// @Description delete waste water data
// @Tags waste water
// @Accept json
// @Produce json
// @Param id path string true "Waste water data ID"
// @Success 204
// @Failure 500 {object} ResponseError
// @Router /waste-water/{id} [delete]
// @Failure 404 {object} ResponseError
func (h *WasteWaterHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := h.service.Delete(ctx.Context(), id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: err.Error()})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
