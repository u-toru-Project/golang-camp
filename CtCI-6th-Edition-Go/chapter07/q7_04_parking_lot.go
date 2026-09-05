package chapter07

import "fmt"

type Vehicle struct {
	Model  string
	Size   int
	Number string
	Parked bool
}

func newVehicle(model string, size int, number string) *Vehicle {
	if model == "" {
		panic("model must be a non-empty string")
	}
	if size < 1 {
		panic("size must be a positive int")
	}
	if number == "" {
		panic("number must be a non-empty string")
	}
	return &Vehicle{Model: model, Size: size, Number: number}
}

func (v *Vehicle) IsParked() bool { return v.Parked }

type Bike struct{ *Vehicle }

func NewBike(model string, size int, number string) *Bike {
	return &Bike{Vehicle: newVehicle(model, size, number)}
}

type Scooter struct{ *Vehicle }

func NewScooter(model string, size int, number string) *Scooter {
	return &Scooter{Vehicle: newVehicle(model, size, number)}
}

type Car struct{ *Vehicle }

func NewCar(model string, size int, number string) *Car {
	return &Car{Vehicle: newVehicle(model, size, number)}
}

type Bus struct{ *Vehicle }

func NewBus(model string, size int, number string) *Bus {
	return &Bus{Vehicle: newVehicle(model, size, number)}
}

type ParkZone struct {
	Capacity       int
	SpaceAvailable int
	Parked         map[int]*Vehicle
	nextToken      int
}

func NewParkZone(capacity ...int) *ParkZone {
	cap := 10
	if len(capacity) > 0 {
		cap = capacity[0]
	}
	if cap < 1 {
		panic("capacity must be positive")
	}
	return &ParkZone{Capacity: cap, SpaceAvailable: cap, Parked: map[int]*Vehicle{}, nextToken: 1000}
}

func (z *ParkZone) Park(vehicle *Vehicle) *int {
	if vehicle == nil {
		panic("vehicle is nil")
	}
	if vehicle.Parked {
		panic("vehicle is already parked")
	}
	if !z.IsSpaceAvailable(vehicle.Size) {
		return nil
	}
	token := z.Register(vehicle)
	z.SpaceAvailable -= vehicle.Size
	vehicle.Parked = true
	return &token
}

func (z *ParkZone) IsSpaceAvailable(size int) bool { return z.SpaceAvailable-size >= 0 }

func (z *ParkZone) Register(vehicle *Vehicle) int {
	token := z.nextToken
	z.nextToken++
	z.Parked[token] = vehicle
	return token
}

func (z *ParkZone) Depark(token int) *Vehicle {
	parked, ok := z.Parked[token]
	if !ok {
		panic("Invalid token or vehicle not found")
	}
	delete(z.Parked, token)
	parked.Parked = false
	z.SpaceAvailable += parked.Size
	return parked
}

func (z *ParkZone) ListParkedVehicles() []*Vehicle {
	out := make([]*Vehicle, 0, len(z.Parked))
	for _, v := range z.Parked {
		out = append(out, v)
	}
	return out
}

func RunQ704() {
	zone := NewParkZone()
	bike := NewBike("Suzuki Access", 1, "MH14AB1234")
	token := zone.Park(bike.Vehicle)
	fmt.Printf("Parked token %d, remaining %d\n", *token, zone.SpaceAvailable)
	zone.Depark(*token)
}
