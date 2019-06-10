package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	tm "github.com/buger/goterm"
)

// Pads an integer with zeroes to the left
// i.e 01
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

func clearTime() {
	tm.MoveCursor(tm.Width()-30, tm.Height())
	tm.Printf("                               ")
}

func printDuration(start time.Time) {
	currentTime := time.Since(start)

	tm.MoveCursor(tm.Width()-30, tm.Height())
	tm.Printf("command running since: %s:%s:%s", getHours(currentTime), getMinutes(currentTime), getSeconds(currentTime))
	tm.MoveCursor(0, tm.Height())
	tm.Flush()
}

func printDurationAndOutput(output *bytes.Buffer) {
	ticker := time.NewTicker(time.Nanosecond)

	outputAccumulator := ""
	outputPrint := ""

	start := time.Now()
	printDuration(start)

	for range ticker.C {
		currentOutput := output.String()

		if strings.Compare(currentOutput, outputAccumulator) == 1 {
			outputPrint = strings.Replace(currentOutput, outputAccumulator, "", 1)
			outputAccumulator = currentOutput

			clearTime()
			tm.MoveCursor(0, tm.Height())
			tm.Printf("%s", string(outputPrint))
			tm.Flush()
		}

		printDuration(start)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "ERROR: You need to provide a command to execute.")
		os.Exit(1)
	}

	program := strings.Join(os.Args[1:2], "")
	args := strings.Join(os.Args[2:], " ")

	cmd := exec.Command(program, args)

	var output bytes.Buffer

	cmd.Stdout = &output
	cmd.Stderr = &output

	go printDurationAndOutput(&output)

	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(os.Stderr, "\nERROR: Please try to put your command into a script and execute that.\n")
		os.Exit(1)
	}

	time.Sleep(time.Second)
}
