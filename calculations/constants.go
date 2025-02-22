package calculations

import "gonum.org/v1/gonum/mat"

func G_Vec() *mat.VecDense {
	return mat.NewVecDense(2, []float64{0, -9.80665})
}

func G_Scl() float64 {
	return 9.80665
}
