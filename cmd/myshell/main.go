package main

import (
	"bufio"
	"fmt"
	"os"
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

		cmdNotFound(input)
	}
}

func cmdNotFound(cmd string) {
	// Trim spaces because the line endings are still CRLF.
	cmd = strings.TrimSpace(cmd)
	fmt.Printf("%s: command not found\n", cmd)
}
