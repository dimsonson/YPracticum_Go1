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
	println(0, "INFO", msg)
}

func (logE *LogExtended) Warnln(msg string) {
	println(1, "WARN", msg)
}

func (logE *LogExtended) Errorln(msg string){
	println(2, "ERR", msg)
}

type LogLevel int

const (
	Info LogLevel = iota
	Warning
	Error
)

type LogExtended struct {
	*log.Logger
	logLevel LogLevel // LogLevel это enum
}

func (l LogLevel) enumIndex() int {
	return int(l)
}

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.logLevel <= srcLogLvl {
		return
	}

	l.Logger.Println(prefix + msg)
}
