package main

import (
	"fmt"
	"math/rand"
)

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// SOAL 1: OPERASI HIMPUNAN KOMPLEKS
func solveSoal1() {
	N := 135
	limit := N

	gen := func(n int) []int {
		s := []int{}
		for len(s) < n {
			val := rand.Intn(limit) + 1
			if !contains(s, val) {
				s = append(s, val)
			}
		}
		return s
	}

	A, B, C := gen(5), gen(5), gen(5)
	fmt.Printf("Generating Sets... (N=%d)\n", N)
	fmt.Printf("A: %v | B: %v | C: %v\n", A, B, C)

	var xor []int
	for _, v := range A {
		if !contains(B, v) {
			xor = append(xor, v)
		}
	}
	for _, v := range B {
		if !contains(A, v) {
			xor = append(xor, v)
		}
	}

	var diff []int
	for _, v := range xor {
		if !contains(C, v) {
			diff = append(diff, v)
		}
	}

	var inter []int
	for _, v := range A {
		if contains(C, v) {
			inter = append(inter, v)
		}
	}

	R := diff
	for _, v := range inter {
		if !contains(R, v) {
			R = append(R, v)
		}
	}

	fmt.Printf("Hasil Operasi R: %v\n", R)
}

// SOAL 2: ANALISIS PASANGAN SUBSET
func solveSoal2() {
	M := 9
	K := 26

	S := make([]int, M)
	for i := range S {
		S[i] = rand.Intn(50) + 1
	}

	fmt.Printf("Set S: %v | Target K: %d\n", S, K)
	fmt.Println("Subset 2-Elemen (Sum > K):")

	count := 0
	for i := 0; i < len(S); i++ {
		for j := i + 1; j < len(S); j++ {
			sum := S[i] + S[j]
			if sum > K {
				count++
				fmt.Printf("%d. {%d, %d} (Sum=%d)\n", count, S[i], S[j], sum)
			}
		}
	}
	fmt.Printf("Total: %d Pasangan\n", count)
}
