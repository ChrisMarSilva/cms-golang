package main

// log "github.com/sirupsen/logrus"
// "os"

//"github.com/rs/zerolog"
//"time"

//"go.uber.org/zap"

// go get -u "github.com/Sirupsen/logrus"
// go get -u github.com/rs/zerolog/log
// go get -u go.uber.org/zap

func main() {

	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// log.Print("hello world")
	// log.Debug().
	// 	Str("Scale", "833 cents").
	// 	Float64("Interval", 833.09).
	// 	Msg("Fibonacci is everywhere")
	// log.Debug().
	// 	Str("Name", "Tom").
	// 	Send()

	// logger, _ := zap.NewProduction()
	// defer logger.Sync() // flushes buffer, if any
	// sugar := logger.Sugar()
	// sugar.Infow("failed to fetch URL",
	// 	// Structured context as loosely typed key-value pairs.
	// 	"url", "url",
	// 	"attempt", 3,
	// 	"backoff", time.Second,
	// )
	// sugar.Infof("Failed to fetch URL: %s", "url")

	// log.SetFormatter(&log.JSONFormatter{})
	//log.SetFormatter(&log.TextFormatter{})

	// log.SetFormatter(&log.TextFormatter{
	// 	DisableColors: true,
	// 	FullTimestamp: true,
	// })

	// log.SetReportCaller(true)

	// log.SetOutput(os.Stdout)
	// log.SetLevel(log.WarnLevel)
	// log.SetLevel(log.InfoLevel)

	// requestLogger := log.WithFields(log.Fields{"request_id": 123, "user_ip": 444})
	// requestLogger.Info("something happened on that request") //# will log request_id and user_ip
	// requestLogger.Warn("something not great happened")

	// log.WithFields(log.Fields{"animal": "walrus"}).Info("A walrus appears")
	// log.WithFields(log.Fields{"foo": "foo", "bar": "bar"})
	// log.Trace("Something very low level.")
	// log.Debug("Useful debugging information.")
	// log.Info("Something noteworthy happened!")
	// log.Warn("You should probably take a look at this.")
	// log.Error("Something failed but I'm not quitting.")
	// log.WithFields(log.Fields{
	// 	"animal": "walrus",
	// 	"size":   10,
	// }).Info("A group of walrus emerges from the ocean")

	// log.WithFields(log.Fields{
	// 	"omg":    true,
	// 	"number": 122,
	// }).Warn("The group's number increased tremendously!")

	// log.WithFields(log.Fields{
	// 	"omg":    true,
	// 	"number": 100,
	// }).Fatal("The ice breaks!")

	// // A common pattern is to re-use fields between logging statements by re-using
	// // the logrus.Entry returned from WithFields()
	// contextLogger := log.WithFields(log.Fields{
	// 	"common": "this is a common field",
	// 	"other":  "I also should be logged always",
	// })

	// contextLogger.Info("I'll be logged with common and other field")
	// contextLogger.Info("Me too")

	// log.Fatal("Bye.")
	// log.Panic("I'm bailing.")
}

// import (
// 	"log"
// 	"os"
// )

// var (
// 	WarningLogger *log.Logger
// 	InfoLogger    *log.Logger
// 	ErrorLogger   *log.Logger
// )

// func init() {
// 	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
// 	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
// 	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
// }

// func main() {
// 	InfoLogger.Println("Starting the application...")
// 	InfoLogger.Println("Something noteworthy happened")
// 	WarningLogger.Println("There is something you should know about")
// 	ErrorLogger.Println("Something went wrong")
// }
