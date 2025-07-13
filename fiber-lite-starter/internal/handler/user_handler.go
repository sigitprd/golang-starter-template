package handler

import (
	"fiber-lite-starter/internal/dto"
	"fiber-lite-starter/internal/service"
	"fiber-lite-starter/pkg/errmsg"
	"fiber-lite-starter/pkg/response"
	"fiber-lite-starter/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"net/http"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.UserRequest

	if err := c.BodyParser(&req); err != nil {
		log.Info().Err(err).Msg("handler::CreateUser - Failed to parse request body")
		code, errs := errmsg.Errors(err, &req)
		return c.Status(code).JSON(response.Error(errs))
	}

	if err := c.Locals("validator").(func(interface{}) error)(&req); err != nil {
		log.Info().Err(err).Msg("handler::CreateUser - Validation failed")
		code, errs := errmsg.Errors(err, &req)
		return c.Status(code).JSON(response.Error(errs))
	}

	res, err := h.Service.Create(c.Context(), req)
	if err != nil {
		log.Warn().Err(err).Msg("handler::CreateUser - Service returned error")
		code, errs := errmsg.Errors(err, &req)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(http.StatusCreated).JSON(response.Success(res, "User created successfully"))
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	results, err := h.Service.Get(c.Context())
	if err != nil {
		log.Warn().Err(err).Msg("handler::GetUsers - Service returned error")
		return c.Status(http.StatusInternalServerError).JSON(response.Error("Failed to retrieve users"))
	}
	return c.Status(http.StatusOK).JSON(response.Success(results, "Users retrieved successfully"))
}

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		log.Warn().Msg("handler::GetUserById - User ID is required")
		return c.Status(http.StatusBadRequest).JSON(response.Error("User ID is required"))
	}

	// Validate the user ID format
	if !utils.IsValidUUID(id) {
		log.Warn().Msg("handler::GetUserById - Invalid User ID format")
		return c.Status(http.StatusBadRequest).JSON(response.Error("Invalid User ID format"))
	}

	result, err := h.Service.GetById(c.Context(), id)
	if err != nil {
		log.Warn().Err(err).Msg("handler::GetUserById - Service returned error")
		code, errs := errmsg.Errors(err, &id)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(http.StatusOK).JSON(response.Success(result, "User retrieved successfully"))
}
