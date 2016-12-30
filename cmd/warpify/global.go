package warpify

import "github.com/pressly/warpdrive/config"

type warpConf struct {
	bundleVersion string
	bundlePath    string
	documentPath  string
	platform      string
	defaultCycle  string
	forceUpdate   bool
	reloadTask    Callback
	api           *api
	warpFile      *config.ClientConfig
	reloadReady   chan struct{}
}

// conf is a global warpConf for this package
var conf warpConf
