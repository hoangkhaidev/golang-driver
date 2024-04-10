package log

import (
	"runtime"

	log "github.com/sirupsen/logrus"
)

func Error(errMsg string) {
	_, file, line, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	  }).Error(`==============> `, errMsg)
}
