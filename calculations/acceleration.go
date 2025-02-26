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
	vNorm := mat.Norm(v, 2) // Норма (модуль) вектора скорости

	// g_vect = (0, -g)
	gVect := mat.NewVecDense(2, []float64{0, -G_Scl()})

	// g(v * vNorm) / vM^2
	temp := mat.NewVecDense(2, nil)
	temp.ScaleVec(G_Scl()*vNorm/math.Pow(vM, 2), v)

	// a = gVect - temp
	a := mat.NewVecDense(2, nil)
	a.SubVec(gVect, temp)

	return a
}
