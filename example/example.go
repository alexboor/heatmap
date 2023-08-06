package main

import (
	"github.com/alexboor/heatmap"
	"image/color"
)

func main() {

	data := [][]int{
		{1, 3, 5, 7},
		{1, 1, 2, 5},
		{5, 3, 2, 1},
	}

	opts := heatmap.Options{
		Width:      500,
		Height:     350,
		Background: color.RGBA{0xff, 0xff, 0xff, 0xff},
	}

	hm := heatmap.New(opts)

	hm.Draw(data, "test.png")
}
