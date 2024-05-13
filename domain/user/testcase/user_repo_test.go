package test_user

import (
	"context"
	"os"
	"testing"
	"time"

	"user-service/domain/user/model"
	"user-service/domain/user/repo"

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

		_, err := repo.InsertUser(ctx, user, "a")

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

		_, err := repo.InsertUser(ctx, user, "a")

		assert.NoError(t, err)
	})

	t.Run("Success Get User by Email", func(t *testing.T) {

		_, err := repo.GetUserByEmail(ctx, "john@example.com", "a")

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

	t.Run("Error UpsertUserToken", func(t *testing.T) {

		err := repo.UpsertUserToken(ctx, model.UserToken{
			Id:                    "cde",
			Token:                 "a",
			TokenExpiredAt:        time.Time{},
			RefreshToken:          "b",
			RefreshTokenExpiredAt: time.Time{},
		})

		assert.Error(t, err)
	})

	t.Run("Success UpsertUserToken insert", func(t *testing.T) {

		err := repo.UpsertUserToken(ctx, model.UserToken{
			Id:                    userId,
			Token:                 "a",
			TokenExpiredAt:        time.Time{},
			RefreshToken:          "b",
			RefreshTokenExpiredAt: time.Time{},
		})

		assert.NoError(t, err)
	})

	t.Run("succes UpsertUserToken update", func(t *testing.T) {

		err := repo.UpsertUserToken(ctx, model.UserToken{
			Id:             userId,
			Token:          "a",
			TokenExpiredAt: time.Time{},
		})

		assert.NoError(t, err)
	})

	t.Run("succes UpsertUserToken update refresh token", func(t *testing.T) {

		err := repo.UpsertUserToken(ctx, model.UserToken{
			Id:             userId,
			Token:          "a",
			RefreshToken:   "b",
			TokenExpiredAt: time.Time{},
		})

		assert.NoError(t, err)
	})

}
