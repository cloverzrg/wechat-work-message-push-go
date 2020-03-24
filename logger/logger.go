package logger

import (
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

var Entry = logrus.NewEntry(Logger)

var Error = Entry.Error
var Errorf = Entry.Errorf
var Errorln = Entry.Errorln

var Info = Entry.Info
var Infof = Entry.Infof

var Print = Entry.Info
var Printf = Entry.Infof
var Println = Entry.Println

var Debug = Entry.Debug
var Debugf = Entry.Debugf
var Debugln = Entry.Debugln

var Panicf = Entry.Panicf
var Panic = Entry.Panic

var Trace = Entry.Trace
var Tracef = Entry.Tracef

var Warn = Entry.Warn
var Warnf = Entry.Warnf

var Fatal = Entry.Fatal
var Fatalf = Entry.Fatalf

func init() {
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Logger.SetReportCaller(true)
}
