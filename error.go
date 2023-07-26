package main

import (
	"fmt"
	"log"
)

func IfFatal(err error, message string) {
	if err != nil {
		fmt.Println(err)
		log.Fatalln(message)
	}
}
