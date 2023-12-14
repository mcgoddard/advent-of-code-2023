package main

import (
	"testing"
)

func TestConvertToFlagsStart(t *testing.T) {
	input := ".##"
	output := convertToFlags(input)
	expected := uint32(1)
	if output != expected {
		t.Fatalf("Incorrect: %q", output)
	}
}

func TestConvertToFlagsEnd(t *testing.T) {
	input := "##."
	output := convertToFlags(input)
	expected := uint32(4)
	if output != expected {
		t.Fatalf("Incorrect: %q", output)
	}
}

func TestConvertToFlagsMultiple(t *testing.T) {
	input := ".#."
	output := convertToFlags(input)
	expected := uint32(5)
	if output != expected {
		t.Fatalf("Incorrect: %q", output)
	}
}

func TestDifferenceZero(t *testing.T) {
	input1 := uint32(5)
	input2 := uint32(5)
	output := difference(input1, input2)
	expected := uint32(0)
	if output != expected {
		t.Fatalf("Incorrect: %q", output)
	}
}

func TestDifferenceOne(t *testing.T) {
	input1 := uint32(4)
	input2 := uint32(5)
	output := difference(input1, input2)
	expected := uint32(1)
	if output != expected {
		t.Fatalf("Incorrect: %q", output)
	}
}

func TestDifferenceTwo(t *testing.T) {
	input1 := uint32(7)
	input2 := uint32(5)
	output := difference(input1, input2)
	expected := uint32(2)
	if output != expected {
		t.Fatalf("Incorrect: %q", output)
	}
}

func TestPowerOfTwoTrue(t *testing.T) {
	inputs := []uint32{1, 2, 4, 8, 16}
	for _, input := range inputs {
		if !isPowerOfTwo(input) {
			t.Fatalf("Incorrect: %q", input)
		}
	}
}

func TestPowerOfTwoFalse(t *testing.T) {
	inputs := []uint32{0, 5, 7, 3}
	for _, input := range inputs {
		if isPowerOfTwo(input) {
			t.Fatalf("Incorrect: %q", input)
		}
	}
}
