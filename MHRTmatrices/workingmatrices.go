package main

import "math"

func MultiplyMultipleMatrices(matrix [][4][4]float64) ([4][4]float64, [][4][4]float64) {
	var result [4][4]float64 = matrix[0]
	var course_decision = [][4][4]float64{}
	course_decision = append(course_decision, matrix[0])
	for i := 1; i < len(matrix); i++ {
		course_decision = append(course_decision, result)
		result = MultiplyMatrix(result, matrix[i])
		if i != len(matrix) {
			course_decision = append(course_decision, result)
		}
	}

	// multiplying matrices and storing result

	return result, course_decision
}

func MultiplyMatrix(matrixA [4][4]float64, matrixB [4][4]float64) [4][4]float64 {
	var total float64 = 0
	var result [4][4]float64

	// multiplying matrices and storing result
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				total = float64(total) + matrixA[i][k]*matrixB[k][j]
			}
			result[i][j] = total
			total = 0
		}
	}
	return result
}

func rotX(Q float64) [4][4]float64 {
	matrix := [4][4]float64{
		{1, 0, 0, 0},
		{0, math.Cos(Q), -math.Sin(Q), 0},
		{0, math.Sin(Q), math.Cos(Q), 0},
		{0, 0, 0, 1},
	}
	return matrix
}

func rotY(Q float64) [4][4]float64 {
	matrix := [4][4]float64{
		{math.Cos(Q), 0, math.Sin(Q), 0},
		{0, 1, 0, 0},
		{-math.Sin(Q), 0, math.Cos(Q), 0},
		{0, 0, 0, 1},
	}
	return matrix
}

func rotZ(Q float64) [4][4]float64 {
	matrix := [4][4]float64{
		{math.Cos(Q), -math.Sin(Q), 0, 0},
		{math.Sin(Q), math.Cos(Q), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	return matrix
}

func smX(Dx float64) [4][4]float64 {
	matrix := [4][4]float64{
		{1, 0, 0, Dx},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	return matrix
}

func smY(Dy float64) [4][4]float64 {
	matrix := [4][4]float64{
		{1, 0, 0, 0},
		{0, 1, 0, Dy},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	return matrix
}

func smZ(Dz float64) [4][4]float64 {
	matrix := [4][4]float64{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, Dz},
		{0, 0, 0, 1},
	}
	return matrix
}
