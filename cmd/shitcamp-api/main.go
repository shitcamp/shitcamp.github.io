package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shitcamp-unofficial/shitcamp/pkg/config"
	"github.com/shitcamp-unofficial/shitcamp/pkg/log"
	"github.com/shitcamp-unofficial/shitcamp/pkg/server"
	logger "github.com/sirupsen/logrus"
)

func main() {
	var configFile, logFile string
	flag.StringVar(&configFile, "config", "", "config file")
	flag.StringVar(&logFile, "log", "", "log file to save to")
	flag.Parse()

	log.Init(logFile)

	err := config.LoadConfig(configFile)
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.ReleaseMode)
	if config.GetConfig().Debug {
		gin.SetMode(gin.DebugMode)
	}

	log.SetDebug(config.GetConfig().Debug)

	s := server.New()

	go func() {
		time.Sleep(1 * time.Second)

		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh

		err = s.Stop()
		if err != nil {
			logger.WithError(err).Error("stopping_server")
			fmt.Println("Error shutting down server:", err)
		}

		logger.Info("stopped_server")
	}()

	err = s.Start()
	if err != nil {
		panic(fmt.Sprintf("Error starting server:", err))
	}
}
