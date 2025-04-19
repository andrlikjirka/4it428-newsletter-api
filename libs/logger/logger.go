package logger

import "log"

func Info(a ...any) {
	log.Println("[INFO]", a)
}

func Warning(a ...any) {
	log.Println("[WARNING]", a)
}

func Error(a ...any) {
	log.Println("[ERROR]", a)
}

func Debug(a ...any) {
	log.Println("[DEBUG]", a)
}
