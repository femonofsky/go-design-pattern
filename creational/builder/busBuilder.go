package builder


type BusBuilder struct {
	v VehicleProduct
}

func (bus *BusBuilder) SetWheel () BuildProcess {
	bus.v.Wheel = 4 * 2
	return bus
}

func (bus *BusBuilder) SetSeats () BuildProcess {
	bus.v.Seats = 30
	return bus
}

func (bus *BusBuilder) SetStructure () BuildProcess {
	bus.v.Structure = "Bus"
	return bus
}

func (bus *BusBuilder) GetVehicle () VehicleProduct {
	return bus.v
}
