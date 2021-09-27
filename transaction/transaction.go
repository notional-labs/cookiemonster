package transaction

var (
	transactionSeperator = "\n\n==========================================\n"
)

type Transaction interface {
	Execute() error
	Report()
}
