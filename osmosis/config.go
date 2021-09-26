package osmosis

import (
	"os"
	"path/filepath"
)

var (
	ChainId = "osmosis-1"
	Node    = "http://95.217.196.54:2001"
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
