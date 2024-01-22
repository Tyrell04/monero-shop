package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"monero-shop-api/internal/adapter/config"
)

/*
* DB is a wrapper for PostgreSQL database connection
* that uses pgxpool as database driver.
* It also holds a reference to squirrel.StatementBuilderType
* which is used to build SQL queries that compatible with PostgreSQL syntax
 */
type DB struct {
	*pgxpool.Pool
	QueryBuilder *squirrel.StatementBuilderType
}

// New creates a new PostgreSQL database instance
func New(ctx context.Context, config *config.Config) (*DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &DB{
		db,
		&psql,
	}, nil
}

// Close closes the database connection
func (db *DB) Close() {
	db.Pool.Close()
}
