package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gorilla/mux"
	"github.com/notional-labs/cookiemonster/accountmanager"
	"github.com/notional-labs/cookiemonster/db"
	"github.com/notional-labs/cookiemonster/invest"
	"github.com/notional-labs/cookiemonster/query"
	"github.com/notional-labs/cookiemonster/transaction"
)

func InitAPI() {
	fmt.Println(5)
	router := mux.NewRouter().StrictSlash(true)
	db.DefaultAddressToCMKeyNameDB = db.AddressToCMKeyDB{
		DB: db.MustOpenDB(db.DefaultAddressToCMKeyNameDBDir),
	}

	db.DefautlAddressToCMAddressDB = db.AddressToCMAddressDB{
		DB: db.MustOpenDB(db.DefaultAddressToCMAddressDBDir),
	}

	accountmanager.DefaultAccountManager = *accountmanager.MustLoadAccountManagerFromFile("/.cookiemonster/accountmanager.json")

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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	m := &map[string]string{}
	json.Unmarshal(body, m)

	userAddress := (*m)["address"]

	addressToCMKeyDB := db.DefaultAddressToCMKeyNameDB
	cmKeyForUserAddress, err := addressToCMKeyDB.GetCMKeyNameForAddress(userAddress)

	if err != nil {
		panic(err)
	}
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	m := &map[string]string{}
	json.Unmarshal(body, m)
	fmt.Println(0)
	userAddress := (*m)["address"]
	poolId := (*m)["pool-id"]
	fmt.Println(userAddress)
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

	fmt.Println(0)

	err = investment.InvestWithOutClaim()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	fmt.Println(0)

	w.WriteHeader(200)
}

type AddressResponse struct {
	Address string
}

func CheckAccount(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	m := &map[string]string{}
	json.Unmarshal(body, m)

	userAddress := (*m)["address"]

	addressToCMAddressDB := db.DefautlAddressToCMAddressDB

	var cmAddress string
	for {
		cmAddress, err = addressToCMAddressDB.GetCMAddressForAddress(userAddress)
		if err == nil {
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	addr := AddressResponse{
		Address: cmAddress,
	}

	json.NewEncoder(w).Encode(addr)
	w.WriteHeader(200)

}

func Deposit(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	m := &map[string]string{}
	json.Unmarshal(body, m)
	// desiredValue := m["tx-hash"]
	fmt.Println((*m)["tx-hash"])
	res, err := query.QueryTxWithRetry((*m)["tx-hash"], 5)
	fmt.Println(9)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	} else if res.Code != 0 {
		// fmt.Println(1)

		w.WriteHeader(http.StatusNotFound)
		panic("ddd")
	}
	// fmt.Println(1)

	tx := res.GetTx()
	msg := tx.GetMsgs()[0]

	bankMsg := msg.(*banktypes.MsgSend)

	if bankMsg.ToAddress != accountmanager.DefaultAccountManager.MasterAddress {
		w.WriteHeader(http.StatusNotFound)
		panic("wrong deposit address")
	}

	acc := bankMsg.FromAddress
	amount := bankMsg.Amount[0]

	am := accountmanager.DefaultAccountManager
	cmAddress, err := db.DefautlAddressToCMAddressDB.GetCMAddressForAddress(acc)
	if err != nil {
		panic(err)
	}

	var cmAddressBz sdk.AccAddress
	fmt.Println(cmAddress)
	if cmAddress == "" {
		cmAddressBz, err = am.RegisterAccountForAddress(acc)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			panic(err)

		}
	} else {
		cmAddressBz, _ = sdk.AccAddressFromBech32(cmAddress)
	}
	fmt.Println(cmAddressBz)
	sendFundToCmAccountAddressOfUser := transaction.BankSendOption{
		ToAddr: cmAddressBz,
		Denom:  "uosmo",
		Amount: amount.Amount,
	}
	fmt.Println("len of priv key", len(am.MasterKey))
	bankSendTx := transaction.BankSendTx{
		BankSendOpt: sendFundToCmAccountAddressOfUser,
		KeyName:     accountmanager.MasterKey,
	}

	err = transaction.HandleTx(bankSendTx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	addr := AddressResponse{
		Address: cmAddressBz.String(),
	}
	json.NewEncoder(w).Encode(addr)

	w.WriteHeader(200)

}
