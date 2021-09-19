package wallet

type WalletLoader interface {
	loadWallets() ([]Wallet, error)
}

type JsonWalletLoader struct {
	filePath string
}

func (x JsonWalletLoader) loadWallets() ([]Wallet, error) {
	// open filepath
	var wallets []Wallet = make([]Wallet, 0)
	return wallets, nil
}
