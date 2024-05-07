package repo

import (
	"context"
	"fmt"
	"user-service/domain/user"

	"user-service/domain/user/model"

	"github.com/go-pg/pg/v10"
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

	query := `INSERT INTO users (id, full_name, email, phone_number, user_password, created_by, created_at) values ('%s', '%s', '%s', '%s','%s', '%s', now())`
	query = fmt.Sprintf(query, req.Id, req.FullName, req.Email, req.PhoneNumber, req.UserPassword, req.Id)

	_, err = r.gopg.ExecContext(ctx, query)

	if err != nil {
		return
	}

	userId = req.Id

	return
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (res model.Users, err error) {

	err = r.gopg.ModelContext(ctx, &res).Where("email=?", email).First()

	if err != nil {
		if err != pg.ErrNoRows {
			return
		}
	}

	return
}
