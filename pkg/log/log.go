package log

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init(logFile string) {
	logFormatter := new(log.JSONFormatter)

	log.SetFormatter(logFormatter)
	log.SetLevel(log.InfoLevel)

	if logFile == "" {
		log.SetOutput(os.Stdout)
		return
	}

	logPath, err := filepath.Abs(logFile)
	if err != nil {
		panic(err)
	}
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     30, // days
		Compress:   false,
	}
	log.SetOutput(lumberjackLogger)
}

func SetDebug(debug bool) {
	if debug {
		log.SetLevel(log.DebugLevel)
	}
}
