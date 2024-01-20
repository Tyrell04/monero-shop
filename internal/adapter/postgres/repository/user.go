package repository

import (
	"context"
	"github.com/google/uuid"
	"monero-shop-api/internal/adapter/postgres"
	"monero-shop-api/internal/core/domain"
	"monero-shop-api/internal/core/util"
)

// UserRepository implements port.UserRepository interface and provides an access to the postgres database
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

// GetUserByName returns a user by name
func (ur *UserRepository) GetUserByName(ctx context.Context, name string) (*domain.User, error) {
	var user domain.User
	query := ur.db.QueryBuilder.Select("*").
		From("users").
		Where("name = ?", name).
		Limit(1)

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

	return &user, nil
}

// GetUserByID returns a user by id
func (ur *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var user domain.User
	query := ur.db.QueryBuilder.Select("*").
		From("users").
		Where("user_id = ?", id).
		Limit(1)

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

	return &user, nil
}

// UpdateUser updates a user in the database
func (ur *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	var err error
	if err != nil {
		return nil, err
	}
	query := ur.db.QueryBuilder.Update("users").
		Set("name", user.Name).
		Set("password", user.Password).
		Where("user_id = ?", user.ID).
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

// DeleteUser deletes a user from the database
func (ur *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	query := ur.db.QueryBuilder.Delete("users").
		Where("user_id = ?", id)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = ur.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
