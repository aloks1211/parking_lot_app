package utils

import "github.com/parking_lot/src/models"

func GetMockCar(regNum string, color string) models.Car{
	car := models.Car{
		RegNumber: regNum,
		Color: color,
	}
	return car
}