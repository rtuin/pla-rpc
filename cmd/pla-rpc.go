package main

import (
	pla "github.com/rtuin/go-plalib"
	. "github.com/rtuin/pla-rpc"
	"net/http"
)

func main() {
	var log = SetupLogging()
	log.Infof("Pla-RPC master by Richard Tuin - Starting server.\n")

	targets, err := pla.LoadTargets("Plafile.yml")
	if err != nil {
		panic(err)
	}

	http.ListenAndServe("localhost:7777", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Debugf("Request for %v\n", req.RequestURI)
		res.Header().Add("Content-Type", "application/json")

		// args := os.Args[1:]
		// calledTarget := "all"
		// if len(args) > 0 {
		// 	calledTarget = args[0]
		// }
		calledTarget := req.RequestURI[1:]

		// var params []string
		// if len(args) > 1 {
		// params = args[1:]
		// }

		var params []string
		var err = pla.RunTargetByName(calledTarget, targets, false, params)
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
		}
	}))
}
