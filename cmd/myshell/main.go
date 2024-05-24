package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Wait for user input
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		input = strings.TrimSpace(input)
		switch {
		case strings.HasPrefix(input, "exit"):
			exitCmd(input)
		case strings.HasPrefix(input, "echo"):
			echoCmd(input)
		default:
			cmdNotFound(input)
		}
	}
}

func cmdNotFound(cmd string) {
	fmt.Printf("%s: command not found\n", cmd)
}

func exitCmd(cmd string) {
	parser := strings.SplitN(cmd, " ", 2)

	exitCode, err := strconv.ParseInt(parser[1], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	os.Exit(int(exitCode))
}

func echoCmd(cmd string) {
	_, print, flag := strings.Cut(cmd, "echo")
	if flag {
		fmt.Println(print)
	}
}
