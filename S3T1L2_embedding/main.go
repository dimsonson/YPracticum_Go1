package main

import (
	"fmt"
	"log"
)

func main() {
	logger := new(LogExtended)
	logger.SetLogLevel(Warning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	//logger.Println("Debug")
	fmt.Println(Info)
}

func (l *LogExtended) SetLogLevel(logLvl LogLevel) {
	l.logLevel = logLvl
}

func (l *LogExtended) Infoln(msg string) {
	l.println(0, "INFO", msg)
}

func (l *LogExtended) Warnln(msg string) {
	l.println(1, "WARN", msg)
}

func (l *LogExtended) Errorln(msg string) {
	l.println(2, "ERR", msg)
}

type LogLevel int

const (
	Info LogLevel = iota
	Warning
	Error
)

//var p LogExtended

type LogExtended struct {
	*log.Logger
	logLevel LogLevel // LogLevel это enum
}

/* func (l LogLevel) enumIndex() int {
	return int(l)
} */

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	if srcLogLvl < l.logLevel {
		return
	}

	l.Logger.Println(prefix + msg)
}
