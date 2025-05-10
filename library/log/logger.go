package log

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "test", log.Ldate|log.Ltime|log.Lshortfile)
