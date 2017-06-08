package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"

	cfg "../../property/util"
	log "../logrus"
)

var l = false
var mu = &sync.Mutex{}

func GetLog() {
	if !l {
		initLog()
	}
}

func initLog() {
	mu.Lock()
	defer mu.Unlock()
	if !l {
		log.SetFormatter(&log.JSONFormatter{})

		dir := cfg.GetValue("log", "log.path")
		logName := cfg.GetValue("log", "log.filename")
		bufStr := cfg.GetValue("log", "log.buffer")
		bufSize, atoiErr := strconv.Atoi(bufStr)
		if atoiErr != nil {
			bufSize = 1
		}
		_, err := os.Open(dir)
		if err != nil {
			os.MkdirAll(dir, os.ModeDir)

		}
		file, err := os.OpenFile(dir+"/"+logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

		if err == nil {
			log.SetOutput(bufio.NewWriterSize(file, bufSize))
		} else {
			log.SetOutput(os.Stdout)
			log.Error("Failed to load file, using default stderr")
		}
		//配置文件
		//ErrorLevel
		level := cfg.GetValue("log", "log.level")

		le := log.ErrorLevel
		if strings.EqualFold(level, "info") {
			le = log.InfoLevel
		}
		if strings.EqualFold(level, "debug") {
			le = log.DebugLevel
		}
		if strings.EqualFold(level, "error") {
			le = log.ErrorLevel
		}
		if strings.EqualFold(level, "warn") {
			le = log.WarnLevel
		}
		log.SetLevel(le)
		l = true
	}
}

/*
// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
*/

func Error(val ...interface{}) {
	GetLog()
	log.Error(val)
}

func Info(val ...interface{}) {
	GetLog()
	log.Info(val)
}
func Warn(val ...interface{}) {
	GetLog()
	log.Warn(val)
}

func Debug(val ...interface{}) {
	GetLog()
	log.Debug(val)
}
