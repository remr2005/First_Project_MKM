package calculations

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func Velocity(v, a *mat.VecDense, t float64) *mat.VecDense {
	res := mat.NewVecDense(2, []float64{0, 0})
	res.ScaleVec(t, a)
	res.AddVec(res, v)
	return res
}

func Velocity_Fix(v, a0, a1 *mat.VecDense, t float64) *mat.VecDense {
	res := mat.NewVecDense(2, []float64{0, 0})
	res.AddVec(a1, a0)
	res.ScaleVec(t/2, res)
	res.AddVec(res, v)
	return res
}

func MakeVelocity(alpha, v float64) *mat.VecDense {
	return mat.NewVecDense(2, []float64{v * math.Cos(alpha*math.Pi/180), v * math.Sin(alpha*math.Pi/180)})
}
