package RL

import (
	"math"
)

func RegressionLR(x_values [][]float64, y_values []float64, epochs int, learning_rate float64) []float64 {
	if len(x_values) == 0 || len(y_values) == 0 {
		panic("Input arrays must not be empty.")
	}
	dimension := len(x_values[0])
	a_current := make([]float64, dimension+1)
	for i := 0; i < epochs; i++ {
		a_current = update(a_current, x_values, y_values, learning_rate)
	}
	return a_current
}




func update(a_currentt []float64, x_values [][]float64, y_values []float64, learning_rate float64) []float64 {
	a_current := a_currentt[0 : len(a_currentt)-1]
	b := a_currentt[len(a_currentt)-1]
	var length = len(y_values)
	var dimension = len(x_values[0])
	var s float64
	for l := 0; l < dimension; l++ {
		for i := 0; i < length; i++ {
			s = 0
			for j := 0; j < dimension; j++ {
				s += x_values[i][j] * a_current[j]
			}
			var two_over_n = float64(2) / float64(length)
			a_current[l] += learning_rate * two_over_n * x_values[i][l] * (y_values[i] - s - b)
		}
	}
	for i := 0; i < length; i++ {
		s = 0
		for j := 0; j < dimension; j++ {
			s += x_values[i][j] * a_current[j]
		}
		var two_over_n = float64(2) / float64(length)
		b += learning_rate * (two_over_n * (y_values[i] - s - b))
	}

	return append(a_current, b)
}



func Seg(X []float64, B1 []float64, B0 float64) float64 {
	sum := B0
	for i := 0; i < len(X); i++ {
		sum += X[i] * B1[i]
	}
	a := 1 / (1 + math.Exp(-sum))
	return a
}




func J(B1 []float64, B0 float64, Y []float64, X [][]float64) float64 {
	m := len(Y)
	sum := 0.
	for i := 0; i < m; i++ {
		h := Seg(X[i], B1, B0)
		if h <= 0. {
			h = 0.0000001
		}
		if h >= 1. {
			h = 0.9999999
		}
		sum += Y[i]*math.Log(h) + (1-Y[i])*math.Log(1-h)
	}
	resultat := -sum / float64(m)
	return resultat
}
func RegressionLG(Y []float64, X [][]float64, ite int, alp float64) []float64 {
	
	B0 := 0.
	B1 := make([]float64, len(X[0]))
	sum := 0.
	for i := 0; i < ite; i++ {
		sum = 0.
		for k := 0; k < len(Y); k++ {
			sum += Seg(X[k], B1, B0) - Y[k]
		}
		B0 = B0 - alp*sum
		for j := 0; j < len(X[0]); j++ {
			sum = 0.
			for k := 0; k < len(Y); k++ {
				sum += (Seg(X[k], B1, B0) - Y[k]) * X[k][j]
			}

			B1[j] = B1[j] - alp*sum
		}
	}
	return append(B1, B0)
}
