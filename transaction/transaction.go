package transaction

var (
	transactionSeperator = "\n\n==========================================\n"
	DefaultGas           = 200000
)

type Transaction interface {
	Execute() (string, error)
	Report()
	Prompt()
}

type Transactions []Transaction
