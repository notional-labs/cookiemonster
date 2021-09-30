package transaction

type Transaction interface {
	Execute() (string, error)
	Report()
	Prompt()
	// Type() string
}

type Transactions []Transaction
