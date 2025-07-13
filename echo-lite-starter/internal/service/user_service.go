package service

import (
	"context"
	"echo-lite-starter/config"
	"echo-lite-starter/internal/dto"
	"echo-lite-starter/internal/entity"
	"echo-lite-starter/internal/repository/port"
	"echo-lite-starter/pkg/errmsg"
	"echo-lite-starter/pkg/utils"
)

type UserService interface {
	Get(ctx context.Context) ([]dto.UserResponse, error)
	GetById(ctx context.Context, id string) (dto.UserResponse, error)
	Create(ctx context.Context, req dto.UserRequest) (dto.UserResponse, error)
}

type UserServiceImpl struct {
	cfg        *config.Config
	repository port.RepositoryRegistry
}

func NewUserService(repo port.RepositoryRegistry) UserService {
	return &UserServiceImpl{
		cfg:        config.Envs,
		repository: repo,
	}
}

func (s *UserServiceImpl) Get(ctx context.Context) ([]dto.UserResponse, error) {
	userRepo := s.repository.GetUserRepository()
	users, err := userRepo.Get(ctx)
	if err != nil {
		return nil, err
	}

	var responses []dto.UserResponse
	for _, user := range users {
		responses = append(responses, dto.UserResponse{
			Id:    user.Id,
			Email: user.Email,
			Role:  user.Role,
		})
	}

	return responses, nil
}

func (s *UserServiceImpl) GetById(ctx context.Context, id string) (dto.UserResponse, error) {
	userRepo := s.repository.GetUserRepository()
	user, err := userRepo.GetById(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		Id:    user.Id,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (s *UserServiceImpl) Create(ctx context.Context, req dto.UserRequest) (dto.UserResponse, error) {
	userRepo := s.repository.GetUserRepository()

	// Cek email sudah terdaftar
	existing, err := userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return dto.UserResponse{}, err
	}

	if existing {
		return dto.UserResponse{}, errmsg.NewCustomErrors(409, errmsg.WithErrors("email", "Email sudah terdaftar"))
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return dto.UserResponse{}, errmsg.NewCustomErrors(500, errmsg.WithMessage("Gagal mengenkripsi password"))
	}

	// Simpan user
	user := &entity.UserDB{
		Id:       utils.GenerateID(),
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "user",
	}
	if err = userRepo.Create(ctx, user); err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		Id:    user.Id,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}
