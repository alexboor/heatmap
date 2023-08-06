package heatmap

import (
	"fmt"
	"testing"
)

func TestAnalyse(t *testing.T) {
	tests := []struct {
		Data              [][]int
		WantHorizontalLen int
		WantVerticalLen   int
		WantMax           int
		WantMin           int
	}{
		// case 1
		{
			Data: [][]int{
				{1, 3, 5, 7},
				{1, 1, 2, 5},
				{5, 3, 2, 1},
			},
			WantHorizontalLen: 4,
			WantVerticalLen:   3,
			WantMax:           7,
			WantMin:           1,
		},

		// case 2
		{
			Data: [][]int{
				{1, 2, 3},
				{0, 1, 2},
				{50, 3, 2},
			},
			WantHorizontalLen: 3,
			WantVerticalLen:   3,
			WantMax:           50,
			WantMin:           0,
		},

		// case 3
		{
			Data: [][]int{
				{},
				{},
			},
			WantHorizontalLen: 0,
			WantVerticalLen:   2,
			WantMax:           0,
			WantMin:           0,
		},

		// case 4
		{
			Data: [][]int{
				{1, 2, 3},
				{1},
			},
			WantHorizontalLen: 3,
			WantVerticalLen:   2,
			WantMax:           3,
			WantMin:           0,
		},

		// case 5
		{
			Data: [][]int{
				{1, 2},
				{},
			},
			WantHorizontalLen: 2,
			WantVerticalLen:   2,
			WantMax:           2,
			WantMin:           0,
		},

		// case 5
		{
			Data: [][]int{
				{1, 2},
			},
			WantHorizontalLen: 2,
			WantVerticalLen:   1,
			WantMax:           2,
			WantMin:           1,
		},
	}

	for k, tt := range tests {
		name := fmt.Sprintf("case %d", k+1)
		t.Run(name, func(t *testing.T) {
			h, v, max, min := analyse(tt.Data)

			if h != tt.WantHorizontalLen {
				t.Errorf("Horizontal Len error: got %d, want %d", h, tt.WantHorizontalLen)
			}
			if v != tt.WantVerticalLen {
				t.Errorf("Vertical Len error: got %d, want %d", h, tt.WantVerticalLen)
			}
			if max != tt.WantMax {
				t.Errorf("Max value error: got %d, want %d", max, tt.WantMax)
			}
			if min != tt.WantMin {
				t.Errorf("Min value error: got %d, want %d", min, tt.WantMin)
			}

		})
	}

}
