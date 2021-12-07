package api

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/notional-labs/cookiemonster/accountmanager"
	"github.com/notional-labs/cookiemonster/db"
	"github.com/notional-labs/cookiemonster/query"
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
	var err error
	for {
		address, err = db.GetRegisterAddressForAddress(database, acc)
		if err == nil {
			break
		}
	}

	w.Header().Set("address", address)

}

func Deposit(w http.ResponseWriter, r *http.Request) {
	txHash := ""
	time.Sleep(5 * time.Second)
	res, err := query.QueryTxWithRetry(txHash, 5)
	if err != nil {
		w.Header().Set("error", err.Error())
	} else if res.Code != 0 {
		w.Header().Set("error", "tx fail with code "+strconv.Itoa(int(res.Code)))
	} else {
		w.Header().Set("fund recieved", "true")
	}
	acc := ""

	am := accountmanager.DefaultAccountManager

	privKey, err := am.CreateNewPrivKeyForAddress(acc)

	privKey.PubKey().Address()

}
