package flannel

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	loggerError  = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.LUTC|log.Ltime|log.Lshortfile)
	loggerAccess = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LUTC|log.Ltime)
)

func logOut(logger *log.Logger, level string, r *http.Request, format string, v ...interface{}) {
	logger.SetPrefix(fmt.Sprintf("[%6s] %s ", level, reqID(r)))
	logger.Output(3, fmt.Sprintf(format, v...))
}

func logAccess(r *http.Request, format string, v ...interface{}) {
	logOut(loggerAccess, "access", r, format, v...)
}

// SetErrorLogOutput sets the output writer for the access log (os.Stderr by
// default).
func SetErrorLogOutput(w io.Writer) {
	loggerError.SetOutput(w)
}

// SetAccessLogOutput sets the output writer for the access log (os.Stdout by
// default).
func SetAccessLogOutput(w io.Writer) {
	loggerAccess.SetOutput(w)
}

// LogInfo logs an info level error.
func LogInfo(r *http.Request, format string, v ...interface{}) {
	logOut(loggerError, "info", r, format, v...)
}

// LogWarn logs a warning level error.
func LogWarn(r *http.Request, format string, v ...interface{}) {
	logOut(loggerError, "warn", r, format, v...)
}

// LogError logs an error level error.
func LogError(r *http.Request, format string, v ...interface{}) {
	logOut(loggerError, "error", r, format, v...)
}
