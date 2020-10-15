package usecase

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rizalhamdana/widya-test/helper"
	"github.com/rizalhamdana/widya-test/modules/api/auth/repo"
	"github.com/rizalhamdana/widya-test/modules/share/model"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecaseImpl struct {
	repo repo.AuthRepository
}

func NewAuthUsecaseImpl(repo repo.AuthRepository) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{repo: repo}
}

func (u *AuthUsecaseImpl) RegisterUser(user *model.User) (*string, error) {
	ctx := "usecase_auth.register_user"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Register user")
		return nil, errors.New("There is something wrong in our API, please try again later")
	}
	user.Password = string(hashedPassword)
	createResult := u.repo.CreateUser(user)

	if createResult.Error != nil {
		helper.Log(logrus.ErrorLevel, createResult.Error.Error(), ctx, "Register user")
		return nil, createResult.Error
	}

	createSuccessMessage := createResult.Result.(string)
	return &createSuccessMessage, nil
}

func (u *AuthUsecaseImpl) LoginUser(email, password string) (*string, error) {
	ctx := "usecase_auth.login_user"

	getResult := u.repo.GetUserByEmailAndPassword(email, password)

	if getResult.Error != nil {
		helper.Log(logrus.ErrorLevel, getResult.Error.Error(), ctx, "Login user")
		return nil, getResult.Error
	}

	if getResult.Count <= 0 {
		err := errors.New("Invalid email or password")
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Login user")
		return nil, err
	}
	user := getResult.Result.(model.User)
	token, err := u.createToken(user.Email, user.ID, user.Password, "OulCEetwinW6oYrNktgI")
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Login user")
		return nil, err
	}
	tokenModel := model.StoredToken{
		IDUser: user.ID,
		Token:  token,
	}
	u.repo.SaveToken(&tokenModel, user.ID)

	return &token, nil

}

func (u *AuthUsecaseImpl) createToken(email string, id string, pass string, signature string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"email":    email,
		"pass":     pass,
		"login_at": time.Now().Format(time.RFC3339),
	})
	tokenString, err = token.SignedString([]byte(signature))
	return tokenString, err
}
