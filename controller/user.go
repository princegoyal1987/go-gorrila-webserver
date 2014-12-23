package controller

import (
	"encoding/json"
	"fmt"
	"github.com/princegoyal1987/go-gorrila-webserver/models"
	"net/http"
	"strconv"
)


func createNew(facebookId, name, email, deviceId, bcryptPassword string) (*models.User,error) {
	user := &models.User{0, name, deviceId, email, facebookId, bcryptPassword}
	if err := models.Dbm.Insert(user); err != nil {
		panic(err)
	}
	return user,nil
}

func getWithFacebookId(facebookId string) (*models.User,error) {
	var user *models.User
	users, err := models.Dbm.Select(models.User{}, `select * from User where FacebookId = ?`, facebookId)
	if err != nil {
		panic(err)
	}
	if len(users) != 0 {
		user = users[0].(*models.User)
	}
	return user,nil
}


func UserNew(w http.ResponseWriter, r *http.Request) {
	facebookId := r.URL.Query().Get("facebook_id")
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")
	deviceId := r.URL.Query().Get("device_id")
	bcryptPassword := ""
	var user *models.User

        user,_ = getWithFacebookId(facebookId)

	if user == nil {
	    user,_ = createNew(facebookId,name,email,deviceId,bcryptPassword)
        }

	userJson, _ := json.Marshal(user)
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, string(userJson))
}

func GetData(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	userCurrencies, err := models.Dbm.Select(models.UserCurrency{}, `select * from UserCurrency where UserId = ?`, userId)
	if err != nil {
		panic(err)
	}

	if len(userCurrencies) == 0 {
		var userCurrency *models.UserCurrency
		userIdInt, _ := strconv.ParseInt(userId, 0, 64)

                userCurrency = &models.UserCurrency{UserId:userIdInt, CurrencyName:"silver", Amount:10000}

		if err := models.Dbm.Insert(userCurrency); err != nil {
			panic(err)
		}

		userCurrencies = append(userCurrencies, userCurrency)
	}

	userJson, _ := json.Marshal(userCurrencies)
	fmt.Fprintf(w, string(userJson))
}
