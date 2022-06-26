package api

import (
	"fmt"
	"net/http"
)

func CryptMyWallet(w http.ResponseWriter, r *http.Request) {
	_ = r.URL.Query()
	fmt.Println("Placeholder for My Wallet")
}
