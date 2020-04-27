package binarysearch

import "testing"

var tests = []struct {
	input []int
	val int
	want int
}{
	{
		input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		val: 5,
		want:4,
	},
	{
		input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		val: 0,
		want:-1,
	},
}

func TestBinarySearch(t *testing.T) {
	for _, test := range tests {
		if got := BinarySearch(test.input, test.val); got != test.want {
			t.Errorf("BinarySearch(%v, %v) = %v", test.input, test.val, test.want)
		}
	}
}

func TestRecursionBinarySearch(t *testing.T) {
	for _, test := range tests {
		if got := RecursionBinarySearch(test.input, test.val); got != test.want {
			t.Errorf("BinarySearch(%v, %v) = %v", test.input, test.val, test.want)
		}
	}
}