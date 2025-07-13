package port

import (
	"context"
	"fiber-lite-starter/internal/entity"
)

type UserRepository interface {
	Get(ctx context.Context) ([]*entity.UserDB, error)
	GetById(ctx context.Context, id string) (*entity.UserDB, error)
	FindByEmail(ctx context.Context, email string) (*entity.UserDB, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	Create(ctx context.Context, user *entity.UserDB) error
}
