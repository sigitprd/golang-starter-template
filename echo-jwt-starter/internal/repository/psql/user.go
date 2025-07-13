package psql

import (
	"context"
	"database/sql"
	"echo-jwt-starter/internal/entity"
	"echo-jwt-starter/internal/repository/port"
	"echo-jwt-starter/pkg/errmsg"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type UserRepository struct {
	DB DBExecutor
}

func NewUserRepositoryImpl(db DBExecutor) port.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.UserDB, error) {
	var user entity.UserDB
	// your implementation here
	query := `
		SELECT u.id, u.email, u.password, u.role
		FROM public.users u
		WHERE u.email = $1 AND u.deleted_at IS NULL
		LIMIT 1
	`

	if err := r.DB.QueryRowContext(ctx, query, email).
		Scan(
			&user.Id,
			&user.Email,
			&user.Password,
			&user.Role,
		); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error().Err(err).Str("email", email).Msg("repo::FindByEmail - User not found")
			return nil, errmsg.NewCustomErrors(404, errmsg.WithMessage(errmsg.UserNotFound))
		}
		log.Error().Err(err).Str("email", email).Msg("repo::FindByEmail - Failed to get user")
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM public.users u
			WHERE u.email = $1 AND u.deleted_at IS NULL
		);
	`
	var exists bool
	if err := r.DB.QueryRowContext(ctx, query, email).Scan(&exists); err != nil {
		log.Error().Err(err).Str("email", email).Msg("repo::ExistsByEmail - Failed to check user existence")
		return false, errmsg.NewCustomErrors(500, errmsg.WithMessage("Failed to check user existence"))
	}
	return exists, nil
}

func (r *UserRepository) Create(ctx context.Context, user *entity.UserDB) error {
	query := `
		INSERT INTO public.users (id, email, password, role)
		VALUES ($1, $2, $3, $4);
	`
	result, err := r.DB.ExecContext(ctx, query, user.Id, user.Email, user.Password, user.Role)
	if err != nil {
		log.Error().Err(err).Str("email", user.Email).Msg("repo::Create - Failed to create user")
		return errmsg.NewCustomErrors(500, errmsg.WithMessage("Failed to create user"))
	}
	// Check if the insert was successful
	if rowsAffected, err := result.RowsAffected(); err != nil {
		log.Error().Err(err).Str("email", user.Email).Msg("repo::Create - Failed to check rows affected")
		return errmsg.NewCustomErrors(500, errmsg.WithMessage("Failed to create user"))
	} else if rowsAffected != 1 {
		log.Error().Str("email", user.Email).Int64("rowsAffected", rowsAffected).Msg("repo::Create - Unexpected number of rows affected")
		return errmsg.NewCustomErrors(500, errmsg.WithMessage("Failed to create user"))
	}
	return nil
}
