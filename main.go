package main

import (
	"main/loops"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"gonum.org/v1/plot/plotter"
)

func main() {
	ch := make(chan plotter.XYs, 1)
	go loops.Loop(750, 150, 0.2, 45, 0, 0, ch)
	myApp := app.New()
	w := myApp.NewWindow("Image")
	w.Resize(fyne.NewSize(500, 500))
	image := canvas.NewImageFromFile("plot.png")
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)
	go func() {
		for points := range ch {
			loops.Draw(points)
			image.Refresh()
		}
	}()
	w.ShowAndRun()
}
