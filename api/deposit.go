package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/notional-labs/cookiemonster/db"
)

func InitAPI() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/deposit", Deposit).Methods("POST")
	router.HandleFunc("/check-account", CheckAccount).Methods("POST")
	// router.HandleFunc("/deposit/get-address", CheckAccount)

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func CheckAccount(w http.ResponseWriter, r *http.Request) {
	acc := ""

	database := db.DefaultRegisteredAccountDB

	var address string
	for {

	}

}

func Deposit(w http.ResponseWriter, r *http.Request) {

	r.Header().Set("deposit-address", "")

}
