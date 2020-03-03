package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestCreateParkingLot(t *testing.T) {
	N := 6
	parkingLots, err := CreateParkingLotFunc(N)
	assert.Nil(t, err)
	assert.Equal(t, len(parkingLots.ParkingSlots), N)
	assert.Equal(t, parkingLots.TotalParkedCars, 0)

}

func TestCreateParkingLot_isFull(t *testing.T) {
	N := 6
	parkingLots, err := CreateParkingLotFunc(N)
	assert.Nil(t, err)
	assert.Equal(t, len(parkingLots.ParkingSlots), N)
	assert.Equal(t, parkingLots.Size, N)
}

func TestCreateZeroParkingLot(t *testing.T) {
	N := 0
	parkingLots, err := CreateParkingLotFunc(N)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), InvalidParkingLotError.UserMessage)
	assert.Nil(t, parkingLots)
}

func TestCreateNegativeParkingLot(t *testing.T) {
	N := -1
	parkingLots, err := CreateParkingLotFunc(N)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), InvalidParkingLotError.UserMessage)
	assert.Nil(t, parkingLots)
}

func TestParkCar(t *testing.T) {
	N := 1
	parkingLots, _ := CreateParkingLotFunc(N)
	mockCar := GetMockCar("KA-01-HH-1234", "White")
	err := Park(parkingLots, mockCar)
	assert.Nil(t, err)
	assert.Equal(t, parkingLots.TotalParkedCars, 1)
	assert.Equal(t, len(parkingLots.EmptyParkinglots), 0)
	slot, ok := parkingLots.GetSlotById(1)
	assert.True(t, ok)
	assert.False(t, slot.IsAvailable)
	assert.Equal(t, slot.ParkedCar.SlotId, 1)
}

func TestParkCar_NilParkingLot(t *testing.T) {
	mockCar := GetMockCar("KA-01-HH-1234", "White")
	err := Park(nil, mockCar)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), ParkingLotNotDefined.UserMessage)
}

func TestParkCar_ParkingLotisFull(t *testing.T) {
	N := 1
	parkingLots, _ := CreateParkingLotFunc(N)
	mockCar := GetMockCar("KA-01-HH-1234", "White")
	err := Park(parkingLots, mockCar)

	assert.Nil(t, err)
	mockCar = GetMockCar("KA-01-HH-9999", "White")
	err = Park(parkingLots, mockCar)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), ParkingFullError.UserMessage)
}

func TestCarUnpark(t *testing.T) {
	N := 1
	parkingLots, _ := CreateParkingLotFunc(N)
	mockCar := GetMockCar("KA-01-HH-1234", "White")
	Park(parkingLots, mockCar)
	slot, _ := parkingLots.GetSlotById(1)
	err := Leave(parkingLots, slot.Id)

	assert.Nil(t, err)
	assert.Equal(t, len(parkingLots.EmptyParkinglots), 1)
	assert.True(t, slot.IsAvailable)
	assert.Equal(t, parkingLots.TotalParkedCars, 0)
	assert.Nil(t, slot.ParkedCar)
}

func TestCarUnpark_EmptyParkingSlot(t *testing.T) {
	N := 1
	parkingLots, _ := CreateParkingLotFunc(N)

	slot, _ := parkingLots.GetSlotById(1)
	err := Leave(parkingLots, slot.Id)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), ParkingLotEmptyError.UserMessage)

}

func TestCarUnpark_ParkingSlotNotExists(t *testing.T) {
	N := 1
	parkingLots, _ := CreateParkingLotFunc(N)

	_, ok := parkingLots.GetSlotById(2)
	mockCar := GetMockCar("KA-01-HH-1234", "White")
	Park(parkingLots, mockCar)

	assert.False(t, ok)
	err := Leave(parkingLots, 2)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), NoSlotByIDFound.UserMessage)

}

func TestCarUnpark_ParkingEmpty(t *testing.T) {
	N := 1
	parkingLots, _ := CreateParkingLotFunc(N)

	slot, ok := parkingLots.GetSlotById(1)
	mockCar := GetMockCar("KA-01-HH-1234", "White")
	Park(parkingLots, mockCar)
	assert.True(t, ok)

	slot.IsAvailable = true
	err := Leave(parkingLots, 1)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), SlotEmptyError.UserMessage)

}

func TestUnParkCar_NilParkingLot(t *testing.T) {
	err := Leave(nil, 1)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), ParkingLotNotDefined.UserMessage)
}

