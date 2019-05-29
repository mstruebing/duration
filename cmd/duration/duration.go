package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func printDuration() {
	start := time.Now()

	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		currentTime := time.Since(start)

		fmt.Printf("\rcommand running since: %.0fh - %.0fm - %.0fs", currentTime.Hours(), currentTime.Minutes(), currentTime.Seconds())
	}
}

func main() {
	program := strings.Join(os.Args[1:2], "")
	args := strings.Join(os.Args[2:], " ")

	go printDuration()
	output, err := exec.Command(program, args).Output()

	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%v", string(output[:]))
}
