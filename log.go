package rlog

import (
    "fmt"
    "os"
    "github.com/sirupsen/logrus"
    nested "github.com/cyberys/nested-logrus-formatter"
)

var (
    log *logrus.Logger
    f *os.File
    err error
)

func init() {

    logfile := os.Getenv("LOGFILE")
    if logfile == "" {
        f = os.Stdout
    } else {
        f, err = os.OpenFile(logfile, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0600)
    }
    if err != nil {
        fmt.Printf("error opening file: %v", err)
    }
//    defer f.Close()

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
            HideKeys:         true,
            NoColors:         true,
            ShowFullLevel:    true,
        },
        Level: ll,
    }
    log.SetOutput(f)

}

func Info(msg string, vars ...interface{}) {
    log.Infof(msg, vars...)
}

func Debug(msg string, vars ...interface{}) {
    log.Debugf(msg, vars...)
}

func Error(msg string, vars ...interface{}) {
    log.Errorf(msg, vars...)
}

func Trace(msg string, vars ...interface{}) {
    log.Tracef(msg, vars...)
}

func Warn(msg string, vars ...interface{}) {
    log.Warningf(msg, vars...)
}

func Fatal(msg string, vars ...interface{}) {
    log.Fatalf(msg, vars...)
}

func Level() string {
    return log.Level.String()
}
