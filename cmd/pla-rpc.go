package main

import (
	"fmt"
	pla "github.com/rtuin/go-plalib"
	"net/http"
	// "os"
)

func main() {
	fmt.Println("Pla-RPC master by Richard Tuin - The remote procedure call version of Pla.\n")

	targets, err := pla.LoadTargets("Plafile.yml")
	if err != nil {
		panic(err)
	}

	http.ListenAndServe("localhost:7777", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Printf("Request for %v\n", req.RequestURI)
		fmt.Printf("Targets: %v\n", targets)

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
		pla.RunTargetByName(calledTarget, targets, false, params)
	}))
}
