package logger

import "fmt"

type ILogger interface {
	Info()
	Infof()
	Debug()
	Debugf()
	Warn()
	Warnf()
	Error()
	Errorf()
}

type logger struct {
	//log *logrus.Logger
}

func NewInstance() (ILogger, error) {
	//log = logrus.New()
	//formatter := new(logrus.TextFormatter)
	//formatter.TimestampFormat = "2006-01-02 15:04:05"
	//formatter.FullTimestamp = true
	//log.SetFormatter(formatter)
	//
	//etrackerweb.SetLogger(log.WithField("p", "WEB"))
	//mysqlclient.SetLogger(log.WithField("p", "MYSQL"))
	//
	//switch conf.Log.Formatter {
	//case 1:
	//	// jsonFormatter
	//	formatter := logrus.JSONFormatter{
	//		TimestampFormat: "2006-01-02 15:04:05",
	//	}
	//	log.SetFormatter(&formatter)
	//	break
	//default:
	//	formatter := new(logrus.TextFormatter)
	//	formatter.TimestampFormat = "2006-01-02 15:04:05"
	//	formatter.FullTimestamp = true
	//	log.SetFormatter(formatter)
	//}
	//switch conf.Log.Output {
	//case 1:
	//	// log to file
	//	log.SetOutput(&lumberjack.Logger{
	//		Filename:   filepath.Join(conf.System.LogDir, "etracker.log"),
	//		MaxSize:    500, // megabytes
	//		MaxBackups: 3,
	//		MaxAge:     28,    //days
	//		Compress:   false, // disabled by default
	//	})
	//default:
	//	// log to the terminal
	//}
	//var level = logrus.PanicLevel
	//switch conf.Log.Level {
	//case 1:
	//	level = logrus.FatalLevel
	//	break
	//case 2:
	//	level = logrus.ErrorLevel
	//	break
	//case 3:
	//	level = logrus.WarnLevel
	//	break
	//case 4:
	//	level = logrus.InfoLevel
	//	break
	//case 5:
	//	level = logrus.DebugLevel
	//	break
	//case 6:
	//	level = logrus.TraceLevel
	//	break
	//}
	//log.SetLevel(level)

	return &logger{}, nil
}

func (l *logger) Info() {
	fmt.Println("info")
}
func (l *logger) Infof() {
	fmt.Println("infof")
}
func (l *logger) Debug() {
	fmt.Println("debug")
}
func (l *logger) Debugf() {
	fmt.Println("debugf")
}
func (l *logger) Warn() {
	fmt.Println("warn")
}
func (l *logger) Warnf() {
	fmt.Println("warnf")
}
func (l *logger) Error() {
	fmt.Println("error")
}
func (l *logger) Errorf() {
	fmt.Println("errorf")
}
