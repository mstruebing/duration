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

const VERSION string = "1.0.0"

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
	// Prints every 250 milliseconds
	ticker := time.NewTicker(time.Millisecond * 250)

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

func printVersion() {
	fmt.Printf("duration v%s\n", VERSION)
}

func printHelp() {
	fmt.Println("USAGE: duration [h|help|v|version|<command>]")
	fmt.Println("where")
	fmt.Println("\t[v|version] - prints the version")
	fmt.Println("\t[h|help] - prints the help")
	fmt.Println("\t<command> - a SINGLE command or script to execute")
	fmt.Println("")
	fmt.Println("would work:")
	fmt.Println("\tduration sleep 5")
	fmt.Println("\tduration script.sh")
	fmt.Println("would NOT work:")
	fmt.Println("\tduration sleep 5 && sleep 4 (use ie a bash script instead)")
}

func isFlag(short string, long string) func(string) bool {
	return func(maybeFlag string) bool {
		trimmedFlag := strings.Trim(maybeFlag, "-")
		return trimmedFlag == short || trimmedFlag == long
	}
}

func isHelpFlag(input string) bool {
	return isFlag("h", "help")(input)
}

func isVersionFlag(input string) bool {
	return isFlag("v", "version")(input)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "ERROR: You need to provide a command to execute.")
		printHelp()
		os.Exit(1)
	}

	programOrFlag := strings.Join(os.Args[1:2], "")
	args := strings.Join(os.Args[2:], " ")

	if isVersionFlag(programOrFlag) {
		printVersion()
		os.Exit(0)
	}

	if isHelpFlag(programOrFlag) {
		printHelp()
		os.Exit(0)
	}

	cmd := exec.Command(programOrFlag, args)

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
