package product

import "time"

type Product struct {
	name        string
	description string
	category    string
	units       int
	createdAt   time.Time
}

func (p *Product) AddUnits() {
	p.units++
}

func (p *Product) Units() int {
	return p.units
}
