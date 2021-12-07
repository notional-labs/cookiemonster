package osmosis

import (
	"os"
	"path/filepath"
)

var (
	ChainId = "osmosis-1"
	Node    = "http://162.55.132.230:2001"
	HomeDir = DefaultNodeHome()
)

func DefaultNodeHome() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	DefaultNodeHome := filepath.Join(userHomeDir, ".auto-invest")
	return DefaultNodeHome
}
