package handler

import (
	"echo-jwt-starter/internal/dto"
	"echo-jwt-starter/internal/service"
	"echo-jwt-starter/middleware"
	"echo-jwt-starter/pkg/errmsg"
	"echo-jwt-starter/pkg/response"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type AuthHandler struct {
	Service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		log.Info().Err(err).Msg("handler::Login - Failed to bind request body")
		code, errs := errmsg.Errors(err, &req)
		return c.JSON(code, response.Error(errs))
	}

	if err := c.Validate(&req); err != nil {
		log.Info().Err(err).Msg("handler::Login - Validation failed")
		code, errs := errmsg.Errors(err, &req)
		return c.JSON(code, response.Error(errs))
	}

	res, err := h.Service.Login(c.Request().Context(), req)
	if err != nil {
		log.Warn().Err(err).Msg("handler::Login - Service returned error")
		code, errs := errmsg.Errors(err, &req)
		return c.JSON(code, response.Error(errs))
	}

	return c.JSON(http.StatusOK, response.Success(res, "Login berhasil"))
}

func (h *AuthHandler) Refresh(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.JSON(http.StatusUnauthorized, response.Error("Unauthorized"))
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	res, err := h.Service.RefreshToken(c.Request().Context(), tokenStr)
	if err != nil {
		log.Warn().Err(err).Msg("handler::Refresh - Service returned error")
		code, errs := errmsg.Errors(err, &tokenStr)
		return c.JSON(code, response.Error(errs))
	}

	return c.JSON(http.StatusOK, response.Success(res, "Access token refreshed"))
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		log.Info().Err(err).Msg("handler::Register - Failed to bind request body")
		code, errs := errmsg.Errors(err, &req)
		return c.JSON(code, response.Error(errs))
	}
	if err := c.Validate(&req); err != nil {
		log.Info().Err(err).Msg("handler::Register - Validation failed")
		code, errs := errmsg.Errors(err, &req)
		return c.JSON(code, response.Error(errs))
	}

	res, err := h.Service.Register(c.Request().Context(), req)
	if err != nil {
		log.Warn().Err(err).Msg("handler::Register - Service returned error")
		code, errs := errmsg.Errors(err, &req)
		return c.JSON(code, response.Error(errs))
	}

	return c.JSON(http.StatusCreated, response.Success(res, "Registrasi berhasil"))
}

func (h *AuthHandler) Profile(c echo.Context) error {
	userId := middleware.GetUserIDFromContext(c)
	role := middleware.GetRoleFromContext(c)

	data := map[string]string{
		"user_id": userId,
		"role":    role,
	}
	return c.JSON(http.StatusOK, response.Success(data, "Profile loaded"))
}
