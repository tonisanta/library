package product

import "time"

type Product struct {
	name        string
	description string
	category    string
	units       int
	createdAt   time.Time
	updatedAt   time.Time
}

func (p *Product) AddUnits(units int) {
	p.units += units
}

func (p *Product) Units() int {
	return p.units
}
