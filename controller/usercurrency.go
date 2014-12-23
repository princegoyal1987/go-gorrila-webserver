package controller

import (
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/princegoyal1987/go-gorrila-webserver/models"
    "fmt"
)


func NewUserCurrency(userId int64,currencyName string,amount int64) (*models.UserCurrency,error) {
    userCurrency := &models.UserCurrency{UserId:userId,CurrencyName:currencyName,Amount:amount}
    if err := models.Dbm.Insert(userCurrency); err!=nil {
        panic(err)
    }
    return userCurrency,nil
}

func UserCurrencyUpdate(w http.ResponseWriter, r *http.Request) {
	userId,_ := strconv.ParseInt(r.URL.Query().Get("user_id"),10,64)
	amount,_ := strconv.ParseInt(r.URL.Query().Get("amount"),10,64)
        currencyName := r.URL.Query().Get("currency_name")
	var userCurrency models.UserCurrency
	err := models.Dbm.SelectOne(&userCurrency, `select * from UserCurrency where UserId = ? and CurrencyName = ?`, userId, currencyName)
	if err != nil {
		panic(err)
	}
	userCurrency.Amount = amount
	models.Dbm.Update(&userCurrency)
	
	if err != nil {
		panic(err)
	}
	
	json, _ := json.Marshal(userCurrency)
        fmt.Fprintf(w, string(json)) 
}

func UserCurrencyGet(w http.ResponseWriter, r *http.Request) {
    userId,_ := strconv.ParseInt(r.URL.Query().Get("user_id"),10,64)
    currencyName := r.URL.Query().Get("currency_name")
    var userCurrency *models.UserCurrency
    
    userCurrencies, _ := models.Dbm.Select(models.UserCurrency{}, `select * from UserCurrency where UserId = ? and CurrencyName = ?`, userId,currencyName)
    
    if len(userCurrencies)==0 {
        userCurrency,_ = NewUserCurrency(userId,currencyName,10000); //starting amount = 10000
    } else {
        userCurrency = userCurrencies[0].(*models.UserCurrency)
    }
    json, _ := json.Marshal(userCurrency)
    fmt.Fprintf(w, string(json)) 
}


