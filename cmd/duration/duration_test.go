package main

import (
	"testing"
	"time"
)

func TestPadTimePart(t *testing.T) {
	result := padTimePart(20)

	if result != "20" {
		t.Errorf("Should not pad if value has two things, got: %v", result)
	}

	result = padTimePart(2)
	if result != "02" {
		t.Errorf("Should pad if value has one thing, got: %v", result)
	}
}

func TestGetSeconds(t *testing.T) {
	result := getSeconds(time.Nanosecond)

	if result != "00" {
		t.Errorf("Should correctly calculate seconds, got: %v", result)
	}

	result = getSeconds(time.Second * 5)

	if result != "05" {
		t.Errorf("Should correctly calculate seconds, got: %v", result)
	}

	result = getSeconds(time.Second * 65)

	if result != "05" {
		t.Errorf("Should correctly calculate seconds, got: %v", result)
	}
}

func TestGetMinutes(t *testing.T) {
	result := getMinutes(time.Nanosecond)

	if result != "00" {
		t.Errorf("Should correctly calculate seconds, got: %v", result)
	}

	result = getMinutes(time.Minute * 5)

	if result != "05" {
		t.Errorf("Should correctly calculate seconds, got: %v", result)
	}

	result = getMinutes(time.Minute * 65)

	if result != "05" {
		t.Errorf("Should correctly calculate seconds, got: %v", result)
	}
}

func TestGetHours(t *testing.T) {
	result := getHours(time.Nanosecond)

	if result != "00" {
		t.Errorf("Should correctly calculate seconds, got: %v", result)
	}

	result = getHours(time.Hour * 5)

	if result != "05" {
		t.Errorf("Should correctly calculate seconds, got: %v", result)
	}

	result = getHours(time.Hour * 25)

	if result != "25" {
		t.Errorf("Should correctly calculate seconds, got: %v", result)
	}
}
