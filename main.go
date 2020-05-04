package main

import (
	"github.com/golang/geo/r2"
	"github.com/kr/pretty"
)

func main() {
	p1 := r2.Point{X: 100, Y: 100}
	p2 := r2.Point{X: 200, Y: 200}

	r := r2.RectFromPoints(p1, p2)

	ps := r.Vertices()

	pretty.Println(ps)

}
