// The MIT License (MIT)

// Copyright (c) 2016 Richard Tuin

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package plarpc

import (
	"encoding/json"
	pla "github.com/rtuin/go-plalib"
	"net/http"
)

type PlaHttpResponse struct {
	Message string `json:"message"`
}

func ServePlaRpc(config Config) {
	log.Debug("Starting server...")

	targets, err := pla.LoadTargets("Plafile.yml")
	if err != nil {
		log.Fatalf("Cannot start pla-rpc: %v", err.Error())
	}

	infinite := make(chan bool)

	go func() {
		http.ListenAndServe(config.BindAddress, http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
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

	log.Infof("Server running at http://%s/", config.BindAddress)
	<-infinite
}
