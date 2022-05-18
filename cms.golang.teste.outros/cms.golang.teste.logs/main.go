package main

import (
	""
)

func main(){


}

/*

https://github.com/ztrue/tracerr

file, _ := os.Create('file.log')
log.SetOutput(file)
log.SetFlags(log.Ldate | log.Lshortfile)
log.SetFlags(log.LstdFlags | log.Lshortfile)
log.SetFlags(log.Lmicroseconds)
log.Println('Hello world')
file.Close()

log.SetOutput(&lumberjack.Logger{
  Filename:   "/var/log/proxy.log",
  MaxSize:    1000, // megabytes
  MaxBackups: 3,
  MaxAge:     1, // days
  Compress:   true, // disabled by default
})

const (
Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
Ltime                         // the time in the local time zone: 01:23:23
Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
Llongfile                     // full file name and line number: /a/b/c/d.go:23
Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
LstdFlags     = Ldate | Ltime // initial values for the standard logger

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
    return fmt.Print(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + " [DEBUG] " + string(bytes))
}

func main() {

    log.SetFlags(0)
    log.SetOutput(new(logWriter))
    log.Println("This is something being logged!")
}

*/