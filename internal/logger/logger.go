package logger

import (
	"log"
	"os"
)

// InitLogger инициализирует логирование.
func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
