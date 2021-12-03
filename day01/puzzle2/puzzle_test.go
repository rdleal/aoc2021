package main

import "testing"

func TestMeasureDepthSlidingWindowIncreases(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name: "Example",
			input: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			want: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := MeasureDepthSlidingWindowIncreases(tc.input)

			if got != tc.want {
				t.Errorf("got sliding window increases: %d; want %d", got, tc.want)
			}
		})
	}
}
