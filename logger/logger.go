package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"path/filepath"
)

type Formatter int

const (
	DefaultFormatter Formatter = iota
	JsonFormatter
	TextFormatter
)

type Output int

const (
	DefaultOutput Output = iota
	StandardOutput
	FileOutput
)

type Level int

const (
	DefaultLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

type ILogger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

type logger struct {
	lrus     *logrus.Logger
	format   Formatter
	output   Output
	dir      string
	fileName string
	level    Level
}

type Option func(*logger) error

func WithFormatter(formatter Formatter) Option {
	return func(l *logger) error {
		l.format = formatter
		return nil
	}
}

func WithFileOutput(directory string, fileName string) Option {
	return func(l *logger) error {
		l.output = FileOutput
		l.dir = directory
		l.fileName = fileName
		return nil
	}
}

func WithLevel(level Level) Option {
	return func(l *logger) error {
		l.level = level
		return nil
	}
}

func NewInstance(opts ...Option) (ILogger, error) {
	// Create default instance
	instance := &logger{
		lrus:   logrus.New(),
		format: DefaultFormatter,
		output: DefaultOutput,
		level:  DefaultLevel,
	}

	for _, opt := range opts {
		err := opt(instance)
		if err != nil {
			return nil, err
		}
	}

	switch instance.format {
	case JsonFormatter:
		formatter := logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
		instance.lrus.SetFormatter(&formatter)
		break
	default:
		formatter := new(logrus.TextFormatter)
		formatter.TimestampFormat = "2006-01-02 15:04:05"
		formatter.FullTimestamp = true
		instance.lrus.SetFormatter(formatter)
	}

	switch instance.output {
	case FileOutput:
		instance.lrus.SetOutput(&lumberjack.Logger{
			Filename:   filepath.Join(instance.dir, instance.fileName),
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,    //days
			Compress:   false, // disabled by default
		})
	default:
		// log to the terminal
	}

	instance.lrus.SetLevel(logrus.InfoLevel)

	switch instance.level {
	case FatalLevel:
		instance.lrus.SetLevel(logrus.FatalLevel)
		break
	case ErrorLevel:
		instance.lrus.SetLevel(logrus.ErrorLevel)
		break
	case WarnLevel:
		instance.lrus.SetLevel(logrus.WarnLevel)
		break
	case InfoLevel:
		instance.lrus.SetLevel(logrus.InfoLevel)
		break
	case DebugLevel:
		instance.lrus.SetLevel(logrus.DebugLevel)
		break
	case TraceLevel:
		instance.lrus.SetLevel(logrus.TraceLevel)
		break
	default:
		instance.lrus.SetLevel(logrus.InfoLevel)
	}

	return instance, nil
}

func (l *logger) Info(args ...interface{}) {
	l.lrus.Info(args)
}
func (l *logger) Infof(format string, args ...interface{}) {
	l.lrus.Infof(format, args)
}
func (l *logger) Debug(args ...interface{}) {
	l.lrus.Debug(args)
}
func (l *logger) Debugf(format string, args ...interface{}) {
	l.lrus.Debugf(format, args)
}
func (l *logger) Warn(args ...interface{}) {
	l.lrus.Warn(args)
}
func (l *logger) Warnf(format string, args ...interface{}) {
	l.lrus.Warnf(format, args)
}
func (l *logger) Error(args ...interface{}) {
	l.lrus.Error(args)
}
func (l *logger) Errorf(format string, args ...interface{}) {
	l.lrus.Errorf(format, args)
}
