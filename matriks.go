package main

import (
	"fmt"
	"math/rand"
)

// SOAL 3: PERKALIAN DAN TRACE MATRIKS
func solveSoal3() {
	ordo := 2
	genMat := func(n int) [][]int {
		m := make([][]int, n)
		for i := range m {
			m[i] = make([]int, n)
			for j := 0; j < n; j++ {
				m[i][j] = rand.Intn(10) + 1
			}
		}
		return m
	}

	A := genMat(ordo)
	B := genMat(ordo)

	R := make([][]int, ordo)
	trace := 0
	for i := 0; i < ordo; i++ {
		R[i] = make([]int, ordo)
		for j := 0; j < ordo; j++ {
			for k := 0; k < ordo; k++ {
				R[i][j] += A[i][k] * B[k][j]
			}
		}
		trace += R[i][i]
	}

	fmt.Printf("Matrix A: %v\nMatrix B: %v\n", A, B)
	fmt.Printf("Hasil Matriks R: %v\nTrace: %d\n", R, trace)
}

// SOAL 4: TRANSFORMASI BARIS
func solveSoal4() {
	ordo := 4
	matrix := make([][]int, ordo)
	for i := range matrix {
		matrix[i] = make([]int, ordo)
		for j := 0; j < ordo; j++ {
			matrix[i][j] = rand.Intn(50) + 1
		}
	}

	fmt.Printf("Matrix M (Generated): %v\n", matrix)

	matrix[0], matrix[ordo-1] = matrix[ordo-1], matrix[0]
	fmt.Printf("Matrix M Terkini: %v\n", matrix)

	maxVal, px, py := matrix[0][0], 0, 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] > maxVal {
				maxVal, px, py = matrix[i][j], i, j
			}
		}
	}
	fmt.Printf("Nilai Maksimum: %d ditemukan di Posisi (%d,%d)\n", maxVal, px, py)
}
