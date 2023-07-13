package main

import "log"

func IfFatal(err error, message string) {
	if err != nil {
		log.Fatalln(message)
	}
}
