package api

import (
	"github.com/golang/geo/r2"
)

func (p *TwoDimensionalPlane) AddDoubleIntersections() {

	for id1, rectangle1 := range p.Rectangles {
		for id2, rectangle2 := range p.Rectangles {
			if id1 == id2 {
				continue
			}

			k := DoubleIntersectedRectangleKey([2]int{id1, id2})
			if _, ok := p.DoubleIntersectedRectangle[k]; ok {
				continue
			}

			rect1 := rectangle1.convertToR2Rect()
			rect2 := rectangle2.convertToR2Rect()

			intersectedRect := rect1.Intersection(rect2)
			if intersectedRect == r2.EmptyRect() {
				continue
			}

			intersectedRectangle := createRectangle(intersectedRect)

			p.DoubleIntersectedRectangle[k] = *intersectedRectangle
		}
	}
}

func (p *TwoDimensionalPlane) AddTripleIntersections() {

	for id1, rectangle1 := range p.Rectangles {
		for id2, rectangle2 := range p.Rectangles {
			if id1 == id2 {
				continue
			}
			for id3, rectangle3 := range p.Rectangles {
				if id1 == id3 || id2 == id3 {
					continue
				}

				k := TripleIntersectedRectangleKey([3]int{id1, id2, id3})
				if _, ok := p.TripleIntersectedRectangle[k]; ok {
					continue
				}

				rect1 := rectangle1.convertToR2Rect()
				rect2 := rectangle2.convertToR2Rect()
				rect3 := rectangle3.convertToR2Rect()

				intersectedRect := rect1.Intersection(rect2).Intersection(rect3)
				if intersectedRect == r2.EmptyRect() {
					continue
				}

				intersectedRectangle := createRectangle(intersectedRect)

				p.TripleIntersectedRectangle[k] = *intersectedRectangle
			}
		}
	}
}
