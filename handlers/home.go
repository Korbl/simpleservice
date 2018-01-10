package handlers

import (
	"fmt"
	"net/http"
)

// home is a simple HTTP handler function which writes a responce 

func home(w http.ResponseWrite, _ *http.Request) {
	fmt.Fprint(w, "Hello! Your request was processed")
}