package postgres

import (
	"context"
	"github.com/stretchr/testify/assert"
	"monero-shop-api/internal/adapter/config"
	"testing"
)

func TestNew(t *testing.T) {
	// Mock configuration
	config := config.Config{
		Database: config.Database{
			Name:     "postgres",
			Password: "postgres",
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
		},
	}

	// Create a new context
	ctx := context.Background()

	// Test New function
	db, err := New(ctx, &config)

	// Assert that no error occurred during database creation
	assert.NoError(t, err)
	// Assert that the returned database object is not nil
	assert.NotNil(t, db)

	// Close the database connection
	db.Close()
}

func TestClose(t *testing.T) {
	// Mock configuration
	config := config.Config{
		Database: config.Database{
			Name:     "postgres",
			Password: "postgres",
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
		},
	}

	// Create a new context
	ctx := context.Background()

	// Create a new database instance
	db, err := New(ctx, &config)
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Test Close function
	db.Close()
	if db.Pool.Ping(ctx) == nil {
		t.Error("database connection is not closed")
	} else {
		t.Log("database connection is closed")
	}
}
