package log

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

var log = logrus.New()
var notebookMode bool

func init() {
	log.SetOutput(ioutil.Discard)
}

func NotebookMode() {
	notebookMode = true
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})
}

func Get() *logrus.Logger {
	if notebookMode {
		log.SetOutput(os.Stdout)
	}
	return log
}
