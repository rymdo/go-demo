package main

import (
	"machine"
	"time"
)

func main() {
	go led()
	go print()
	for {
		time.Sleep(time.Millisecond * 1000)
	}
}

func led() {
	println("led: start")
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		println("led: low")
		led.Low()
		time.Sleep(time.Millisecond * 500)
		println("led: high")
		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}

func print() {
	println("print: start")
	for {
		println("print: Hello There!")
		time.Sleep(time.Millisecond * 1000)
	}
}
