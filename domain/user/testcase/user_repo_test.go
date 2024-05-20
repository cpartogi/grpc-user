package test_user

import (
	"context"
	"os"
	"testing"

	"user-service/domain/user/model"
	"user-service/domain/user/repo"
	"user-service/lib/helper"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo(t *testing.T) {

	godotenv.Load("../../../.testenv")

	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("UNIT_TEST_DB_HOST") + ":" + os.Getenv("UNIT_TEST_DB_PORT"),
		User:     os.Getenv("UNIT_TEST_DB_USER"),
		Password: os.Getenv("UNIT_TEST_DB_PASSWORD"),
		Database: os.Getenv("UNIT_TEST_DB_DATABASE"),
	})
	defer db.Close()

	repo := repo.NewUserRepo(db)

	ctx := context.Background()

	userId := uuid.New().String()

	t.Run("Error Insert User", func(t *testing.T) {

		user := model.Users{
			Id:           "1",
			FullName:     "John Doe",
			Email:        "john@example.com",
			PhoneNumber:  "1234567890",
			UserPassword: "password123",
		}

		_, err := repo.InsertUser(ctx, user)

		assert.Error(t, err)
	})

	t.Run("Success Insert User", func(t *testing.T) {

		user := model.Users{
			Id:           userId,
			FullName:     "John Doe",
			Email:        "john@example.com",
			PhoneNumber:  "1234567890",
			UserPassword: "password123",
		}

		_, err := repo.InsertUser(ctx, user)

		assert.NoError(t, err)
	})

	t.Run("Success Get User by Email", func(t *testing.T) {

		_, err := repo.GetUserByEmail(ctx, "john@example.com")

		assert.NoError(t, err)
	})

	t.Run("Error Get User by Id", func(t *testing.T) {

		_, err := repo.GetUserById(ctx, helper.GenerateRandomString(125))

		assert.Error(t, err)
	})

	t.Run("Success get user by id", func(t *testing.T) {

		_, err := repo.GetUserById(ctx, userId)

		assert.NoError(t, err)
	})

	t.Run("Error Insert user log", func(t *testing.T) {

		err := repo.InsertUserLog(ctx, model.UserLogs{
			Id:           "a",
			UserId:       "b",
			IsSuccess:    false,
			LoginMessage: "",
		})

		assert.Error(t, err)
	})

	t.Run("Success Insert user log", func(t *testing.T) {

		err := repo.InsertUserLog(ctx, model.UserLogs{
			Id:           userId,
			UserId:       userId,
			IsSuccess:    false,
			LoginMessage: "success",
		})

		assert.NoError(t, err)
	})

	t.Run("Error update user", func(t *testing.T) {

		err := repo.UpdateUser(ctx, model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "c",
			PhoneNumber:  "d",
			UserPassword: "e",
		}, "d")

		assert.Error(t, err)
	})

	t.Run("Success update user", func(t *testing.T) {

		err := repo.UpdateUser(ctx, model.Users{
			Id:           userId,
			FullName:     "fullname",
			Email:        "email@example.com",
			PhoneNumber:  "d",
			UserPassword: "e",
		}, "a5a70646-0fbf-4db2-9265-e580f9ce863c")

		assert.NoError(t, err)
	})
}
