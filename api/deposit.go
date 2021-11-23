package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/notional-labs/cookiemonster/accountmanager"
	"github.com/notional-labs/cookiemonster/invest"
)

func InitAPI() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/deposit/register-account", RegisterAccount).Methods("POST")
	router.HandleFunc("/deposit/", Deposit).Methods("POST")
	router.HandleFunc("/deposit/get-address", GetAddress)

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServeTLS(":8000", "localhost.crt", "localhost.key", router))
}

func RegisterAccount(w http.ResponseWriter, r *http.Request) {

}

func Deposit(w http.ResponseWriter, r *http.Request) {
	// // checking header type to make sure json
	// if r.Header.Get("Content-Type") != "application/json" {
	// 	msg := "Content-Type header is not application/json"
	// 	http.Error(w, msg, http.StatusUnsupportedMediaType)
	// 	return
	// }
	// txHash := r.Header.Get("tx-hash")

	// txQueryRespond, err := query.QueryTxWithRetry(txHash)
	// code := txQueryRespond.Code

	// if code != 0 {
	// 	panic(fmt.Errorf("deposit failed"))
	// }

	// tx := txQueryRespond.GetTx()
	// sender := tx.GetMsgs()[0].GetSigners()

	// w.Header().Set("deposit-address", "")

}

func GetAddress(w http.ResponseWriter, r *http.Request) {
	am, err := accountmanager.LoadAccountManagerFromFile("~/accountmanager.json")
	addr := am.CreateNewAccount()
	if err != nil {
		panic(err)
	}
	accountmanager.DumpAccountManagerToFile(am, "~/accountmanager.json")
	investments := am.GetDefaultInvestments()
	invest.DumpInvestmentsToFile("~/cookiemonster/investments.json", investments)

	w.Header().Set("deposit-address", addr)
}
