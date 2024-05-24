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
		if strings.HasPrefix(input, "exit") {
			exitCmd(input)
		}
		cmdNotFound(input)
	}
}

func cmdNotFound(cmd string) {
	// Trim spaces because the line endings are still CRLF.
	cmd = strings.TrimSpace(cmd)
	fmt.Printf("%s: command not found\n", cmd)
}

func exitCmd(cmd string) {
	// Trim spaces because the line endings are still CRLF.
	cmd = strings.TrimSpace(cmd)
	parser := strings.SplitN(cmd, " ", 2)

	exitCode, err := strconv.ParseInt(parser[1], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	os.Exit(int(exitCode))
}
