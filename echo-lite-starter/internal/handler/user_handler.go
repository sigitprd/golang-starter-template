package handler

import (
	"echo-lite-starter/internal/dto"
	"echo-lite-starter/internal/service"
	"echo-lite-starter/pkg/errmsg"
	"echo-lite-starter/pkg/response"
	"echo-lite-starter/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req dto.UserRequest
	if err := c.Bind(&req); err != nil {
		log.Warn().Err(err).Msg("handler::CreateUser - Failed to bind request body")
		code, errs := errmsg.Errors(err, &req)
		return c.JSON(code, response.Error(errs))
	}

	if err := c.Validate(&req); err != nil {
		log.Warn().Err(err).Msg("handler::CreateUser - Validation failed")
		code, errs := errmsg.Errors(err, &req)
		return c.JSON(code, response.Error(errs))
	}

	res, err := h.Service.Create(c.Request().Context(), req)
	if err != nil {
		log.Warn().Err(err).Msg("handler::CreateUser - Service returned error")
		code, errs := errmsg.Errors(err, &req)
		return c.JSON(code, response.Error(errs))
	}

	return c.JSON(http.StatusCreated, response.Success(res, "User created successfully"))
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	results, err := h.Service.Get(c.Request().Context())
	if err != nil {
		log.Warn().Err(err).Msg("handler::GetUsers - Service returned error")
		return c.JSON(500, response.Error("Failed to retrieve users"))
	}

	return c.JSON(http.StatusOK, response.Success(results, "Users retrieved successfully"))
}

func (h *UserHandler) GetUserById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		log.Warn().Msg("handler::GetUserById - User ID is required")
		return c.JSON(http.StatusBadRequest, response.Error("User ID is required"))
	}

	if !utils.IsValidUUID(id) {
		log.Warn().Msg("handler::GetUserById - Invalid User ID format")
		return c.JSON(http.StatusBadRequest, response.Error("Invalid User ID format"))
	}

	result, err := h.Service.GetById(c.Request().Context(), id)
	if err != nil {
		log.Warn().Err(err).Msg("handler::GetUserById - Service returned error")
		code, errs := errmsg.Errors(err, &id)
		return c.JSON(code, response.Error(errs))
	}

	return c.JSON(http.StatusOK, response.Success(result, "User retrieved successfully"))
}
