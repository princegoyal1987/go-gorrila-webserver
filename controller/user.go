package controller

import (
     "fmt"
     "net/http"
)


type Controller struct {
}


func UserNew(w http.ResponseWriter,r *http.Request) {
    w.Header().Add("Content-Type", "text/html")
    fmt.Fprintf(w,"New User")
}
