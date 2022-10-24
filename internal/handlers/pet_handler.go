package handlers

import (
	"net/http"

	"github.com/internal/model"
	"github.com/internal/shared"

	"github.com/gofiber/fiber/v2"
)

type IPetHandler interface {
	GetPetByID(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

type PetHandler struct {
	Handler
}

func NewPetHandler() *PetHandler {
	return &PetHandler{}
}

// GetPetByID godoc.
func (p PetHandler) GetPetByID(ctx *fiber.Ctx) error {
	model := new(model.PetModel)
	model.ID = int64(1)
	model.Name = "Rin Tin Tin"

	return p.Handler.
		SendJSON(ctx, model)
}

// GetAll godoc.
func (p PetHandler) GetAll(ctx *fiber.Ctx) error {
	model := new([]model.PetModel)

	if len(*model) == 0 {
		return shared.NewError(http.StatusNotFound, "no pets found")
	}

	return p.Handler.
		SendJSON(ctx, model)
}

// Create godoc.
func (p PetHandler) Create(ctx *fiber.Ctx) error {
	var model model.PetModel

	var b = int64(0)
	model.ID = int64(0) / b

	return p.Handler.
		SendJSON(ctx, model)
}
