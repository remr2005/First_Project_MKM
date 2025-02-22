package calculations

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func Nul_Acceleration(v *mat.VecDense, h, vM float64) *mat.VecDense {
	res := mat.NewVecDense(2, []float64{0, 0})
	v_norm := v.Norm(2)
	res.ScaleVec(-G_Scl()*v_norm/(vM*vM)*math.Exp(-h/1000), v)
	res.AddVec(res, G_Vec())
	return res
}

func Acceleration(v *mat.VecDense, vM float64) *mat.VecDense {
	res := mat.NewVecDense(2, []float64{0, 0})
	res.ScaleVec(-G_Scl()*v.Norm(2)/math.Pow(vM, 2), v)
	res.AddVec(res, G_Vec())
	return res
}
