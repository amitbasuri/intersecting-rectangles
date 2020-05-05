package api

type AddRectanglesRequest struct {
	Rectangles []Rectangle `json:"rects"`
}

func (p *TwoDimensionalPlane) AddRectangles(addRectRequest AddRectanglesRequest) {
	count := 1
	for _, r := range addRectRequest.Rectangles {
		p.Rectangles[count] = r
		count++
	}
}
