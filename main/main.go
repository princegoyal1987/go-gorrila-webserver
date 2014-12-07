package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/princegoyal1987/go-gorrila-webserver/controller"
	"net/http"
)

var r = new(mux.Router)

type handlerFunc func(http.ResponseWriter, *http.Request) error

func (f handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := f(w, r)
	if err != nil {
		panic("and it failed")
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "sadfaf")
	return nil
}

func init1() {
	r.Handle("/", handlerFunc(homeHandler)).Name("home")
	r.HandleFunc("/1", controller.HomeHandler1).Name("home1")
        r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}

type funcType func() error

func main() {
	init1()
	http.Handle("/", r)
	http.ListenAndServe(":8081", nil)
}
