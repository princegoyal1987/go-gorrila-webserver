package models

import (
	"fmt"
)

type UserCurrency struct {
	UserId             int64
	CurrencyId			string
	Amount				int64
}

func (u *UserCurrency) String() string {
	return fmt.Sprintf("UserId(%v) %v=%v", u.UserId,u.CurrencyId,u.Amount);
}
