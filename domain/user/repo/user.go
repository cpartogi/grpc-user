package repo

import (
	"context"
	"fmt"
	"time"
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

func (r *UserRepo) InsertUser(ctx context.Context, req model.Users, requestId string) (userId string, err error) {
	functionName := "repo.InsertUser"

	query := `INSERT INTO users (id, full_name, email, phone_number, user_password, created_by, created_at) values ('%s', '%s', '%s', '%s','%s', '%s', now())`
	query = fmt.Sprintf(query, req.Id, req.FullName, req.Email, req.PhoneNumber, req.UserPassword, req.Id)

	_, err = r.gopg.ExecContext(ctx, query)

	if err != nil {
		logger.GetLogger(functionName, err.Error(), requestId, req, nil)
		return
	}

	logger.GetLogger(functionName, "", requestId, req, userId)

	userId = req.Id

	return
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email, requestId string) (res model.Users, err error) {
	functionName := "repo.GetUserByEmail"

	err = r.gopg.ModelContext(ctx, &res).Where("email=?", email).First()

	if err != nil {
		if err != pg.ErrNoRows {
			logger.GetLogger(functionName, err.Error(), requestId, email, res)
			return
		} else {
			logger.GetLogger(functionName, "", requestId, email, res)
			return res, nil
		}
	}

	logger.GetLogger(functionName, "", requestId, email, res)

	return
}

func (r *UserRepo) InsertUserLog(ctx context.Context, req model.UserLogs) (err error) {

	req.Id = uuid.New().String()

	query := `INSERT INTO user_logs (id, user_id, is_success, login_message, created_at) values ('%s', '%s', '%t', '%s', now())`
	query = fmt.Sprintf(query, req.Id, req.UserId, req.IsSuccess, req.LoginMessage)

	_, err = r.gopg.ExecContext(ctx, query)

	if err != nil {
		return
	}

	return

}

func (r *UserRepo) UpsertUserToken(ctx context.Context, req model.UserToken) (err error) {

	// check data exist
	err = r.gopg.ModelContext(ctx, &model.UserToken{}).Where("user_id=?", req.Id).First()
	var query string
	id := uuid.New().String()

	tokenExpiredAt := req.TokenExpiredAt.UTC().Format(time.RFC3339)
	refreshTokenExpiredAt := req.RefreshTokenExpiredAt.UTC().Format(time.RFC3339)

	if err != nil {
		if err == pg.ErrNoRows {
			query = `INSERT INTO user_tokens (id, user_id, token, token_expired_at, refresh_token, refresh_token_expired_at, created_at) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', now()) `
			query = fmt.Sprintf(query, id, req.Id, req.Token, tokenExpiredAt, req.RefreshToken, refreshTokenExpiredAt)
		} else {
			return
		}
	} else {

		if req.RefreshToken != "" {
			query = `UPDATE user_tokens SET  token = '%s', token_expired_at = '%s', refresh_token = '%s',  refresh_token_expired_at = '%s', updated_at = now()  WHERE user_id = '%s' `
			query = fmt.Sprintf(query, req.Token, tokenExpiredAt, req.RefreshToken, refreshTokenExpiredAt, req.Id)
		} else {
			query = `UPDATE user_tokens SET  token = '%s', token_expired_at = '%s', updated_at = now()  WHERE user_id = '%s' `
			query = fmt.Sprintf(query, req.Token, tokenExpiredAt, req.Id)
		}
	}

	_, err = r.gopg.ExecContext(ctx, query)

	if err != nil {
		return
	}

	return
}
