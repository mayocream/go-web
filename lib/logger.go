package lib

import (
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	// Log as JSON instead of the default ASCII formatter.
	Logger.SetFormatter(&Logger.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	// Logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	Logger.SetLevel(Logger.WarnLevel)
}