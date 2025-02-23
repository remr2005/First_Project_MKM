package main

import (
	"fmt"
	"main/calculations"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Создаем новый график
	p := plot.New()
	p.Title.Text = "График y = x^2"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	var vM float64 = 150
	v := mat.NewVecDense(2, []float64{525, 525})
	// a0 := calculations.G_Vec()
	// a0 := calculations.Acceleration(v, vM)
	a0 := calculations.Nul_Acceleration(v, 0, vM)
	a := calculations.Acceleration(v, vM)
	a.CopyVec(a0)
	fmt.Println(a0)
	r := mat.NewVecDense(2, []float64{0, 0})
	// Данные для графика
	dt := 0.1
	points := make(plotter.XYs, 1000000000)
	for i := range points[1:] {
		r = calculations.Coordinate(r, v, a, dt)
		points[i].X = r.AtVec(0)
		v = calculations.Velocity(v, a0, dt)
		a := calculations.Nul_Acceleration(v, r.AtVec(1), vM)
		v = calculations.Velocity_Fix(v, a0, a, dt)
		a0.CopyVec(a)
		points[i].Y = r.AtVec(1)
		if r.AtVec(1) < 0 {
			points = points[:i+1]
			break
		}
		// Выводим значения ускорения и скорости на каждой итерации
		fmt.Printf("Итерация %d:\n", i+1)
		fmt.Println(r)
		fmt.Printf("Скорость (v): %v\n", v)
		fmt.Printf("Ускорение (a): %v\n\n", a)
	}

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
