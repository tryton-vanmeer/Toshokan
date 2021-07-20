package utils

import "log"

func Error(msg interface{}) {
	log.Fatalf("\033[31mERROR\033[0m:) %s", msg)
}
