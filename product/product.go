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

func (p *Product) AddUnit() {
	p.units++
}

func (p *Product) AddUnits(unitsToAdd int) {
	p.units += unitsToAdd
}

func (p *Product) Units() int {
	return p.units
}
