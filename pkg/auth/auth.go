package auth

import (
	"errors"
	"goblog/app/models/user"
	"goblog/pkg/session"

	"gorm.io/gorm"
)

func _getUid() string {
	_uid := session.Get("uid")
	uid, ok := _uid.(string)
	if ok && len(uid) > 0 {
		return uid
	}
	return ""
}

func User() user.User {

	uid := _getUid()

	if len(uid) > 0 {

		user, err := user.Get(uid)

		if err == nil {
			return user
		}

	}

	return user.User{}
}

func Attemp(email string, password string) error {

	_user, err := user.GetByEmail(email)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("用户名或密码错误")
		} else {
			return errors.New("服务器错误")
		}
	}

	if !_user.ComparePassword(password) {
		return errors.New("用户名或密码错误")
	}

	session.Put("uid", _user.GetStringId())

	return nil
}

func Login(_user user.User) {
	session.Put("uid", _user.GetStringId())
}

func Logout() {
	session.Forget("uid")
}

func Check() bool {
	return len(_getUid()) > 0
}
