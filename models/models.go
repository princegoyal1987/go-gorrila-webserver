package models


type User struct {
	UserId         int64
	Name           string
	DeviceId       string
	Email          string
	FacebookId     string
	HashedPassword string
}



type UserCurrency struct {
        UserCurrencyId int64
	UserId     int64
	CurrencyName string
	Amount     int64
}





