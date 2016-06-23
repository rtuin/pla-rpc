package plarpc

import (
	"encoding/json"
	pla "github.com/rtuin/go-plalib"
	"net/http"
)

type PlaHttpResponse struct {
	Message string `json:"message"`
}

func ServePlaRpc() {
	log.Debug("Starting server...")

	targets, err := pla.LoadTargets("Plafile.yml")
	if err != nil {
		log.Fatalf("Cannot start pla-rpc: %v", err.Error())
	}

	infinite := make(chan bool)

	go func() {
		http.ListenAndServe("localhost:7777", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			var responseMessage *PlaHttpResponse
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
				switch err.Code {
				case pla.TARGET_RUN_ERROR:
					res.WriteHeader(http.StatusInternalServerError)
					responseMessage = &PlaHttpResponse{err.Error.Error()}
					break
				case pla.TARGET_NOT_FOUND:
					res.WriteHeader(http.StatusNotFound)
					responseMessage = &PlaHttpResponse{err.Error.Error()}
					break
				}
			}

			if responseMessage == nil {
				res.WriteHeader(http.StatusNoContent)
				return
			}

			json, jsonError := json.Marshal(*responseMessage)
			if jsonError != nil {
				log.Error(jsonError)
				http.Error(res, "Something went wrong inside pla-rpc, please check the logs and file a bug-report.", http.StatusInternalServerError)
			}
			res.Header().Add("Content-Type", "application/json")
			res.Write(json)
		}))
	}()

	log.Infof("Server running at http://localhost:7777/")
	<-infinite
}
