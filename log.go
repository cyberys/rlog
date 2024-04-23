package rlog

import (
    "fmt"
    "os"
    "sync"

    "github.com/sirupsen/logrus"
    nested "github.com/cyberys/nested-logrus-formatter"
)

var (
    log *logrus.Logger
    f *os.File
    err error
    queueNumber int
    queueNumberSet sync.Once
)

func init() {
    logfile := os.Getenv("LOGFILE")
    if logfile == "" {
        f = os.Stdout
    } else {
        f, err = os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
        if err != nil {
            fmt.Printf("Error opening log file %s, defaulting to standard output: %v\n", logfile, err)
            f = os.Stdout
        }
    }

    lvl, ok := os.LookupEnv("RLOG_LEVEL")
    if !ok {
        lvl = "info"
    }
    ll, err := logrus.ParseLevel(lvl)
    if err != nil {
        ll = logrus.DebugLevel
    }

    log = &logrus.Logger{
        Out: f,
        Formatter: &nested.Formatter{
            TimestampFormat: "2006-01-02 15:04:05",
            HideKeys:        true,
            NoColors:        true,
            ShowFullLevel:   true,
        },
        Level: ll,
    }
}

func SetQueueNumber(num int) {
    queueNumberSet.Do(func() {
        queueNumber = num
    })
}

func Info(msg string, vars ...interface{}) {
    log.WithField("queue", queueNumber).Infof(msg, vars...)
}

func Debug(msg string, vars ...interface{}) {
    log.WithField("queue", queueNumber).Debugf(msg, vars...)
}

func Error(msg string, vars ...interface{}) {
    log.WithField("queue", queueNumber).Errorf(msg, vars...)
}

func Trace(msg string, vars ...interface{}) {
    log.WithField("queue", queueNumber).Tracef(msg, vars...)
}

func Warn(msg string, vars ...interface{}) {
    log.WithField("queue", queueNumber).Warningf(msg, vars...)
}

func Fatal(msg string, vars ...interface{}) {
    log.WithField("queue", queueNumber).Fatalf(msg, vars...)
}

func Level() string {
    return log.Level.String()
}
