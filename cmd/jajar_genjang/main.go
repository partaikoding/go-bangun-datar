package main

import (
	"fmt"

	"github.com/partaikoding/go-bangun-datar/pkg/bd"
)

func main() {

	// diketahui sebuah Jajar Genjang dengan panjang
	// sisi panjang 8cm, sisi lebar 10cm, tinggi 5cm
	jajarGenjang := bd.JajarGenjang{
		SisiPanjang: 8,  // sisi datar
		SisiLebar:   12, // sisi miring
		Tinggi:      5,
	}

	// ditanya: berapa luas bidang Jajar Genjang tersebut ?
	fmt.Printf(
		`Luas Jajar Genjang dengan 
		sisi panjang %dcm & tinggi %dcm 
		adalah: %d cm persegi`,
		jajarGenjang.SisiPanjang,
		jajarGenjang.Tinggi,
		jajarGenjang.Luas(),
	)
	fmt.Println()
	// ditanya: berapa keliling Jajar Genjang tersebut ?
	fmt.Printf(
		`Keliling Jajar Genjang dengan 
		sisi panjang %dcm & sisi lebar %dcm
		adalah: %d cm`,
		jajarGenjang.SisiPanjang,
		jajarGenjang.SisiLebar,
		jajarGenjang.Keliling(),
	)
	fmt.Println()

}
