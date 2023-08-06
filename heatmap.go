package heatmap

import (
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"math/bits"
)

type Heatmap struct {
	Title          string
	BackgroudColor color.RGBA
	Width          int
	Height         int
}

type Options struct {
	Width      int
	Height     int
	Background color.RGBA
}

func New(opt Options) *Heatmap {
	hm := &Heatmap{
		Width:          opt.Width,
		Height:         opt.Height,
		BackgroudColor: opt.Background,
	}

	return hm
}

func (hm Heatmap) Draw(data [][]int) error {

	dest := image.NewRGBA(image.Rect(0, 0, hm.Width, hm.Height))
	gc := draw2dimg.NewGraphicContext(dest)

	gc.SetFillColor(hm.BackgroudColor)
	gc.SetStrokeColor(color.RGBA{A: 0xff})
	gc.SetLineWidth(1)

	hLen, vLen, max, min := analyse(data)

	fmt.Printf("%d x %d; max: %d, min: %d\n", hLen, vLen, max, min)

	draw2dimg.SaveToPngFile("test.png", dest)
	return nil
}

// analyse returns width, height of the given matrix and max and min values
// utilizing the function to get useful data in single interation O(x*y)
// TODO: add test
func analyse(d [][]int) (width int, height int, max int, min int) {
	max = (1 << bits.UintSize) / -2 // set minimal int value
	min = (1<<bits.UintSize)/2 - 1  // set maximal int value

	for _, row := range d {
		x := 0
		for _, i := range row {
			x = x + 1
			if i > max {
				max = i
			}
			if i < min {
				min = i
			}
		}

		if x >= width {
			width = x
		} else {
			min = 0 // if horizontal row has fewer items, that means min value is 0 TODO: proper handle negative values
		}

		height = height + 1
	}

	// TODO: proper handle negative values #2 - this is not uint
	if width == 0 {
		max = 0
		min = 0
	}

	return width, height, max, min
}
