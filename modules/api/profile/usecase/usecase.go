package usecase

import "github.com/rizalhamdana/widya-test/modules/share/model"

type ProfileUsecase interface {
	GetProfile(ID string) (*model.User, error)
}
