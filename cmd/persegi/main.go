package main

import (
	"fmt"

	"github.com/partaikoding/go-bangun-datar/pkg/bd"
)

func main() {

	// diketahui sebuah persegi dengan panjang sisi 4cm
	persegi := bd.Persegi{Sisi: 4}

	// ditanya: berapa luas bidang persegi tersebut ?
	fmt.Printf(
		`Luas persegi dengan panjang sisi %dcm adalah: %d cm persegi`,
		persegi.Sisi, persegi.Luas(),
	)
	fmt.Println()
	// ditanya: berapa keliling persegi tersebut ?
	fmt.Printf(
		`Keliling persegi dengan panjang sisi %dcm adalah: %d cm`,
		persegi.Sisi, persegi.Keliling(),
	)
	fmt.Println()

}
