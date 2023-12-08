package main

import (
	"fmt"
	"testing"
)

func TestNoOverlap(t *testing.T) {
	list := [][]Range{
		{
			Range{
				start: 6,
				end:   11,
			},
			Range{
				start: 6,
				end:   11,
			},
		},
	}
	ranges := []Range{
		{
			start: 1,
			end:   6,
		},
	}
	newRanges := splitRanges(list, ranges)
	expected := Range{
		start: 1,
		end:   6,
	}
	if newRanges[0] != expected {
		t.Fatalf("Incorrect newRanges: %q", newRanges)
	}
}

func TestExactOverlap(t *testing.T) {
	list := [][]Range{
		{
			Range{
				start: 1,
				end:   6,
			},
			Range{
				start: 6,
				end:   11,
			},
		},
	}
	ranges := []Range{
		{
			start: 1,
			end:   6,
		},
	}
	newRanges := splitRanges(list, ranges)
	expected := []Range{
		{
			start: 6,
			end:   11,
		},
	}
	fmt.Println("NewRanges: ", newRanges)
	for i, newRange := range newRanges {
		if newRange != expected[i] {
			t.Fatalf("Incorrect newRanges: %v", newRanges)
		}
	}
	if len(newRanges) != len(expected) {
		t.Fatalf("Incorrect newRanges: %v", newRanges)
	}
}

func TestPartialOverlapStart(t *testing.T) {
	list := [][]Range{
		{
			Range{
				start: 1,
				end:   6,
			},
			Range{
				start: 10,
				end:   16,
			},
		},
	}
	ranges := []Range{
		{
			start: 4,
			end:   8,
		},
	}
	newRanges := splitRanges(list, ranges)
	expected := []Range{
		{
			start: 13,
			end:   15,
		},
		{
			start: 6,
			end:   8,
		},
	}
	fmt.Println("NewRanges: ", newRanges)
	for i, newRange := range newRanges {
		if newRange != expected[i] {
			t.Fatalf("Incorrect newRanges: %v", newRanges)
		}
	}
	if len(newRanges) != len(expected) {
		t.Fatalf("Incorrect newRanges: %v", newRanges)
	}
}

func TestPartialOverlapEnd(t *testing.T) {
	list := [][]Range{
		{
			Range{
				start: 1,
				end:   6,
			},
			Range{
				start: 11,
				end:   16,
			},
		},
	}
	ranges := []Range{
		{
			start: 4,
			end:   8,
		},
	}
	newRanges := splitRanges(list, ranges)
	expected := []Range{
		{
			start: 14,
			end:   16,
		},
		{
			start: 6,
			end:   8,
		},
	}
	fmt.Println("NewRanges: ", newRanges)
	for i, newRange := range newRanges {
		if newRange != expected[i] {
			t.Fatalf("Incorrect newRanges: %v", newRanges)
		}
	}
	if len(newRanges) != len(expected) {
		t.Fatalf("Incorrect newRanges: %v", newRanges)
	}
}

func TestExactSubrangeOverlap(t *testing.T) {
	list := [][]Range{
		{
			Range{
				start: 1,
				end:   6,
			},
			Range{
				start: 6,
				end:   11,
			},
		},
	}
	ranges := []Range{
		{
			start: 2,
			end:   5,
		},
	}
	newRanges := splitRanges(list, ranges)
	expected := []Range{
		{
			start: 7,
			end:   10,
		},
	}
	fmt.Println("NewRanges: ", newRanges)
	for i, newRange := range newRanges {
		if newRange != expected[i] {
			t.Fatalf("Incorrect newRanges: %v", newRanges)
		}
	}
	if len(newRanges) != len(expected) {
		t.Fatalf("Incorrect newRanges: %v", newRanges)
	}
}

func TestExactSuperRangeOverlap(t *testing.T) {
	list := [][]Range{
		{
			Range{
				start: 2,
				end:   10,
			},
			Range{
				start: 12,
				end:   20,
			},
		},
	}
	ranges := []Range{
		{
			start: 1,
			end:   11,
		},
	}
	newRanges := splitRanges(list, ranges)
	expected := []Range{
		{
			start: 12,
			end:   20,
		},
		{
			start: 1,
			end:   2,
		},
		{
			start: 10,
			end:   11,
		},
	}
	fmt.Println("NewRanges: ", newRanges)
	for i, newRange := range newRanges {
		if newRange != expected[i] {
			t.Fatalf("Incorrect newRanges: %v", newRanges)
		}
	}
	if len(newRanges) != len(expected) {
		t.Fatalf("Incorrect newRanges: %v", newRanges)
	}
}

func TestMultipleOverlaps(t *testing.T) {
	list := [][]Range{
		{
			Range{
				start: 2,
				end:   10,
			},
			Range{
				start: 22,
				end:   30,
			},
		},
		{
			Range{
				start: 12,
				end:   20,
			},
			Range{
				start: 32,
				end:   40,
			},
		},
	}
	ranges := []Range{
		{
			start: 1,
			end:   21,
		},
	}
	newRanges := splitRanges(list, ranges)
	expected := []Range{
		{
			start: 22,
			end:   30,
		},
		{
			start: 32,
			end:   40,
		},
		{
			start: 1,
			end:   2,
		},
		{
			start: 10,
			end:   12,
		},
		{
			start: 20,
			end:   21,
		},
	}
	fmt.Println("NewRanges: ", newRanges)
	for i, newRange := range newRanges {
		if newRange != expected[i] {
			t.Fatalf("Incorrect newRanges: %v", newRanges)
		}
	}
	if len(newRanges) != len(expected) {
		t.Fatalf("Incorrect newRanges: %v", newRanges)
	}
}
