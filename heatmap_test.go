package heatmap

import "testing"

func TestAnalyse(t *testing.T) {
	data := [][]int{
		{1, 3, 5, 7},
		{0, 1, 2, 5},
		{5, 3, 2, 1},
	}

	w, h, max, min := analyse(data)

	if w != 4 && h != 3 && max != 7 && min != 0 {
		t.Fail()
	}

}
