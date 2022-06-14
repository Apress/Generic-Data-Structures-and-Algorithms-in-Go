
package main

import (
	"image/color"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type Point struct {
	X float64 
	Y float64
}

func definePoints(cities []Point, tour []int) plotter.XYs {
	pts := make(plotter.XYs, len(cities) + 1)
	pts[0].X = cities[0].X
	pts[0].Y = cities[0].Y
	for i := 1; i < len(cities); i++ {
		pts[i].X = cities[tour[i]].X
		pts[i].Y = cities[tour[i]].Y
	}
	pts[len(cities)].X = cities[0].X
	pts[len(cities)].Y = cities[0].Y
	return pts
}

func DrawTour(cities []Point, tour []int) {
	data := definePoints(cities, tour) // plotter.XYs
	p := plot.New()
	p.Title.Text = "TSP Tour"
	lines, points, err := plotter.NewLinePoints(data)
	if err != nil {
		panic(err)
	}
	lines.Color = color.RGBA{R: 255, A: 255}
	points.Shape = draw.PyramidGlyph{}
	points.Color = color.RGBA{B: 255, A: 255}
	p.Add(lines, points)
	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "tour.png"); err != nil {
		panic(err)
	}
}

func main() {
	numCities := 4
	cities := make([]Point, numCities)
	cities[0] = Point{0.0, 0.0}
	cities[1] = Point{3.0, 0.0}
	cities[2] = Point{3.0, 4.0}
	cities[3] = Point{1.0, 11.0}
	tour := []int{0, 3, 1, 2}
	DrawTour(cities, tour)
}


