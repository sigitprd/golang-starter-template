package handler

import (
	"fiber-jwt-starter/internal/dto"
	"fiber-jwt-starter/internal/service"
	"fiber-jwt-starter/middleware"
	"fiber-jwt-starter/pkg/errmsg"
	"fiber-jwt-starter/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

type AuthHandler struct {
	Service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		log.Info().Err(err).Msg("handler::Login - Failed to parse request body")
		code, errs := errmsg.Errors(err, &req)
		return c.Status(code).JSON(response.Error(errs))
	}

	if err := c.Locals("validator").(func(interface{}) error)(&req); err != nil {
		log.Info().Err(err).Msg("handler::Login - Validation failed")
		code, errs := errmsg.Errors(err, &req)
		return c.Status(code).JSON(response.Error(errs))
	}

	res, err := h.Service.Login(c.Context(), req)
	if err != nil {
		log.Warn().Err(err).Msg("handler::Login - Service returned error")
		code, errs := errmsg.Errors(err, &req)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(http.StatusOK).JSON(response.Success(res, "Login berhasil"))
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Error("Unauthorized"))
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	res, err := h.Service.RefreshToken(c.Context(), tokenStr)
	if err != nil {
		log.Warn().Err(err).Msg("handler::Refresh - Service returned error")
		code, errs := errmsg.Errors(err, &tokenStr)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(http.StatusOK).JSON(response.Success(res, "Access token refreshed"))
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		log.Info().Err(err).Msg("handler::Register - Failed to parse request body")
		code, errs := errmsg.Errors(err, &req)
		return c.Status(code).JSON(response.Error(errs))
	}

	if err := c.Locals("validator").(func(interface{}) error)(&req); err != nil {
		log.Info().Err(err).Msg("handler::Register - Validation failed")
		code, errs := errmsg.Errors(err, &req)
		return c.Status(code).JSON(response.Error(errs))
	}

	res, err := h.Service.Register(c.Context(), req)
	if err != nil {
		log.Warn().Err(err).Msg("handler::Register - Service returned error")
		code, errs := errmsg.Errors(err, &req)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(http.StatusCreated).JSON(response.Success(res, "Registrasi berhasil"))
}

func (h *AuthHandler) Profile(c *fiber.Ctx) error {
	userId := middleware.GetUserIDFromContext(c)
	role := middleware.GetRoleFromContext(c)

	data := map[string]string{
		"user_id": userId,
		"role":    role,
	}
	return c.Status(http.StatusOK).JSON(response.Success(data, "Profile loaded"))
}
