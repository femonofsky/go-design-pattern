package builder



type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheel () BuildProcess {
	b.v.Wheel = 2
	return b
}

func (b *BikeBuilder) SetSeats () BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure () BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

