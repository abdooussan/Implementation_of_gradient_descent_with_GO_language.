package RL
import (
	"math"
	"testing"
)
func Test_Regression_with_empty_arrays(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Regression with empty arrays did not panic.")
		}}()
	RegressionLR([][]float64{}, []float64{}, 200, 0.5)
	RegressionLG([]float64{}, [][]float64{},100000,0.00001)
}
func Test_Regression_with_non_empty_arrays(t *testing.T) {
	x_values := [][]float64{{1, 0}, {2, 0}, {3, 0}}
	y_values := []float64{1, 3, 5}
	x_values2 := [][]float64{{1, 0}, {2, 0}, {2.6, 0}}
	y_values2 := []float64{2, 4, 5}
	a_current := RegressionLR(x_values, y_values, 10000, 0.005)
	a_current2 := RegressionLG(y_values2, x_values2, 1000000,0.00001)
	assertDelta(t, 17.4, a_current2[0], 0.1)
	assertDelta(t, 2, a_current[0], 0.01)
	assertDelta(t, 17.4, a_current2[0], 0.1)
}
func assertDelta(t *testing.T, x float64, y float64, delta float64) {
	abs_difference := math.Abs(float64(x - y))
	if abs_difference > float64(delta) {
		t.Error("Difference between x and y is greater than delta!")
	}}
