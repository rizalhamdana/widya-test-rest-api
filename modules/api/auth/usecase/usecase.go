package usecase

import "github.com/rizalhamdana/widya-test/modules/share/model"

type AuthUsecase interface {
	RegisterUser(user *model.User) (*string, error)
	LoginUser(email, password string) (*string, error)
}
