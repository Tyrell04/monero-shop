package repository

import (
	"context"
	"monero-shop-api/internal/adapter/postgres"
	"monero-shop-api/internal/core/domain"
	"monero-shop-api/internal/core/util"
	"testing"
)

func dbSetup(con context.Context) (*postgres.DB, error) {
	db, err := postgres.New(con, &util.Config{
		Database: util.Database{
			Name:     "postgres",
			Password: "postgres",
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
		},
	})
	if err != nil {
		return db, err
	}
	return db, nil
}

func TestCreateUser(t *testing.T) {
	con := context.Background()
	db, err := dbSetup(con)
	if err != nil {
		t.Error(err)
	}
	user := &UserRepository{
		db: db,
	}
	randomString, err := util.RandomString(5)
	userTest, err := user.CreateUser(con, &domain.User{
		Name:     "test" + randomString,
		Password: "test",
	})
	if err != nil {
		t.Error(err)
	}
	if userTest.Name != string("test"+randomString) {
		t.Error("user name is not test")
	}
	if userTest.Password != "test" {
		t.Error("user password is not test")
	}
}
