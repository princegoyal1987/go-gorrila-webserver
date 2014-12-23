package controller

import (
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/princegoyal1987/go-gorrila-webserver/models"
    "fmt"
)
func Update(w http.ResponseWriter, r *http.Request) {
	userId,_ := strconv.ParseInt(r.URL.Query().Get("user_id"),10,64)
	amount,_ := strconv.ParseInt(r.URL.Query().Get("amount"),10,64)
        currencyId := r.URL.Query().Get("currency_id")
	var userCurrency models.UserCurrency
	err := models.Dbm.SelectOne(&userCurrency, `select * from UserCurrency where UserId = ? and CurrencyId = ?`, userId, currencyId)
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
