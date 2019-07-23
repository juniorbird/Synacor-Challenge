package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	msg, err := parseArgs(os.Args)
	resultifier(msg, err)
}

func parseArgs(args []string) (msg string, err error) {
	hasMsg := len(args) > 1

	if !hasMsg {
		msg = "please supply a message"
		err = errors.New("no message")
	} else {
		msg = args[1]
		err = nil
	}

	return msg, err
}

func resultifier(msg string, err error) {
	fmt.Printf("the message is: %s\n", msg)
	if err != nil || len(msg) < 1 {
		os.Exit(1)
	}
	os.Exit(0)
}
