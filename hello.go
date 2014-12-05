package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "hello1"
)

var router = new(mux.Router)

func init1() {
    router.HandleFunc("/",homeHandler).Name("home")
    router.HandleFunc("/1",hello1.HomeHandler).Name("home1")
}

func homeHandler(w http.ResponseWriter,r *http.Request) {
    w.Header().Add("Content-Type", "text/html")
    fmt.Fprintf(w,"sadfaf")
}

func ab() {
}

func main() {
    init1()
    http.Handle("/", router)
    http.ListenAndServe(":8081",nil)
}
