package user_models_domain

import (
	"fmt"

	"github.com/karasuneo/go-di/src/utils"
)

type User struct {
	id   UserId
	name string
	mail string
}

func NewUser(
	id *string,
	name string,
	mail string,
) (*User, error) {
	userId, err := NewUserId(id)
	if err != nil {
		return nil, err
	}

	if len(name) > 50 {
		return nil, fmt.Errorf("name is too long")
	}

	if !utils.ValidateEmail(&mail) {
		return nil, fmt.Errorf("invalid email")
	}

	return &User{
		id:   *userId,
		name: name,
		mail: mail,
	}, nil
}

func (u *User) GetId() string {
	return u.id.GetValue()
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetMail() string {
	return u.mail
}
