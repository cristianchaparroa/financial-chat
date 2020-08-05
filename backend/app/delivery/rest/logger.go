package rest

import (
	"fmt"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

const (
	loggerFile = "/tmp/logger.log"
)

func setupLogger(s *server) {

	if !fileExists(loggerFile) {
		createFile()
	}

	// setup logger
	lumberjackLogRotate := &lumberjack.Logger{
		Filename:   loggerFile,
		MaxSize:    5,  // Max megabytes before log is rotated
		MaxBackups: 90, // Max number of old log files to keep
		MaxAge:     60, // Max number of days to retain log files
		Compress:   true,
	}

	mw := io.MultiWriter(os.Stdout, lumberjackLogRotate)
	log.SetOutput(mw)
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.JSONFormatter{})

	setupGinLogger(s, mw)

	/* This configuration is just in case you want to log everything to Logstash
	 and view the logs with Kibana
	logStashServer := os.Getenv("LOG_STASH_SERVER")
	hook := graylog.NewGraylogHook(logStashServer, map[string]interface{}{})
	log.AddHook(hook) */
}

// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func createFile() {
	_, err := os.Create(loggerFile)
	if err != nil {
		fmt.Print(err)
		log.Println(err)
	}
}

func setupGinLogger(s *server, mw io.Writer) {
	l := log.New()
	l.SetOutput(mw)
	l.SetLevel(log.TraceLevel)
	l.SetFormatter(&log.JSONFormatter{})
	s.router.Use(ginrus.Ginrus(l, time.RFC3339, false))
}
