package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

// https://zetcode.com/golang/exec-command/
func runCommandWithParams(cmd string, cmd_args ...string) {
	args := strings.Join(cmd_args, " ")

	command := exec.Command(cmd, args)

	var out bytes.Buffer
	command.Stdout = &out

	err := command.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func runCommand(cmd string) {
	command := exec.Command("cmd")

	var out bytes.Buffer
	command.Stdout = &out

	err := command.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func clearConsole() {
	runCommand("cls") // Clear the console
}

func displayFrame(frame string) {
	clearConsole()
	fmt.Println(frame)
}

func animate(frames []string, delay time.Duration) {
	for _, frame := range frames {
		displayFrame(frame)
		time.Sleep(delay)
	}
}

func main() {
	frames := []string{
		"Frame 1",
		"Frame 2",
		"Frame 3",
		"Frame 4",
	}

	animate(frames, 500*time.Millisecond)
}
