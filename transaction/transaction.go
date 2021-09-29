package transaction

var (
	transactionSeperator = "\n\n==========================================\n"
	DefaultGas           = 200000
)

type Transaction interface {
	Execute()
	Report()
	Prompt()
}

type Transactions []Transaction
