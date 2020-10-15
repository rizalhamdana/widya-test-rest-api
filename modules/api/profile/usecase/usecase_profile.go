package usecase

import (
	"github.com/rizalhamdana/widya-test/helper"
	"github.com/rizalhamdana/widya-test/modules/api/profile/repo"
	"github.com/rizalhamdana/widya-test/modules/share/model"
	"github.com/sirupsen/logrus"
)

type ProfileUsecaseImpl struct {
	repo repo.ProfileRepository
}

func NewProfileUsecaseImpl(repo repo.ProfileRepository) *ProfileUsecaseImpl {
	return &ProfileUsecaseImpl{repo: repo}
}

func (u *ProfileUsecaseImpl) GetProfile(ID string) (*model.User, error) {
	ctx := "usecase_profile.get_profile"

	getResult := u.repo.GetUserProfileById(ID)
	if getResult.Error != nil {
		helper.Log(logrus.ErrorLevel, getResult.Error.Error(), ctx, "Get Profile")
		return nil, getResult.Error
	}

	profile := getResult.Result.(model.User)
	return &profile, nil
}
