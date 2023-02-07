package bd

type JajarGenjang struct {
	SisiPanjang int
	SisiLebar   int
	Tinggi      int
}

func (p *JajarGenjang) Keliling() int {
	return 2 * (p.SisiPanjang + p.SisiLebar)
}

func (p *JajarGenjang) Luas() int {
	return p.SisiPanjang * p.Tinggi
}
