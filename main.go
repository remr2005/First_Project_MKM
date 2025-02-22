package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// Создаем вектор размерности 3 с заданными значениями
	v := mat.NewVecDense(3, []float64{1.0, 2.0, 3.0})

	// Вывод вектора
	fmt.Printf("Vector:\n%v\n", mat.Formatted(v))
}
