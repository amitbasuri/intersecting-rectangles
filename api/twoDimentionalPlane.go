package api

import (
	"fmt"
	"sort"
)

type RectangleSet map[*Rectangle]struct{}

type TwoDimensionalPlane struct {
	Rectangles                 map[int]Rectangle
	DoubleIntersectedRectangle map[string]Rectangle
	TripleIntersectedRectangle map[string]Rectangle
}

func NewTwoDimPlane() *TwoDimensionalPlane {
	return &TwoDimensionalPlane{
		Rectangles:                 make(map[int]Rectangle),
		DoubleIntersectedRectangle: make(map[string]Rectangle),
		TripleIntersectedRectangle: make(map[string]Rectangle),
	}
}

func DoubleIntersectedRectangleKey(a [2]int) string {
	sa := sort.IntSlice([]int{a[0], a[1]})
	sa.Sort()
	return fmt.Sprintf("%d_%d", sa[0], sa[1])
}

func TripleIntersectedRectangleKey(a [3]int) string {
	sa := sort.IntSlice([]int{a[0], a[1], a[2]})
	sa.Sort()
	return fmt.Sprintf("%d_%d_%d", sa[0], sa[1], sa[2])
}
