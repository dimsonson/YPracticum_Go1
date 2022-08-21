package main

import (
	"log"
	"os"
)

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(Warning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}

func (l *LogExtended) SetLogLevel(logLvl LogLevel) {
	l.logLevel = logLvl
}

func (l *LogExtended) Infoln(msg string) {
	l.println(Info, "INFO ", msg)
}

func (l *LogExtended) Warnln(msg string) {
	l.println(Warning, "WARN ", msg)
}

func (l *LogExtended) Errorln(msg string) {
	l.println(Error, "ERR ", msg)
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

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	if srcLogLvl < l.logLevel {
		return
	}

	l.Logger.Println(prefix + msg)
}

func NewLogExtended() *LogExtended {
	return &LogExtended{
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
		logLevel: Error,
	}
}

/* package main

import (
    "log"
    "os"
)

type LogLevel int

func (l LogLevel) IsValid() bool {
    switch l {
    case LogLevelInfo, LogLevelWarning, LogLevelError:
        return true
    default:
        return false
    }
}

const (
    LogLevelError LogLevel = iota
    LogLevelWarning
    LogLevelInfo
)

type LogExtended struct {
    *log.Logger
    logLevel LogLevel
}

func (l *LogExtended) SetLogLevel(logLvl LogLevel) {
    if !logLvl.IsValid() {
        return
    }
    l.logLevel = logLvl
}

func (l *LogExtended) Infoln(msg string) {
    l.println(LogLevelInfo, "INFO ", msg)
}

func (l *LogExtended) Warnln(msg string) {
    l.println(LogLevelWarning, "WARN ", msg)
}

func (l *LogExtended) Errorln(msg string) {
    l.println(LogLevelError, "ERR ", msg)
}

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
    if l.logLevel < srcLogLvl {
        return
    }

    l.Logger.Println(prefix + msg)
}

func NewLogExtended() *LogExtended {
    return &LogExtended{
        Logger:   log.New(os.Stderr, "", log.LstdFlags),
        logLevel: LogLevelError,
    }
}

func main() {
    logger := NewLogExtended()
    logger.SetLogLevel(LogLevelWarning)
    logger.Infoln("Не должно напечататься")
    logger.Warnln("Hello")
    logger.Errorln("World")
}  
*/
