package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
)

type AdminUser struct {
	UserId             int
	Name               string
	Username, Password string
	HashedPassword     []byte
}

func (u *AdminUser) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}

var adminUserRegex = regexp.MustCompile("^\\w*$")

func (user *AdminUser) Validate(v *revel.Validation) {
	v.Check(user.Username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	ValidatePassword(v, user.Password).
		Key("AdminUser.Password")

	v.Check(user.Name,
		revel.Required{},
		revel.MaxSize{100},
	)
}
