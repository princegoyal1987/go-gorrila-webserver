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
	UserId     int64
	CurrencyId string
	Amount     int64
}





