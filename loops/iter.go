package loops

import (
	"main/calculations"
	"time"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot/plotter"
)

func Iter(r, v, a *mat.VecDense, dt, vM float64) {
	r.CopyVec(calculations.Coordinate(r, v, a, dt)) // Копируем новые координаты в r
	a0 := mat.NewVecDense(2, nil)
	a0.CopyVec(a) // Сохраняем текущее ускорение
	v.CopyVec(calculations.Velocity(v, a0, dt))
	a.CopyVec(calculations.Nul_Acceleration(v, r.AtVec(1), vM))
	v.CopyVec(calculations.Velocity_Fix(v, a0, a, dt))
}

func Loop(v0, vM, dt, alpha, x, y float64, ch chan plotter.XYs) {
	v := calculations.MakeVelocity(alpha, v0)
	a := mat.NewVecDense(2, nil) // nil → нулевые значения
	r := mat.NewVecDense(2, []float64{x, y})

	points := make(plotter.XYs, 0, 10000) // Уменьшаем начальный размер
	points = append(points, plotter.XY{X: r.AtVec(0), Y: r.AtVec(1)})

	for i := 1; ; i++ {
		Iter(r, v, a, dt, vM)
		if r.AtVec(1) < 0 {
			break
		}
		points = append(points, plotter.XY{X: r.AtVec(0), Y: r.AtVec(1)})
		ch <- points
		time.Sleep(300 * time.Millisecond)
	}
	close(ch)
}
