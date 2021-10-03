package osmosis

import (
	"os"
	"path/filepath"
)

var (
	ChainId = "osmosis-1"
	Node    = "https://osmosis-1.technofractal.com:443"
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
