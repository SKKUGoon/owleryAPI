package api

import (
	"fmt"
	"net/http"
)

func Landing(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Landing Page for Owlery API")
	if err != nil {
		fmt.Println("Server not responding")
	}
}
