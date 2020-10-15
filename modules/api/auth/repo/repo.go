package repo

import (
	"github.com/rizalhamdana/widya-test/modules/share"
	"github.com/rizalhamdana/widya-test/modules/share/model"
)

type AuthRepository interface {
	CreateUser(user *model.User) share.ResultRepository
	GetUserByEmailAndPassword(email, password string) share.ResultRepository
	SaveToken(token *model.StoredToken, idUser string) share.ResultRepository
}
