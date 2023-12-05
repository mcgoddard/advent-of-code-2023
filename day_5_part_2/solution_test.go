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
				end:   10,
			},
		},
	}
	ranges := []Range{
		{
			start: 1,
			end:   5,
		},
	}
	newRanges := splitRanges(list, ranges)
	expected := Range{
		start: 1,
		end:   5,
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
				end:   5,
			},
			Range{
				start: 6,
				end:   10,
			},
		},
	}
	ranges := []Range{
		{
			start: 1,
			end:   5,
		},
	}
	newRanges := splitRanges(list, ranges)
	expected := []Range{
		{
			start: 6,
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

func TestPartialOverlapStart(t *testing.T) {
	list := [][]Range{
		{
			Range{
				start: 1,
				end:   5,
			},
			Range{
				start: 10,
				end:   15,
			},
		},
	}
	ranges := []Range{
		{
			start: 4,
			end:   7,
		},
	}
	newRanges := splitRanges(list, ranges)
	expected := []Range{
		{
			start: 13,
			end:   14,
		},
		{
			start: 6,
			end:   7,
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
				end:   5,
			},
			Range{
				start: 11,
				end:   15,
			},
		},
	}
	ranges := []Range{
		{
			start: 4,
			end:   7,
		},
	}
	newRanges := splitRanges(list, ranges)
	expected := []Range{
		{
			start: 14,
			end:   15,
		},
		{
			start: 6,
			end:   7,
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
				end:   5,
			},
			Range{
				start: 6,
				end:   10,
			},
		},
	}
	ranges := []Range{
		{
			start: 2,
			end:   4,
		},
	}
	newRanges := splitRanges(list, ranges)
	expected := []Range{
		{
			start: 7,
			end:   9,
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
				end:   9,
			},
			Range{
				start: 12,
				end:   19,
			},
		},
	}
	ranges := []Range{
		{
			start: 1,
			end:   10,
		},
	}
	newRanges := splitRanges(list, ranges)
	expected := []Range{
		{
			start: 1,
			end:   1,
		},
		{
			start: 12,
			end:   19,
		},
		{
			start: 10,
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

func TestMultipleOverlaps(t *testing.T) {
	list := [][]Range{
		{
			Range{
				start: 2,
				end:   9,
			},
			Range{
				start: 22,
				end:   29,
			},
		},
		{
			Range{
				start: 12,
				end:   19,
			},
			Range{
				start: 32,
				end:   39,
			},
		},
	}
	ranges := []Range{
		{
			start: 1,
			end:   20,
		},
	}
	newRanges := splitRanges(list, ranges)
	expected := []Range{
		{
			start: 1,
			end:   1,
		},
		{
			start: 22,
			end:   29,
		},
		{
			start: 10,
			end:   11,
		},
		{
			start: 32,
			end:   39,
		}, {
			start: 20,
			end:   20,
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
