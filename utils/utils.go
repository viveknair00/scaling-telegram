package utils

import (
	"log"
)

func Subtract(x, y int) (res int) {
	return x - y
}

func Add(x, y int) (res int) {
	return x + y
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
