package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/princegoyal1987/go-gorrila-webserver/controller"
        "github.com/princegoyal1987/go-gorrila-webserver/models"
	"net/http"
        "github.com/gorilla/context"
)

var r = new(mux.Router)


func AddUserDataToContext(handler http.Handler) http.Handler {
    ourFunc := func(w http.ResponseWriter, r *http.Request) {
        models.InitDB()
	user_id := r.URL.Query()["user_id"]
        context.Set(r,"user_id",user_id)
	fmt.Print(user_id);
        handler.ServeHTTP(w, r)
    }
    return http.HandlerFunc(ourFunc)
}


func init1() {
	r.Handle("/", AddUserDataToContext(http.HandlerFunc(controller.HomeHandler1))).Name("home")
        r.Handle("/user/new", AddUserDataToContext(http.HandlerFunc(controller.UserNew))).Name("home")
        r.Handle("/usercurrency/get",AddUserDataToContext(http.HandlerFunc(controller.UserCurrencyGet)))
        r.Handle("/usercurrency/update",AddUserDataToContext(http.HandlerFunc(controller.UserCurrencyUpdate)))

        r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

}

func main() {
	init1()
	http.Handle("/", r)
	http.ListenAndServe(":8081", nil)
}
