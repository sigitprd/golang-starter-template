package service

import (
	"context"
	"fiber-jwt-starter/config"
	"fiber-jwt-starter/internal/dto"
	"fiber-jwt-starter/internal/entity"
	"fiber-jwt-starter/internal/repository/port"
	"fiber-jwt-starter/pkg/errmsg"
	"fiber-jwt-starter/pkg/jwthandler"
	"fiber-jwt-starter/pkg/utils"
	"net/http"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	RefreshToken(ctx context.Context, tokenStr string) (dto.LoginResponse, error)
	Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error)
}

type AuthServiceImpl struct {
	cfg        *config.Config
	repository port.RepositoryRegistry
}

func NewAuthService(repo port.RepositoryRegistry) AuthService {
	return &AuthServiceImpl{
		cfg:        config.Envs,
		repository: repo,
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	// 1. Get user by email
	userRepo := s.repository.GetUserRepository()
	user, err := userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	// 2. Compare password
	if err = utils.ComparePassword(user.Password, req.Password); err != nil {
		return dto.LoginResponse{}, errmsg.NewCustomErrors(400, errmsg.WithMessage("Password salah"))
	}

	// 3. Generate tokens
	accessToken, err := jwthandler.GenerateToken(jwthandler.Payload{
		ID:              user.Id,
		Role:            user.Role,
		Subject:         jwthandler.AccessToken,
		ExpirationHours: s.cfg.Guard.JwtTtlHours, // or use config.Envs.Guard.JwtTtlHours
	})
	if err != nil {
		return dto.LoginResponse{}, errmsg.NewCustomErrors(500, errmsg.WithMessage("Gagal membuat access token"))
	}

	refreshToken, err := jwthandler.GenerateToken(jwthandler.Payload{
		ID:              user.Id,
		Role:            user.Role,
		Subject:         "refresh_token",
		ExpirationHours: s.cfg.Guard.JwtRefreshTtlDays * 24, // Convert days to hours
	})
	if err != nil {
		return dto.LoginResponse{}, errmsg.NewCustomErrors(500, errmsg.WithMessage("Gagal membuat refresh token"))
	}

	// 4. Return response
	return dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthServiceImpl) RefreshToken(ctx context.Context, refreshToken string) (dto.LoginResponse, error) {
	claims, err := jwthandler.ParseToken(refreshToken)
	if err != nil || claims.Subject != string(jwthandler.RefreshToken) {
		return dto.LoginResponse{}, errmsg.NewCustomErrors(http.StatusUnauthorized, errmsg.WithMessage("Invalid refresh token"))
	}

	accessToken, err := jwthandler.GenerateToken(jwthandler.Payload{
		ID:              claims.ID,
		Role:            claims.Role,
		Subject:         jwthandler.AccessToken,
		ExpirationHours: s.cfg.Guard.JwtTtlHours, // or use config.Envs.Guard.JwtTtlHours,
	})
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken, // reuse, atau generate ulang kalau mau
	}, nil
}

func (s *AuthServiceImpl) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	userRepo := s.repository.GetUserRepository()

	// Cek email sudah terdaftar
	existing, err := userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	if existing {
		return dto.RegisterResponse{}, errmsg.NewCustomErrors(409, errmsg.WithErrors("email", "Email sudah terdaftar"))
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return dto.RegisterResponse{}, errmsg.NewCustomErrors(500, errmsg.WithMessage("Gagal mengenkripsi password"))
	}

	// Simpan user
	user := &entity.UserDB{
		Id:       utils.GenerateID(),
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "user",
	}
	if err = userRepo.Create(ctx, user); err != nil {
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{
		ID:    user.Id,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}
