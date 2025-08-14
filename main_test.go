package main

// Пишите тесты в этом файле
import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		validate func([]int) bool
	}{
		{
			"Positive size",
			100,
			func(s []int) bool {
				if len(s) != 100 {
					return false
				}
				for _, v := range s {
					if v < 0 || v >= 100*10 {
						return false
					}
				}
				return true
			},
		},
		{
			"Zero size",
			0,
			func(s []int) bool { return s == nil },
		},
		{
			"Negative size",
			-1,
			func(s []int) bool { return s == nil },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if !tt.validate(got) {
				t.Errorf("generateRandomElements(%d) = %v, validation failed", tt.size, got)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Normal case", []int{1, 5, 3, 9, 2}, 9},
		{"Single element", []int{42}, 42},
		{"All equal", []int{7, 7, 7, 7}, 7},
		{"With negatives", []int{-1, -5, -3}, -1},
		{"Empty slice", []int{}, 0},
		{"Max first", []int{10, 3, 5}, 10},
		{"Max last", []int{1, 4, 20}, 20},
		{"All zeros", []int{0, 0, 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.input); got != tt.expected {
				t.Errorf("maximum(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
