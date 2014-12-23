package controller

import (
	"fmt"
	"net/http"
)

func HomeHandler1(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "hello1")
}
