package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func padTimePart(timePart int) string {
	return fmt.Sprintf("%02d", timePart)
}

func getSeconds(time time.Duration) string {
	seconds := int(time.Seconds()) % 60
	return padTimePart(seconds)
}

func getMinutes(time time.Duration) string {
	minutes := int(time.Minutes()) % 60
	return padTimePart(minutes)
}

func getHours(time time.Duration) string {
	hours := int(time.Hours())
	return padTimePart(hours)
}

func printDuration() {
	fmt.Printf("command running since: %s:%s:%s", getHours(0), getMinutes(0), getSeconds(0))
	start := time.Now()

	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		currentTime := time.Since(start)
		fmt.Printf("\rcommand running since: %s:%s:%s", getHours(currentTime), getMinutes(currentTime), getSeconds(currentTime))
	}
}

func printCmdOutput(output string) {
	fmt.Printf("\n\nOutput:\n%s", output)
}

func main() {
	program := strings.Join(os.Args[1:2], "")
	args := strings.Join(os.Args[2:], " ")

	go printDuration()
	cmd := exec.Command(program, args)

	// Currently the output is printed at the end of the program
	// We can not differntiate between stdout and stderr anymore
	// I couldn't find a good solution to print realtime while also
	// printing the current duration readable until now.
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	printCmdOutput(string(output))
}
