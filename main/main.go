package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/princegoyal1987/go-gorrila-webserver/controller"
	"net/http"
        "github.com/gorilla/context"
)

var r = new(mux.Router)


func AddUserDataToContext(handler http.Handler) http.Handler {
    ourFunc := func(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Query()["user_id"]
        context.Set(r,"user_id",user_id)
	fmt.Print(user_id);
        handler.ServeHTTP(w, r)
    }
    return http.HandlerFunc(ourFunc)
}


func homeHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "sadfaf")
	return nil
}

func init1() {
	r.Handle("/", AddUserDataToContext(http.HandlerFunc(controller.HomeHandler1))).Name("home")
	r.HandleFunc("/1", controller.HomeHandler1).Name("home1")
        r.HandleFunc("/UserNew", controller.UserNew).Name("UserNew")
        r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}

func main() {
	init1()
	http.Handle("/", r)
	http.ListenAndServe(":8081", nil)
}
