package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
)

type User struct {
	UserId         int64
	Name           string
	DeviceId       string
	Email          string
	FacebookId     string
	HashedPassword string
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.UserId)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *User) Validate(v *revel.Validation) {

}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}
