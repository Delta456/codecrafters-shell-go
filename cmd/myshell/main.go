package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		input = strings.TrimSpace(input)
		switch {
		case strings.HasPrefix(input, "exit"):
			exitCmd(input)
		case strings.HasPrefix(input, "echo"):
			echoCmd(input)
		case strings.HasPrefix(input, "type"):
			typeCmd(input)
		case strings.HasPrefix(input, "pwd"):
			pwdCmd(input)
		case strings.HasPrefix(input, "cd"):
			cdCmd(input)
		default:
			execCmd(input)
		}
	}
}

func cmdNotFound(cmd string) {
	fmt.Fprintf(os.Stderr, "%s: not found\n", cmd)
}

func exitCmd(cmd string) {
	parser := strings.SplitN(cmd, " ", 2)
	if len(parser) == 1 {
		os.Exit(0)
	}

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

func typeCmd(cmd string) {
	parser := strings.SplitN(cmd, " ", 2)
	introspectCmd := strings.TrimSpace(parser[1])

	switch introspectCmd {
	case "echo", "exit", "type":
		fmt.Printf("%s is a shell builtin\n", introspectCmd)
	default:
		if path, ok := cmdinPath(introspectCmd); ok {
			fmt.Printf("%s is %s\n", introspectCmd, path)
		} else {
			cmdNotFound(introspectCmd)
		}
	}
}

func cmdinPath(cmd string) (string, bool) {
	env := os.Getenv("PATH")
	pathEnvVars := strings.Split(env, ":")

	for _, path := range pathEnvVars {
		if _, err := os.Stat(filepath.Join(path, cmd)); err == nil {
			return filepath.Join(path, cmd), true
		}
	}
	return "", false

}

func execCmd(cmd string) {
	cmds := strings.Split(cmd, " ")
	execCmd := exec.Command(cmds[0], cmds[1:]...)

	execCmd.Stderr = os.Stderr
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout

	err := execCmd.Run()
	if err != nil {
		cmdNotFound(cmds[0])
	}

}

func pwdCmd(_ string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid dir")
	}
	fmt.Println(dir)
}

func cdCmd(cmd string) {
	parser := strings.SplitN(cmd, " ", 2)
	introspectDir := strings.TrimSpace(parser[1])

	// Redirect to $HOME when dir is ~
	if introspectDir == "~" {
		introspectDir = os.Getenv("HOME")
	}

	err := os.Chdir(introspectDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", introspectDir)
	}
}
