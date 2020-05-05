package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoDimensionalPlane_AddDoubleIntersections(t *testing.T) {
	twoDimPlane := NewTwoDimPlane()
	request := AddRectanglesRequest{Rectangles: []Rectangle{
		{
			X: 100,
			Y: 100,
			W: 50,
			H: 50,
		},
		{
			X: 70,
			Y: 90,
			W: 50,
			H: 30,
		},
	}}

	twoDimPlane.AddRectangles(request)
	twoDimPlane.AddDoubleIntersections()

	expected := Rectangle{
		X: 100,
		Y: 90,
		W: 20,
		H: 30,
	}

	assert.Equal(t, expected, twoDimPlane.DoubleIntersectedRectangle["1_2"], "Must be equal")

}

func TestTwoDimensionalPlane_AddTripleIntersections(t *testing.T) {
	twoDimPlane := NewTwoDimPlane()
	request := AddRectanglesRequest{Rectangles: []Rectangle{
		{
			X: 100,
			Y: 100,
			W: 50,
			H: 50,
		},
		{
			X: 75,
			Y: 75,
			W: 50,
			H: 50,
		},
		{
			X: 60,
			Y: 60,
			W: 50,
			H: 50,
		},
	}}

	twoDimPlane.AddRectangles(request)
	twoDimPlane.AddDoubleIntersections()
	twoDimPlane.AddTripleIntersections()

	expected := Rectangle{
		X: 100,
		Y: 60,
		W: 10,
		H: 10,
	}

	assert.Equal(t, expected, twoDimPlane.TripleIntersectedRectangle["1_2_3"], "Must be equal")

}
