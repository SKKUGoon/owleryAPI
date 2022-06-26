package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"log"
	"net/http"
	"owlery/api"
	"time"
)

func init() {
	const INIT_ASCII = `
   ___         _              
  / _ \__ __ _| |___ _ _ _  _ 
 | (_) \ V  V / / -_) '_| || |
  \___/ \_/\_/|_\___|_|  \_, |
    _   ___ ___          |__/ 
   /_\ | _ \_ _|              
  / _ \|  _/| |               
 /_/ \_\_| |___|              

Hosting on http://localhost:8080
                              `
	fmt.Println(INIT_ASCII)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", api.Landing).Name("api landing")

	// sub-route 0. security
	subSec := r.PathPrefix("/api/v1/key").Subrouter()
	subSec.Path("/new").
		HandlerFunc(api.NewRegister).
		Methods("POST").
		Queries("username", "{username}").
		Name("api key creator")

	// sub-route 1. crypto
	subMktCrypt := r.PathPrefix("/api/v1/crypt").Subrouter()
	subMktCrypt.Path("/strat").
		HandlerFunc(api.CryptOrder).
		Methods("GET").
		Queries("apiKey", "{apiKey}").
		Name("crypto strategy return")
	subMktCrypt.Path("/wallet").
		HandlerFunc(api.CryptMyWallet).
		Methods("GET").
		Queries("apiKey", "{apiKey}").
		Name("my binance wallet status")

	// sub-route 2. korean market
	subMktKor := r.PathPrefix("/api/v1/kor").Subrouter()
	subMktKor.Path("/eventdriven").
		HandlerFunc(api.KorEventDriven).
		Methods("GET").
		Queries("apiKey", "{apiKey}").
		Name("korean market notice board for event driven strategy")
	subMktKor.Path("/rate").
		HandlerFunc(api.KorBond).
		Methods("GET").
		Queries("apiKey", "{apiKey}").
		Name("korean market korean bonds rates")

	// sub-route 3. foreign market

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
