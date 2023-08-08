package heatmap

import (
	"fmt"
	"image/color"
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

func TestCellColor(t *testing.T) {
	tests := []struct {
		Name string
		Data []int // min, max, current
		Want color.RGBA
	}{
		{
			"Minimal 0-10 0",
			[]int{0, 10, 0},
			color.RGBA{G: 255, A: 255},
		},
		{
			"Middle 0-10 5",
			[]int{0, 10, 5},
			color.RGBA{G: 255, R: 255, A: 255},
		},
		{
			"Maximal 0-10 10",
			[]int{0, 10, 10},
			color.RGBA{R: 255, A: 255},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			res := cellColor(tt.Data[0], tt.Data[1], tt.Data[2])
			if res != tt.Want {
				t.Errorf("want %v, given %v", tt.Want, res)
			}
		})
	}

}
