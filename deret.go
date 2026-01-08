package main

import (
	"fmt"
	"math"
)

// SOAL 5: RELASI REKURENS ITERATIF
func solveSoal5() {
	C1, C2, N := 3, 5, 14
	a := make([]int, N+1)
	a[0], a[1] = 1, 1

	fmt.Printf("INPUT: C1=%d, C2=%d, N=%d\n", C1, C2, N)
	for i := 2; i <= N; i++ {
		a[i] = (C1 * a[i-1]) + (C2 * a[i-2])
	}
	fmt.Printf("HASIL AKHIR Suku ke-%d: %d\n", N, a[N])
}

// SOAL 6: ANALISIS KEDEKATAN DERET GEOMETRI
func solveSoal6() {
	a := 8.0
	r := 0.1
	N := 15

	sumBerhingga := a * (1 - math.Pow(r, float64(N))) / (1 - r)
	sumTakHingga := a / (1 - r)
	kedekatan := (sumBerhingga / sumTakHingga) * 100

	fmt.Printf("Input Paket: a=%.1f, r=%.1f, N=%d\n", a, r, N)
	fmt.Printf("Sum Berhingga S(%d): %.2f\n", N, sumBerhingga)
	fmt.Printf("Sum Tak Hingga S(inf): %.2f\n", sumTakHingga)
	fmt.Printf("Persentase Kedekatan: %.2f%%\n", kedekatan)
}
