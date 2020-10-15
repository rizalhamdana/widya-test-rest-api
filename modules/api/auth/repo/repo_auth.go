package repo

import (
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/rizalhamdana/widya-test/helper"
	"github.com/rizalhamdana/widya-test/modules/share"
	"github.com/rizalhamdana/widya-test/modules/share/model"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryImpl struct {
	*share.Repository
}

func NewAuthRepositoryImpl(repo *share.Repository) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{repo}
}

func (r *AuthRepositoryImpl) CreateUser(user *model.User) share.ResultRepository {
	ctx := "repo_user.create_user"
	countData, err := r.DBRead.Model((*model.User)(nil)).Where("email = ?", user.Email).Count()

	if countData > 0 {
		err = errors.New("User with this email is already registered")
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot Create User")
		return share.ResultRepository{Error: err}
	}

	user.ID = uuid.NewV1().String()
	res, err := r.DBWrite.Model(user).Returning("*").Insert()
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot Create User")
		return share.ResultRepository{Error: err}
	}

	if res.RowsAffected() <= 0 {
		err := errors.New("Failed to register new user")
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot Create User")
		return share.ResultRepository{Error: err}
	}

	return share.ResultRepository{Result: "User registered successfully"}

}

func (r *AuthRepositoryImpl) GetUserByEmailAndPassword(email, password string) share.ResultRepository {
	ctx := "repo_user.get_user_by_email_and_password"
	user := model.User{}
	err := r.DBRead.Model(&user).Where("email = ?", email).Select()

	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			return share.ResultRepository{Error: nil, Count: 0}
		}
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot Get User")
		return share.ResultRepository{Error: err}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return share.ResultRepository{Error: errors.New("Invalid email or password"), Count: 0}
	}

	return share.ResultRepository{Result: user, Count: 1}

}

func (r *AuthRepositoryImpl) SaveToken(token *model.StoredToken, idUser string) share.ResultRepository {
	ctx := "repo_user.save_token"

	countData, err := r.DBRead.Model((*model.StoredToken)(nil)).Where("id_user = ? ", idUser).Count()
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot Save token User")
		return share.ResultRepository{Error: err}
	}
	if countData > 0 {
		_, err := r.DBWrite.Model(token).Where("id_user = ? ", token.IDUser).UpdateNotZero()
		if err != nil {
			helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot Save token User")
			return share.ResultRepository{Error: err}
		}
	}

	_, err = r.DBWrite.Model(token).Insert()
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot Save token User")
		return share.ResultRepository{Error: err}
	}

	return share.ResultRepository{Result: "Success Save token"}
}
