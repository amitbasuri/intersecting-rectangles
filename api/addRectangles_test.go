package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoDimensionalPlane_AddRectangles(t *testing.T) {

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
			W: 20,
			H: 20,
		},
	}}

	twoDimPlane.AddRectangles(request)

	assert.Len(t, twoDimPlane.Rectangles, 2, "Length must be 2")
	for i, r := range request.Rectangles {
		assert.Equal(t, r, twoDimPlane.Rectangles[i+1], "Must be equal")
	}

}
