package log

import "log"

func Debug(v ...interface{}) {
	log.Println(v...)
}

func Info(v ...interface{}) {
	log.Println(v...)
}

func Warning(v ...interface{}) {
	log.Println(v...)
}

func Error(v ...interface{}) {
	log.Println(v...)
}

func Fatal(v ...interface{}) {
	log.Println(v...)
}
