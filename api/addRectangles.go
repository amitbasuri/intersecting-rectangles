package api

type AddRectanglesRequest struct {
	Rectangles []Rectangle `json:"rects"`
}

func (p *TwoDimensionalPlane) AddRectangles(addRectRequest AddRectanglesRequest) {
	count := 1
	empty := Rectangle{}
	for _, r := range addRectRequest.Rectangles {
		if r == empty || r.W <= 0 || r.H <= 0 {
			continue
		}

		p.Rectangles[count] = r
		count++
	}
}