func TestGetNearestEmptyParkingSLot(t *testing.T) {
	N := 3
	parkingLots, _ := CreateParkingLotFunc(N)

	mockCar := GetMockCar("KA-01-HH-1234", "White")
	Park(parkingLots, mockCar)
	mockCar = GetMockCar("KA-01-HH-9999", "White")
	Park(parkingLots, mockCar)
	mockCar = GetMockCar("KA-01-HH-7778", "White")
	Park(parkingLots, mockCar)

	err := Leave(parkingLots, 2)
	err = Leave(parkingLots, 3)

	assert.Nil(t, err)
	slot := GetNearestEmptyParkingSLot(*parkingLots)
	assert.Equal(t, slot.Id, 2)

}

func TestGetRegistrationNumbersByColor(t *testing.T) {
	expectedRegNums := []string{
		"KA-01-HH-1234",
		"KA-01-HH-9999",
	}
	N := 3
	parkingLots, _ := CreateParkingLotFunc(N)

	mockCar := GetMockCar("KA-01-HH-1234", "White")
	Park(parkingLots, mockCar)
	mockCar = GetMockCar("KA-01-HH-9999", "White")
	Park(parkingLots, mockCar)

	regNums, err := GetRegistrationNumbersByColor(*parkingLots, "White")
	sort.Strings(expectedRegNums)
	sort.Strings(regNums)

	assert.Nil(t, err)
	assert.Equal(t, len(regNums), 2)
	assert.EqualValues(t, regNums, expectedRegNums)
	assert.Equal(t, regNums, expectedRegNums)

}

func TestGetRegistrationNumbersByColor_NotFound(t *testing.T) {
	N := 2
	parkingLots, _ := CreateParkingLotFunc(N)

	mockCar := GetMockCar("KA-01-HH-1234", "White")
	Park(parkingLots, mockCar)
	mockCar = GetMockCar("KA-01-HH-9999", "White")
	Park(parkingLots, mockCar)

	regNums, err := GetRegistrationNumbersByColor(*parkingLots, "Blue")

	assert.Nil(t, regNums)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), NoRegistrationByColorFound.UserMessage)

}

func TestGetSlotByRegNumber(t *testing.T) {
	expectedSlotId := 1

	N := 3
	parkingLots, _ := CreateParkingLotFunc(N)

	mockCar := GetMockCar("KA-01-HH-1234", "White")
	Park(parkingLots, mockCar)
	mockCar = GetMockCar("KA-01-HH-9999", "White")
	Park(parkingLots, mockCar)

	slot, err := GetSlotByRegNumber(*parkingLots, "KA-01-HH-1234")

	assert.Nil(t, err)
	assert.Equal(t, slot, expectedSlotId)

}

func TestGetSlotByRegNumber_NotFound(t *testing.T) {
	N := 2
	parkingLots, _ := CreateParkingLotFunc(N)

	mockCar := GetMockCar("KA-01-HH-1234", "White")
	Park(parkingLots, mockCar)
	mockCar = GetMockCar("KA-01-HH-9999", "White")
	Park(parkingLots, mockCar)

	slot, err := GetSlotByRegNumber(*parkingLots, "KA-01-HH-1235")

	assert.NotNil(t, err)
	assert.Equal(t, slot, 0)
	assert.Equal(t, err.Error(), NoSlotByRegistrationNumberFound.UserMessage)

}

func TestGetSlotsByColor(t *testing.T) {
	expectedSlots := []int{1, 2}

	N := 3
	parkingLots, _ := CreateParkingLotFunc(N)

	mockCar := GetMockCar("KA-01-HH-1234", "White")
	Park(parkingLots, mockCar)
	mockCar = GetMockCar("KA-01-HH-9999", "White")
	Park(parkingLots, mockCar)
	mockCar = GetMockCar("KA-01-HH-7778", "Blue")
	Park(parkingLots, mockCar)

	slots, err := GetSlotsByColor(*parkingLots, "White")
	sort.Ints(expectedSlots)
	sort.Ints(slots)

	assert.Nil(t, err)
	assert.Equal(t, len(slots), 2)
	assert.EqualValues(t, slots, expectedSlots)
}

func TestGetSlotsByColor_NotFound(t *testing.T) {
	N := 2
	parkingLots, _ := CreateParkingLotFunc(N)

	mockCar := GetMockCar("KA-01-HH-1234", "White")
	Park(parkingLots, mockCar)
	mockCar = GetMockCar("KA-01-HH-9999", "White")
	Park(parkingLots, mockCar)

	_, err := GetSlotsByColor(*parkingLots, "Red")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), NoSlotsByColorFound.UserMessage)
}
