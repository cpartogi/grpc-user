package repo

import (
	"context"
	"fmt"
	"user-service/domain/user"

	"user-service/domain/user/model"

	logger "user-service/lib/pkg/logger"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type UserRepo struct {
	gopg *pg.DB
}

func NewUserRepo(gopg *pg.DB) user.UserRepoInterface {
	return &UserRepo{
		gopg: gopg,
	}
}

func (r *UserRepo) InsertUser(ctx context.Context, req model.Users) (userId string, err error) {
	functionName := "user-service.repo.InsertUser"

	query := `INSERT INTO users (id, full_name, email, phone_number, user_password, created_by, created_at) values ('%s', '%s', '%s', '%s','%s', '%s', now())`
	query = fmt.Sprintf(query, req.Id, req.FullName, req.Email, req.PhoneNumber, req.UserPassword, req.Id)

	_, err = r.gopg.ExecContext(ctx, query)

	if err != nil {
		logger.Log(ctx, functionName, err.Error(), req, nil)
		return
	}

	logger.Log(ctx, functionName, "", req, userId)

	userId = req.Id

	return
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (res model.Users, err error) {
	functionName := "user-service.repo.GetUserByEmail"

	err = r.gopg.ModelContext(ctx, &res).Where("email=?", email).First()

	if err != nil {
		if err != pg.ErrNoRows {
			logger.Log(ctx, functionName, err.Error(), email, res)
			return
		} else {
			logger.Log(ctx, functionName, "", email, res)
			return res, nil
		}
	}

	logger.Log(ctx, functionName, "", email, res)

	return
}

func (r *UserRepo) InsertUserLog(ctx context.Context, req model.UserLogs) (err error) {
	functionName := "user-service.repo.InsertUserLog"
	req.Id = uuid.New().String()

	query := `INSERT INTO user_logs (id, user_id, is_success, login_message, created_at) values ('%s', '%s', '%t', '%s', now())`
	query = fmt.Sprintf(query, req.Id, req.UserId, req.IsSuccess, req.LoginMessage)

	_, err = r.gopg.ExecContext(ctx, query)

	if err != nil {
		logger.Log(ctx, functionName, err.Error(), req, nil)
		return
	}

	logger.Log(ctx, functionName, "", req, nil)
	return

}

func (r *UserRepo) GetUserById(ctx context.Context, id string) (res model.Users, err error) {
	functionName := "user-service.repo.GetUserById"
	err = r.gopg.ModelContext(ctx, &res).Where("id=?", id).First()

	if err != nil {
		if err != pg.ErrNoRows {
			logger.Log(ctx, functionName, err.Error(), id, nil)
			return
		}
	}

	logger.Log(ctx, functionName, "", id, nil)
	return
}

func (r *UserRepo) UpdateUser(ctx context.Context, req model.Users, userId string) (err error) {
	functionName := "user-service.repo.UpdateUser"

	query := `UPDATE users SET full_name  = '%s', phone_number = '%s', user_password  = '%s', updated_at = now(), updated_by = '%s' WHERE id = '%s' `
	query = fmt.Sprintf(query, req.FullName, req.PhoneNumber, req.UserPassword, req.Id, userId)

	_, err = r.gopg.ExecContext(ctx, query)

	if err != nil {
		logger.Log(ctx, functionName, err.Error(), req, nil)
		return
	}

	logger.Log(ctx, functionName, "", req, nil)
	return
}
