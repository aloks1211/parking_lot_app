package utils

import (
	"fmt"
	"github.com/parking_lot/src/models"
	"sort"
	"strconv"
)

func CreateParkingLotFunc(N int) (*models.ParkingSlots, error) {
	parkingSlots := models.ParkingSlots{}
	if N <= 0 || N > 6 {
		return nil, &InvalidParkingLotError
	}

	parkingSlots.SetSize(N)
	parkingSlots.ParkingSlots = make(map[int]*models.Slot)

	for i := 1; i <= parkingSlots.Size; {
		newSlot := models.Slot{
			Id:          i,
			IsAvailable: true,
			ParkedCar:   nil,
		}
		parkingSlots.EmptyParkinglots = append(parkingSlots.EmptyParkinglots, &newSlot)
		parkingSlots.SetSlot(i, &newSlot)
		i += 1

	}
	fmt.Printf("Created a parking lot with slots %s \n", strconv.Itoa(len(parkingSlots.ParkingSlots)))
	return &parkingSlots, nil
}

func Status(parkingSlot models.ParkingSlots) error {
	if parkingSlot.TotalParkedCars == 0 {
		return &ParkingLotEmptyError
	}

	var slotIds []int
	for k := range parkingSlot.ParkingSlots {
		slotIds = append(slotIds, k)
	}
	sort.Ints(slotIds)

	fmt.Printf("Slot No.\tRegistration_No\tColour\n")
	for _, slotId := range slotIds {
		slot, ok := parkingSlot.GetSlotById(slotId)
		if !ok {
			return &NoSlotByIDFound
		}
		if !slot.IsAvailable && slot.ParkedCar != nil {
			fmt.Printf("%s\t%s\t%s\n", strconv.Itoa(slot.Id), slot.ParkedCar.RegNumber, slot.ParkedCar.Color)
		}
	}

	return nil

}

func Park(parkingSlot *models.ParkingSlots, car models.Car) error {
	if parkingSlot == nil {
		return &ParkingLotNotDefined
	}
	if parkingSlot.TotalParkedCars == parkingSlot.Size {
		return &ParkingFullError
	}

	slot := GetNearestEmptyParkingSLot(*parkingSlot)
	slot.SetisAvailable(false)
	car.SlotId = slot.Id
	slot.ParkedCar = &car
	parkingSlot.IncrementParkedCar()
	parkingSlot.EmptyParkinglots = RemoveLot(*parkingSlot)

	fmt.Printf("Allocated slot number: %s\n", strconv.Itoa(slot.Id))

	return nil

}

func Leave(parkingSlot *models.ParkingSlots, slotId int) error {
	if parkingSlot == nil {
		return &ParkingLotNotDefined
	}

	if parkingSlot.TotalParkedCars == 0 {
		return &ParkingLotEmptyError
	}

	slot, ok := parkingSlot.GetSlotById(slotId)
	if !ok {
		return &NoSlotByIDFound
	}

	if slot.IsAvailable {
		return &SlotEmptyError
	}

	if slot.Id == slotId {
		slot.SetisAvailable(true)
		slot.ParkedCar = nil
		parkingSlot.EmptyParkinglots = append(parkingSlot.EmptyParkinglots, slot)
		fmt.Printf("Slot number %s is free\n", strconv.Itoa(slot.Id))

	}
	parkingSlot.DecrementParkedCar()

	return nil

}

func GetNearestEmptyParkingSLot(parkingSlots models.ParkingSlots) *models.Slot {
	if len(parkingSlots.EmptyParkinglots) == 1 {
		return parkingSlots.EmptyParkinglots[0]
	}
	sort.SliceStable(parkingSlots.EmptyParkinglots, func(i, j int) bool {
		return parkingSlots.EmptyParkinglots[i].Id < parkingSlots.EmptyParkinglots[j].Id
	})

	return parkingSlots.EmptyParkinglots[0]
}

func RemoveLot(parkingSlots models.ParkingSlots) []*models.Slot {
	return parkingSlots.EmptyParkinglots[1:]

}

func GetSlotsByColor(parkingSlot models.ParkingSlots, color string) ([]int, error) {
	var slots []int
	for _, slot := range parkingSlot.ParkingSlots {
		if !slot.IsAvailable && slot.ParkedCar != nil && slot.ParkedCar.Color == color {
			slots = append(slots, slot.Id)
		}

	}
	if len(slots) == 0 {
		return nil, &NoSlotsByColorFound
	}
	return slots, nil
}

func GetSlotByRegNumber(parkingSlot models.ParkingSlots, regNum string) (int, error) {
	var slotId int
	for _, slot := range parkingSlot.ParkingSlots {
		if !slot.IsAvailable && slot.ParkedCar != nil && slot.ParkedCar.RegNumber == regNum {
			slotId = slot.Id
			break
		}

	}
	if slotId == 0 {
		return 0, &NoSlotByRegistrationNumberFound
	}
	return slotId, nil
}

func GetRegistrationNumbersByColor(parkingSlot models.ParkingSlots, color string) ([]string, error) {
	var regNums []string
	for _, slot := range parkingSlot.ParkingSlots {
		if !slot.IsAvailable && slot.ParkedCar != nil && slot.ParkedCar.Color == color {
			regNums = append(regNums, slot.ParkedCar.RegNumber)
		}

	}
	if len(regNums) == 0 {
		return nil, &NoRegistrationByColorFound
	}
	return regNums, nil
}
