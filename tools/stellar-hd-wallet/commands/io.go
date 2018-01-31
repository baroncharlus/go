package commands

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var out io.Writer = os.Stdout

func readString() string {
	line, _ := reader.ReadString('\n')
	return strings.TrimRight(line, "\n")
}

func readUint() uint32 {
	line := readString()
	number, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal("Invalid value")
	}

	return uint32(number)
}

func printf(format string, a ...interface{}) {
	fmt.Fprintf(out, format, a...)
}

func println(a ...interface{}) {
	fmt.Fprintln(out, a...)
}

func readPassword(prompt string) (password string) {
	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintln(os.Stderr)
	}

	// gives our terminal back if we have a problem converting the file
	// descriptor to raw mode.
	defer terminal.Restore(int(os.Stdin.Fd()), oldState)

	// create a password entry terminal raw mode stdin
	t := terminal.NewTerminal(os.Stdin, "")
	pass, err := t.ReadPassword(prompt)
	if err != nil {
		fmt.Fprintln(os.Stderr)
		log.Fatal(err)
	}

	password = string(pass)
	return
}

func savePassword() (password string) {
	password = readPassword("Enter your password: ")
	confirm := readPassword("Confirm your password: ")
	if password != confirm {
		fmt.Println("Passwords did not match. Please try again:")
		savePassword()
	}
	return
}
