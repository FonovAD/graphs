package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	LogInfo(op string, errors, infoField any)
	LogDebug(op string, errors, infoField any)
	LogWarning(op string, errors, infoField any)
	LogError(op string, errors, infoField any)
	LogFatal(op string, errors, infoField any)
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) LogInfo(op string, errors, infoField any) {
	logrus.WithFields(logrus.Fields{
		"op":     op,
		"errors": errors,
	}).Info(infoField)
}

func (l *logger) LogDebug(op string, errors, infoField any) {
	logrus.WithFields(logrus.Fields{
		"op":     op,
		"errors": errors,
	}).Debug(infoField)
}

func (l *logger) LogWarning(op string, errors, infoField any) {
	logrus.WithFields(logrus.Fields{
		"op":     op,
		"errors": errors,
	}).Warn(infoField)
}

func (l *logger) LogError(op string, errors, infoField any) {
	logrus.WithFields(logrus.Fields{
		"op":     op,
		"errors": errors,
	}).Error(infoField)
}

func (l *logger) LogFatal(op string, errors, infoField any) {
	logrus.WithFields(logrus.Fields{
		"op":     op,
		"errors": errors,
	}).Fatal(infoField)
}
