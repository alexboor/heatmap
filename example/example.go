package main

import (
	"github.com/alexboor/heatmap"
	"image/color"
)

func main() {

	data := [][]int{
		{1, 3, 5, 7},
		{0, 1, 2, 5},
		{5, 3, 2, 1},
	}

	opts := heatmap.Options{
		Width:      350,
		Height:     200,
		Background: color.RGBA{0xff, 0xff, 0xff, 0xff},
	}

	hm := heatmap.New(opts)

	hm.Draw(data)
}
