package utils

import "fmt"

type ParkingLotErrors struct {
	error

	UserMessage string
	ErrorCode   string
}

var (
	InvalidParkingLotError = ParkingLotErrors{
		UserMessage: "ParkingLot entered not valid, must be in range [1,6]",
		ErrorCode:   "INVALID_PARKINGLOT_ERROR",
	}
	ParkingFullError = ParkingLotErrors{
		UserMessage: "Sorry, parking lot is full",
		ErrorCode:   "PARKING_FULL_ERROR",
	}
	NoSlotByRegistrationNumberFound = ParkingLotErrors{
		UserMessage: "Slot with registration Number not found!!",
		ErrorCode:   "NO_SLOT_FOUND",
	}
	ParkingLotEmptyError = ParkingLotErrors{
		UserMessage: "Sorry, parking lot is empty",
		ErrorCode:   "PARKING_LOT_EMPTY",
	}
	NoSlotsByColorFound = ParkingLotErrors{
		UserMessage: "No Slots Found",
		ErrorCode:   "NO_SLOTS_FOUND",
	}
	NoRegistrationByColorFound = ParkingLotErrors{
		UserMessage: "No Registration Number Found",
		ErrorCode:   "NO_REGISTRATION_NUMBER_FOUND",
	}
	NoSlotByIDFound = ParkingLotErrors{
		UserMessage: "No Slot By Id Found",
		ErrorCode:   "NO_SLOT_FOUND",
	}
	ParkingLotNotDefined = ParkingLotErrors{
		UserMessage: "Parking lot is not defined",
		ErrorCode:   "NO_PARKING_LOT_DEFINED",
	}
	SlotEmptyError = ParkingLotErrors{
		UserMessage: "Parking slot is empty.",
		ErrorCode:   "SLOT_IS_EMPTY",
	}
)

func (e *ParkingLotErrors) Error() string {
	return fmt.Sprintf("%s" ,e.UserMessage)
}
