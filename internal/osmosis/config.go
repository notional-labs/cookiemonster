package osmosis

import (
	"os"
	"path/filepath"
)

var (
	ChainId = "test"
	Node    = "http://0.0.0.0:26657"
	HomeDir = DefaultNodeHome()
)

func DefaultNodeHome() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome := filepath.Join(userHomeDir, ".osmosisd")
	return DefaultNodeHome
}
