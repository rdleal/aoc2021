package main

import "testing"

func TestCalculateSubmarinePosition(t *testing.T) {
	testCases := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name: "Example",
			input: []string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2",
			},
			want: 150,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CalculateSubmarinePosition(tc.input)

			if got != tc.want {
				t.Errorf("got multiplied position: %d; want %d", got, tc.want)
			}
		})
	}
}
