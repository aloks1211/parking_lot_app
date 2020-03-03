package models

type ParkingSlots struct {
	TotalParkedCars  int
	ParkingSlots     map[int]*Slot
	EmptyParkinglots []*Slot
	Size             int
}

type Slot struct {
	Id          int
	IsAvailable bool
	ParkedCar   *Car
}

func (parkingSlots *ParkingSlots) SetSize(size int) {
	parkingSlots.Size = size

}

func (parkingSlots *ParkingSlots) IncrementParkedCar() {
	parkingSlots.TotalParkedCars += 1

}

func (parkingSlots *ParkingSlots) DecrementParkedCar() {
	parkingSlots.TotalParkedCars -= 1

}

func (parkingSlots *ParkingSlots) GetSlotById(Id int) (*Slot, bool) {
	slot, ok := parkingSlots.ParkingSlots[Id]
	return slot, ok
}

func (parkingSlots *ParkingSlots) SetSlot(Id int, slot *Slot) {
	parkingSlots.ParkingSlots[Id] = slot
}

func (slot *Slot) SetisAvailable(isAvailable bool) {
	slot.IsAvailable = isAvailable

}
