package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := Point{1.0, 1.0}
	p2 := Point{1.0, 1.0}
	fmt.Println(Distance(p1, p2))

}

type Point struct {
	x float64
	y float64
}

func Distance(p1, p2 Point) float64 {
	dist := math.Sqrt(math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2))
	return dist
}
