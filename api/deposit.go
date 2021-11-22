package api_gateway

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/notional-labs/cookiemonster/cmd/auto-farm/cmd"
	"github.com/notional-labs/cookiemonster/command/query"
)

func InitAPI() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/deposit/", Deposit).Methods("POST")
	router.HandleFunc("/deposit/get-address", GetAddress)


	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServeTLS(":8000",router))
}

func Deposit(w http.ResponseWriter, r *http.Request) {
	// checking header type to make sure json
	if r.Header.Get("Content-Type") != "application/json" {
		msg := "Content-Type header is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}
	txHash := r.Header.Get("tx-hash")
	cmd := cmd.NewRootCmd()
	txQueryRespond, err := query.QueryTxWithRetry(,txHash)
	
	
	w.Header().Set("deposit-address", "")
	w.


	
}

func GetAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("address", "sl;adkjfl;ksdjaf")
	accountmanager.

}