package plarpc

import (
	"github.com/op/go-logging"
	pla "github.com/rtuin/go-plalib"
	"os"
)

var log = logging.MustGetLogger("pla-rpc")

func SetupLogging() *logging.Logger {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	var format = logging.MustStringFormatter(`%{time:2006/01/02 15:04:05.000} [%{level}] %{message}`)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatter)
	logging.SetBackend(backendLeveled)
	logging.SetLevel(logging.DEBUG, "")

	pla.RegisterLogger(log)

	return log
}
