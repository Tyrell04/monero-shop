package repository

import (
	"context"
	"github.com/google/uuid"
	"monero-shop-api/internal/adapter/postgres"
	"monero-shop-api/internal/core/domain"
	"monero-shop-api/internal/core/util"
)

/**
 * UserRepository implements port.UserRepository interface
 * and provides an access to the postgres database
 */
type UserRepository struct {
	db     *postgres.DB
	config *util.Config
}

// NewUserRepository creates a new user repository instance
func NewUserRepository(db *postgres.DB, config *util.Config) *UserRepository {
	return &UserRepository{
		db,
		config,
	}
}

// CreateUser creates a new user in the database
func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	var err error
	if err != nil {
		return nil, err
	}
	query := ur.db.QueryBuilder.Insert("users").
		Columns("user_id", "name", "password").
		Values(uuid.New(), user.Name, user.Password).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = ur.db.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
