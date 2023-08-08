package heatmap

import (
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"math"
	"math/bits"
)

type Heatmap struct {
	Title          string
	BackgroudColor color.RGBA
	Width          int
	Height         int
	PaddingX       int
	PaddingY       int
	GCtx           *draw2dimg.GraphicContext
	Dest           *image.RGBA
}

type Options struct {
	Width      int
	Height     int
	Background color.RGBA
}

// New create new heatmap instance
// TODO: polish it
func New(opt Options) *Heatmap {
	dest := image.NewRGBA(image.Rect(0, 0, opt.Width, opt.Height))
	gc := draw2dimg.NewGraphicContext(dest)

	hm := &Heatmap{
		Width:          opt.Width,
		Height:         opt.Height,
		PaddingX:       10,
		PaddingY:       10,
		BackgroudColor: opt.Background,
		GCtx:           gc,
		Dest:           dest,
	}

	//gc.SetFillColor(hm.BackgroudColor)
	//gc.SetStrokeColor(color.RGBA{A: 0xff})

	return hm
}

func (hm Heatmap) Draw(data [][]int, path string) error {

	hLen, vLen, max, min := analyse(data)

	fmt.Printf("%d x %d; max: %d, min: %d\n", hLen, vLen, max, min)

	var (
		startX  float64 = 10
		startY  float64 = 10
		marginX float64 = 2
		marginY float64 = 2

		cellWidth  = cellWidth(float64(hm.Width), float64(hm.PaddingX), marginX, hLen)
		cellHeight = cellHeight(float64(hm.Height), float64(hm.PaddingY), marginY, vLen)
	)

	for _, row := range data {
		x := startX
		for _, i := range row {
			hm.drawCell(x, startY, cellWidth, cellHeight, cellColor(min, max, i))
			x = x + cellWidth + marginX
		}
		startY = startY + cellHeight + marginY
	}

	draw2dimg.SaveToPngFile(path, hm.Dest)
	return nil
}

// drawCell draw cell on the canvas started from given top-left corner coordinate
func (hm Heatmap) drawCell(x float64, y float64, w float64, h float64, c color.RGBA) {
	hm.GCtx.SetFillColor(c)
	hm.GCtx.SetStrokeColor(c)
	hm.GCtx.SetLineWidth(1)
	hm.GCtx.BeginPath()
	hm.GCtx.LineTo(x, y)
	hm.GCtx.LineTo(x+w, y)
	hm.GCtx.LineTo(x+w, y+h)
	hm.GCtx.LineTo(x, y+h)
	hm.GCtx.LineTo(x, y)
	hm.GCtx.FillStroke()
	return
}

// getCellWidth return width of cell calculated from given
func cellWidth(CanvasWidth float64, CanvasPaddingX float64, CellMarginX float64, n int) float64 {
	return (CanvasWidth - (CanvasPaddingX * 2) - ((float64(n) - 1) * CellMarginX)) / float64(n)
}

// getCellHeight return heigh of the cell calculated based on given parameters
func cellHeight(CanvasHeight float64, CanvasPaddingX float64, CellMarginY float64, n int) float64 {
	return (CanvasHeight - (CanvasPaddingX * 2) - ((float64(n) - 1) * CellMarginY)) / float64(n)
}

// cellColor return calculated color based on the min, max and current values
//
//	The default color palette is established using the green - yellow - red color scheme, with green
//	representing the minimum value and red representing the maximum value. The returned value will be
//	positioned within the range defined by these two extremums.
func cellColor(min int, max int, v int) color.RGBA {
	// min value is green = rgba(0, 255, 0, 0)
	// middle is yellow = rgba(255, 255, 0, 0) +255 red chan
	// max is red = rgba(255, 0, 0, 0) - 255 green chan
	// Than means here we have +255 and then -255 steps for the whole range

	rangeMax := 255 + 255
	pShift := (100 * v) / int(math.Abs(float64(max))+math.Abs(float64(min)))
	val := pShift * rangeMax / 100

	var R uint8
	var G uint8

	switch {
	case val <= 255:
		R = uint8(val)
		G = 255
	case val > 255:
		R = 255
		G = 255 - (uint8(val) - 255)
	}

	return color.RGBA{R: R, G: G, B: 0, A: 255}
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
