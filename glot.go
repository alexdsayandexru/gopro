package main

import (
	"github.com/Arafatk/glot"
)

func goGlot() {
	points := [][]float64{{1, 2, 3, 4, 5}, {1, 4, 9, 16, 25}}
	plot, _ := glot.NewPlot(2, true, false)
	plot.SetTitle("Hello world")
	plot.SetXLabel("X")
	plot.SetYLabel("Y")
	plot.AddPointGroup("Points", "circle", points)
	plot.SetFormat("png")
	/*err := plot.SavePlot("data.png")
	if err != nil {
		fmt.Println("error1:", err)
	}*/
	plot.Close()

	/*dimensions := 3
	persist := false
	debug := false
	plot2, _ := glot.NewPlot(dimensions, persist, debug)
	plot2.AddPointGroup("Sample 1", "lines", []float64{2, 3, 4, 1})
	plot2.SetTitle("Test Results")
	plot2.SetZrange(-2, 2)
	err = plot2.SavePlot("1.jpeg")
	if err != nil {
		fmt.Println("error2:", err)
	}

	dimensions = 2
	persist = false
	debug = false
	plot3, _ := glot.NewPlot(dimensions, persist, debug)
	fct := func(x float64) float64 { return (math.Exp(x)) }
	groupName := "Exponential Curve"
	style := "lines"
	pointsX := []float64{1, 2, 3, 4, 5}
	plot3.AddFunc2d(groupName, style, pointsX, fct)
	err = plot3.SavePlot("2.png")
	plot3.Close()
	if err != nil {
		fmt.Println("error3:", err)
	}*/
}
