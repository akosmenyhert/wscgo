// +build windows

package main

//import "fmt"

// LOW = 0
const LOW = 0

// HIGH = 1
const HIGH = 1

// INPUT = 0
const INPUT = 0

// OUTPUT = 1
const OUTPUT = 1

func mcp23017Setup(expansionBase int, address int) {
}

func digitalWrite(pin int, value int) {
	//fmt.Printf("DO[%d]:%d\n", pin, value)
}

func digitalRead(pin int) int {
	return LOW
}

func pinMode(pin int, mode int) {

}
