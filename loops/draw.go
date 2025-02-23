package loops

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func Draw(points plotter.XYs) {
	p := plot.New()
	p.Title.Text = "График координаты пули"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	// Создаем линейный график
	line, err := plotter.NewLine(points)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	p.Add(line)
	// Сохраняем график в файл
	if err := p.Save(500*vg.Points(1), 500*vg.Points(1), "plot.png"); err != nil {
		fmt.Println("Ошибка сохранения:", err)
	}
}
