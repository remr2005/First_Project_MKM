package loops

import (
	"main/calculations"

	"gonum.org/v1/gonum/mat"
)

func Iter(r, v, a *mat.VecDense, dt, vM float64) {
	r.CopyVec(calculations.Coordinate(r, v, a, dt))
	vPred := calculations.Velocity(v, a, dt)
	aNew := calculations.Nul_Acceleration(vPred, r.AtVec(1), vM)
	v.CopyVec(calculations.Velocity_Fix(v, a, aNew, dt))
	a.CopyVec(aNew)
}
