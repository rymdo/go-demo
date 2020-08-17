package main

import (
	"time"
)

func main() {
	for {
		println("Hello tiny world")
		time.Sleep(time.Millisecond * 250)
	}
}
