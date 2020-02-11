package builder


type BuildProcess interface {
	SetWheel() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheel()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess){
	f.builder = b
}

type VehicleProduct struct {
	Wheel int
	Seats int
	Structure string
}



