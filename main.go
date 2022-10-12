package utils

import (
	"log"

	"github.com/TwiN/go-color"
)

func HasError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Info(msg string) {
	println(color.Ize(color.Green, msg))
}
