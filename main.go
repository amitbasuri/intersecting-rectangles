package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/amitbasuri/interesting-rectangles/api"
	"log"
	"os"
	"strings"
)

func main() {
	inputFilePath := flag.String("file", "sample_input.json", "input json file")
	if inputFilePath == nil {
		log.Fatal("input json file missing")
	}
	flag.Parse()

	inputFile, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal("unable to open input file: ", inputFilePath)
	}
	defer inputFile.Close()

	jsonParser := json.NewDecoder(inputFile)
	request := api.AddRectanglesRequest{}
	err = jsonParser.Decode(&request)
	if err != nil {
		log.Fatal("unable to decode input file: ", err.Error())
	}

	// create 2D space
	twoDimPlane := api.NewTwoDimPlane()
	if len(request.Rectangles) > 10 {
		request.Rectangles = request.Rectangles[:10]
	}

	// add rectangles to the 2D space
	twoDimPlane.AddRectangles(request)

	n := len(twoDimPlane.Rectangles)
	fmt.Println("\nInput:")
	for k := 1; k <= n; k++ {
		v := twoDimPlane.Rectangles[k]
		fmt.Printf("     %d: Rectangle at (%d, %0.d), w=%d ,h=%d.\n", k, int(v.X), int(v.Y), v.W, v.H)
	}

	// calculate intersections
	twoDimPlane.AddDoubleIntersections()
	twoDimPlane.AddTripleIntersections()

	// print intersections
	fmt.Println("Intersections:")

	i := 1
	for k, v := range twoDimPlane.DoubleIntersectedRectangle {
		arr := strings.Split(k, "_")
		fmt.Printf("     %d: Between rectangle %s and %s at (%d,%d), w=%d, h=%d.\n", i, arr[0], arr[1], int(v.X), int(v.Y), v.W, v.H)
		i++
	}

	for k, v := range twoDimPlane.TripleIntersectedRectangle {
		arr := strings.Split(k, "_")
		fmt.Printf("     %d: Between rectangle %s, %s and %s at (%d,%d), w=%d, h=%d.\n", i, arr[0], arr[1], arr[2], int(v.X), int(v.Y), v.W, v.H)
		i++
	}
}
