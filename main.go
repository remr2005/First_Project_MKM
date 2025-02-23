package main

import (
	"main/loops"

	"gonum.org/v1/plot/plotter"
)

func main() {
	ch := make(chan plotter.XYs, 1)
	go loops.Loop(750, 150, 0.2, 45, 0, 0, ch)
	for points := range ch {
		loops.Draw(points)
	}
}
