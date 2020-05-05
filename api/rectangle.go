package api

import "github.com/golang/geo/r2"

type Rectangle struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W int     `json:"w"`
	H int     `json:"h"`
}

func (r Rectangle) convertToR2Rect() r2.Rect {
	p1 := r2.Point{
		X: r.X,
		Y: r.Y,
	}

	p2 := r2.Point{
		X: r.X + float64(r.W),
		Y: r.Y - float64(r.H),
	}

	return r2.RectFromPoints(p1, p2)

}

func createRectangle(rect r2.Rect) *Rectangle {
	r := &Rectangle{}
	r.X = rect.Lo().X
	r.Y = rect.Hi().Y
	r.W = int(rect.Hi().X - rect.Lo().X)
	r.H = int(rect.Hi().Y - rect.Lo().Y)

	return r
}
