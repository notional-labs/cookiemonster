package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gorilla/mux"
	"github.com/notional-labs/cookiemonster/accountmanager"
	"github.com/notional-labs/cookiemonster/db"
	"github.com/notional-labs/cookiemonster/invest"
	"github.com/notional-labs/cookiemonster/query"
	"github.com/notional-labs/cookiemonster/transaction"
)

func InitAPI() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/deposit", Deposit).Methods("POST")
	router.HandleFunc("/check-account", CheckAccount).Methods("POST")
	// router.HandleFunc("/identify-pool", IdentifyPool)
	router.HandleFunc("/auto-investing", AutoInvest).Methods("POST")
	router.HandleFunc("/pull-reward", PullReward).Methods("POST")

	// router.HandleFunc("/deposit/get-address", CheckAccount)

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func PullReward(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	userAddress := r.Form.Get("address")

	addressToCMKeyDB := db.DefaultAddressToCMKeyNameDB
	cmKeyForUserAddress, err := addressToCMKeyDB.GetCMKeyNameForAddress(userAddress)

	claimTx := transaction.ClaimTx{
		KeyName: cmKeyForUserAddress,
	}

	err = transaction.HandleTx(claimTx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(200)
}

func AutoInvest(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	userAddress := r.Form.Get("address")
	poolId := r.Form.Get("pool-id")

	addressToCMKeyDB := db.DefaultAddressToCMKeyNameDB
	cmKeyForUserAddress, err := addressToCMKeyDB.GetCMKeyNameForAddress(userAddress)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}

	investment := invest.Investment{
		KeyName:         cmKeyForUserAddress,
		PoolPercentage:  100,
		StakePercentage: 0,
		PoolStrategy: invest.PoolStrategy{
			Distribution: map[string]int{poolId: 100},
		},
		Duration:     "14days",
		StakeAddress: "",
	}

	err = investment.InvestWithOutClaim()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}

	w.WriteHeader(200)
}

type Address struct {
	address string
}

func CheckAccount(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	userAddress := r.Form.Get("address")

	addressToCMAddressDB := db.DefaultAddressToCMAddressDB

	var cmAddress string
	var err error
	for {
		cmAddress, err = addressToCMAddressDB.GetCMAddressForAddress(userAddress)
		if err == nil {
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	addr := Address{
		address: cmAddress,
	}

	json.NewEncoder(w).Encode(addr)
	w.WriteHeader(200)

}

func Deposit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	txHash := r.Form.Get("tx-hash")
	time.Sleep(5 * time.Second)
	res, err := query.QueryTxWithRetry(txHash, 5)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	} else if res.Code != 0 {
		w.WriteHeader(http.StatusNotFound)
		panic("")
	}

	tx := res.GetTx()
	msg := tx.GetMsgs()[0]

	bankMsg := msg.(*banktypes.MsgSend)

	acc := bankMsg.FromAddress
	amount := bankMsg.Amount[0]

	am := accountmanager.DefaultAccountManager

	cmAddressBz, err := am.RegisterAccountForAddress(acc)

	sendFundToCmAccountAddressOfUser := transaction.BankSendOption{
		ToAddr: cmAddressBz,
		Denom:  "uosmo",
		Amount: amount.Amount,
	}

	bankSendTx := transaction.BankSendTx{
		BankSendOpt: sendFundToCmAccountAddressOfUser,
		KeyName:     "master",
	}

	err = transaction.HandleTx(bankSendTx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	addr := Address{
		address: cmAddressBz.String(),
	}
	json.NewEncoder(w).Encode(addr)

	w.WriteHeader(200)

}
