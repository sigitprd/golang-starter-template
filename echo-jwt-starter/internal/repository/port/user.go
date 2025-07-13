package port

import (
	"context"
	"echo-jwt-starter/internal/entity"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.UserDB, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	Create(ctx context.Context, user *entity.UserDB) error
}
