package heatmap

import (
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
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
// utilizing the function to get useful data in single interation
// TODO: add test
func analyse(d [][]int) (width int, height int, max int, min int) {
	for _, row := range d {
		width = 0
		for _, i := range row {
			fmt.Printf("%d", i)
			width = width + 1
			if i > max {
				max = i
			}
		}
		height = height + 1
		fmt.Printf("\n")
	}

	return width, height, max, min
}
