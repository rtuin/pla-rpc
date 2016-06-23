package main

import (
	. "github.com/rtuin/pla-rpc"
)

func main() {
	var log = SetupLogging()
	log.Infof("Pla-RPC master by Richard Tuin\n")

	ServePlaRpc()
}
