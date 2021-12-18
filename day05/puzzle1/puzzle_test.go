package main

import "testing"

func TestCalculateLinesOverlap(t *testing.T) {
	testCases := []struct {
		name  string
		vents [][4]int
		want  int
	}{
		{
			name: "Example",
			vents: [][4]int{
				{0, 9, 5, 9},
				{8, 0, 0, 8},
				{9, 4, 3, 4},
				{2, 2, 2, 1},
				{7, 0, 7, 4},
				{6, 4, 2, 0},
				{0, 9, 2, 9},
				{3, 4, 1, 4},
				{0, 0, 8, 8},
				{5, 5, 8, 2},
			},
			want: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CalculateLinesIntersection(tc.vents)

			if got != tc.want {
				t.Errorf("got %d lines overlaps; want %d", got, tc.want)
			}
		})
	}
}
