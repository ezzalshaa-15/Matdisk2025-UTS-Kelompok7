package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Running all SOAL solutions ===")
	fmt.Println("\n-- SOAL 1 --")
	solveSoal1()
	fmt.Println("\n-- SOAL 2 --")
	solveSoal2()
	fmt.Println("\n-- SOAL 3 --")
	solveSoal3()
	fmt.Println("\n-- SOAL 4 --")
	solveSoal4()
	fmt.Println("\n-- SOAL 5 --")
	solveSoal5()
	fmt.Println("\n-- SOAL 6 --")
	solveSoal6()
}
