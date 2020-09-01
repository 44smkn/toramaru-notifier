package usecase

import (
	"jweb-notifier/domain/model/entity"
	"jweb-notifier/domain/model/vo"
	"jweb-notifier/presentation/param"

	"github.com/pkg/errors"
)

func RegisterUser(u *param.User) error {
	id, err := vo.NewUserId(u.Id)
	if err != nil {
		return errors.Wrapf(err, "creating vo of userid is fail. id: %s", u.Id)
	}
	password, err := vo.NewPassword(u.Password)
	if err != nil {
		return errors.Wrapf(err, "creating vo of userid is fail. password: %s", u.Password)
	}
	user := entity.NewUser(id, password)
	return repo.Register(user)
}
