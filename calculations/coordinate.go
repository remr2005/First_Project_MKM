package calculations

import "gonum.org/v1/gonum/mat"

func Coordinate(r, v, a *mat.VecDense, t float64) *mat.VecDense {
	res := mat.NewVecDense(2, []float64{})
	v_ := mat.NewVecDense(2, []float64{})
	a_ := mat.NewVecDense(2, []float64{})
	v_.ScaleVec(t, v)
	a_.ScaleVec(t*t/2, a)
	res.AddVec(r, v_)
	res.AddVec(res, a_)
	return res
}
