package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cmd, input, err := parseArgs(os.Args)
	resultifier(cmd, input, err)
}

// parseArgs separates out the command from the input
// It dumps you out if you don't provide a command or if the command provided isn't valid
func parseArgs(args []string) (cmd uint16, input []string, err error) {
	hasArgs := len(args) > 1

	if !hasArgs {
		err = errors.New("no command supplied")
		return 0, nil, err
	}

	input = strings.Split(args[1], ",")
	var cmd64 uint64
	cmd64, err = strconv.ParseUint(input[0], 10, 16)
	cmd = uint16(cmd64)
	input = input[1:]

	return cmd, input, err
}

// func routeCommand(cmd uint16) (result string, err error) {
// 	result = fmt.Printf("You told me to do command %d\n", cmd)
// 	return result, err
// }

func resultifier(cmd uint16, input []string, err error) {
	fmt.Printf("the command is: %d\n", cmd)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
