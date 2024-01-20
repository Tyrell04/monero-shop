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
	err = user.DeleteUser(con, userTest.ID)
}

func TestGetUserByName(t *testing.T) {
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
	userTest2, err := user.GetUserByName(con, userTest.Name)
	if err != nil {
		t.Error(err)
	}
	if userTest2.Name != userTest.Name {
		t.Error("user name is not test")
	}
	if userTest2.Password != userTest.Password {
		t.Error("user password is not test")
	}
	err = user.DeleteUser(con, userTest.ID)
}

func TestGetUserByID(t *testing.T) {
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
	userTest2, err := user.GetUserByID(con, userTest.ID)
	if err != nil {
		t.Error(err)
	}
	if userTest2.Name != userTest.Name {
		t.Error("user name is not test")
	}
	if userTest2.Password != userTest.Password {
		t.Error("user password is not test")
	}
	err = user.DeleteUser(con, userTest.ID)
}

func TestUpdateUser(t *testing.T) {
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
	userTest.Name = "test2"
	userTest2, err := user.UpdateUser(con, userTest)
	if err != nil {
		t.Error(err)
	}
	if userTest2.Name != userTest.Name {
		t.Error("user name is not test2")
	}
	if userTest2.Password != userTest.Password {
		t.Error("user password is not test")
	}
	err = user.DeleteUser(con, userTest.ID)
}

func TestDeleteUser(t *testing.T) {
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
	err = user.DeleteUser(con, userTest.ID)
	if err != nil {
		t.Error(err)
	}
	userTest2, err := user.GetUserByID(con, userTest.ID)
	if userTest2 != nil {
		if err != nil {
			t.Error(err)
		}
		t.Error("user should not exist")
	}
}
