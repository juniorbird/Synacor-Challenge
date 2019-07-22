package main

import (
	"fmt"
	"os"
)

func main() {
	hasMsg := len(os.Args) > 1

	if !hasMsg {
		resultifier("please supply a message", true)
	}

	resultifier(os.Args[1], false)

}

func resultifier(msg string, crash bool) {
	fmt.Printf("the message is: %s\n", msg)
	if crash || len(msg) < 1 {
		os.Exit(1)
	}
	os.Exit(0)
}
