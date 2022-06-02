package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debugf(format string, args ...interface{})
}

type Logbook struct {
	log *logrus.Entry
}

var Log = Logbook{}

func LogInit(logbookName string) error {
	var ret error
	var log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	logbookPath := fmt.Sprintf("/var/log/%s.log", logbookName)
	f, err := os.OpenFile(logbookPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = f
		logrus.Infof("Logging to %s", logbookPath)
	} else {
		logrus.Errorf("Failed to open %s, Err: %v", logbookPath, err)
		ret = err
	}
	hostName, _ := os.Hostname()
	Log.log = log.WithFields(logrus.Fields{"HostID": hostName})
	return ret
}

func (logger Logbook) Debugf(msg string, params ...interface{}) {
	if logger.log == nil {
		logrus.Errorf("Logger not initialized properly, LogInit hasn't finished successfully. Original msg: %s", fmt.Sprintf("%s, %v", msg, params))
		return
	}
	logger.log.Debugf(msg, params...)
}
