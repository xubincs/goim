package log

import "github.com/sirupsen/logrus"

func Init(debug bool) {
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func Info(a ...interface{}) {
	logrus.Info(a...)
}

func Infof(format string, a ...interface{}) {
	logrus.Infof(format, a...)
}

func Errorf(format string, a ...interface{}) {
	logrus.Errorf(format, a...)
}

func Error(a ...interface{}) {
	logrus.Error(a...)
}

func Debug(a ...interface{}) {
	logrus.Debug(a...)
}

func Debugf(format string, a ...interface{}) {
	logrus.Debugf(format, a...)
}

func Warn(a ...interface{}) {
	logrus.Warn(a...)
}

func Warnf(format string, a ...interface{}) {
	logrus.Warnf(format, a...)
}