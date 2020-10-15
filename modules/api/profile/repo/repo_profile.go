package repo

import (
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/rizalhamdana/widya-test/helper"
	"github.com/rizalhamdana/widya-test/modules/share"
	"github.com/rizalhamdana/widya-test/modules/share/model"
	"github.com/sirupsen/logrus"
)

type ProfileRepoImpl struct {
	*share.Repository
}

func NewProfileRepoImpl(repo *share.Repository) *ProfileRepoImpl {
	return &ProfileRepoImpl{repo}
}

func (r *ProfileRepoImpl) GetUserProfileById(ID string) share.ResultRepository {
	ctx := "repo_user.create_user"

	profile := model.User{}

	err := r.DBRead.Model(&profile).Where("id = ?", ID).ExcludeColumn("password").Select()
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			err = errors.New("Profile not found")
		}
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Get Profile")
		return share.ResultRepository{Error: err, Count: 0}
	}

	return share.ResultRepository{Result: profile}

}
