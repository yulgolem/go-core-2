package fib

import (
	"testing"
)

func Test_сalc(t *testing.T) {
	cases := []struct {
		Description string
		n           int
		Want        int
	}{
		{"Fibonacci`s number on index 0 is 0", 0, 0},
		{"Fibonacci`s number on index 1 is 1", 1, 1},
		{"Fibonacci`s number on index 2 is 1", 2, 1},
		{"Fibonacci`s number on index 3 is 2", 3, 2},
		{"Fibonacci`s number on index 19 is 4181", 19, 4181},
		{"Fibonacci`s number on index 20 is 6765", 20, 6765},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := count(test.n)
			if got != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	cases := []struct {
		Description string
		n           int
		Want        string
	}{
		{"Should respond with error on less than 0 inputs", -1, "input cannot be less than zero"},
		{"Should respond with error on more than 20 inputs", 23, "input cannot be more than 20"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			_, got := Count(test.n)
			if got.Error() != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}
