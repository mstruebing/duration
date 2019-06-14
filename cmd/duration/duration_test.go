package main

import (
	"testing"
	"time"
)

func TestPadTimePart(t *testing.T) {
	padTimePartTests := []struct {
		input    int
		expected string
	}{
		{20, "20"},
		{2, "02"},
	}

	for _, tt := range padTimePartTests {
		actual := padTimePart(tt.input)

		if actual != tt.expected {
			t.Errorf("padTimePart(%d): expected: %v, got: %v", tt.input, tt.expected, actual)
		}
	}
}

func TestGetSeconds(t *testing.T) {
	getSecondsTest := []struct {
		input    time.Duration
		expected string
	}{
		{time.Nanosecond, "00"},
		{time.Second * 5, "05"},
		{time.Second * 65, "05"},
	}

	for _, tt := range getSecondsTest {
		actual := getSeconds(tt.input)

		if actual != tt.expected {
			t.Errorf("getSeconds(%d): expected: %v, got: %v", tt.input, tt.expected, actual)
		}
	}
}

func TestGetMinutes(t *testing.T) {
	getMinutesTest := []struct {
		input    time.Duration
		expected string
	}{
		{time.Nanosecond, "00"},
		{time.Minute * 5, "05"},
		{time.Minute * 65, "05"},
	}

	for _, tt := range getMinutesTest {
		actual := getMinutes(tt.input)

		if actual != tt.expected {
			t.Errorf("getMinutes(%d): expected: %v, got: %v", tt.input, tt.expected, actual)
		}
	}
}

func TestGetHours(t *testing.T) {
	getHoursTest := []struct {
		input    time.Duration
		expected string
	}{
		{time.Nanosecond, "00"},
		{time.Hour * 5, "05"},
		{time.Hour * 25, "25"},
	}

	for _, tt := range getHoursTest {
		actual := getHours(tt.input)

		if actual != tt.expected {
			t.Errorf("getHours(%d): expected: %v, got: %v", tt.input, tt.expected, actual)
		}
	}
}

func TestIsFlag(t *testing.T) {
	isMyFlag := isFlag("f", "foo")

	isFlagTest := []struct {
		input    string
		expected bool
	}{
		{"f", true},
		{"-f", true},
		{"--f", true},
		{"foo", true},
		{"-foo", true},
		{"--foo", true},
		{"fo", false},
		{"-fo", false},
		{"--fo", false},
		{"abc", false},
		{"--abc", false},
	}

	for _, tt := range isFlagTest {
		actual := isMyFlag(tt.input)

		if actual != tt.expected {
			t.Errorf("isMyFlag(%s): expected: %v, got: %v", tt.input, tt.expected, actual)
		}
	}
}

func TestIsHelpFlag(t *testing.T) {
	isHelpFlagTest := []struct {
		input    string
		expected bool
	}{
		{"h", true},
		{"help", true},
		{"-h", true},
		{"-help", true},
		{"--h", true},
		{"--help", true},
		{"f", false},
		{"foo", false},
		{"-f", false},
		{"-foo", false},
		{"--f", false},
		{"--foo", false},
	}

	for _, tt := range isHelpFlagTest {
		actual := isHelpFlag(tt.input)

		if actual != tt.expected {
			t.Errorf("isHelpFlag(%s): expected: %v, got: %v", tt.input, tt.expected, actual)
		}
	}
}

func TestIsVersionFlag(t *testing.T) {
	isVersionFlagTest := []struct {
		input    string
		expected bool
	}{
		{"v", true},
		{"version", true},
		{"-v", true},
		{"-version", true},
		{"--v", true},
		{"--version", true},
		{"f", false},
		{"foo", false},
		{"-f", false},
		{"-foo", false},
		{"--f", false},
		{"--foo", false},
	}

	for _, tt := range isVersionFlagTest {
		actual := isVersionFlag(tt.input)

		if actual != tt.expected {
			t.Errorf("isVersionFlag(%s): expected: %v, got: %v", tt.input, tt.expected, actual)
		}
	}
}
