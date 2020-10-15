package repo

import "github.com/rizalhamdana/widya-test/modules/share"

type ProfileRepository interface {
	GetUserProfileById(ID string) share.ResultRepository
}
