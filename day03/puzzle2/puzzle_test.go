package main

import "testing"

func TestCalculateLifeSupportRating(t *testing.T) {
	testCases := []struct {
		name  string
		input []string
		want  uint64
	}{
		{
			name: "Example",
			input: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			want: 230,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CalculateCalculateLifeSupportRating(tc.input)

			if got != tc.want {
				t.Errorf("got life support rating: %d; want %d", got, tc.want)
			}
		})
	}
}
