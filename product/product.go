package product

type Product struct {
	name        string
	description string
	category    string
	units       int
}

func (p *Product) AddUnits() {
	p.units++
}

func (p *Product) Units() int {
	return p.units
}
