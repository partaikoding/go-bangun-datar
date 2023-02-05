package bd

type Persegi struct {
	Sisi int
}

func (p *Persegi) Keliling() int {
	return p.Sisi * 4
}

func (p *Persegi) Luas() int {
	return p.Sisi * p.Sisi
}
