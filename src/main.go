package main

import (
	"bufio"
	"fmt"
	"github.com/parking_lot/src/models"
	"github.com/parking_lot/src/utils"
	"os"
	"strconv"
	"strings"
)

var parkingSlot *models.ParkingSlots

func executeCommandFunc(commandList []string) {
	switch commandList[0] {
	case "create_parking_lot":
		size, err := strconv.Atoi(commandList[1])
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		parkingSlot, err = utils.CreateParkingLotFunc(size)
		if err != nil {
			fmt.Println(err.Error())
		}
	case "park":
		car := models.Car{
			RegNumber: commandList[1],
			Color:     commandList[2],
			SlotId:    0,
		}
		err := utils.Park(parkingSlot, car)
		if err != nil {
			fmt.Println(err.Error())
		}
	case "leave":
		slotId, _ := strconv.Atoi(commandList[1])
		err := utils.Leave(parkingSlot, slotId)
		if err != nil {
			fmt.Println(err.Error())
		}
	case "status":
		err := utils.Status(*parkingSlot)
		if err != nil {
			fmt.Println(err.Error())
		}

	case "slot_numbers_for_cars_with_colour":
		slots, err := utils.GetSlotsByColor(*parkingSlot, commandList[1])
		if err != nil {
			fmt.Println(err.Error())
		}
		returnedStr := ""
		for _, slot := range slots {
			returnedStr = returnedStr + strconv.Itoa(slot) + ","
		}

		fmt.Println(strings.Trim(returnedStr, ","))

	case "slot_number_for_registration_number":
		slotId, err := utils.GetSlotByRegNumber(*parkingSlot, commandList[1])
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			fmt.Println(strconv.Itoa(slotId))
		}

	case "registration_numbers_for_cars_with_colour":
		regNums, err := utils.GetRegistrationNumbersByColor(*parkingSlot, commandList[1])
		if err != nil {
			fmt.Printf(err.Error())
		}
		returnedStr := ""
		for _, regNum := range regNums {
			returnedStr = returnedStr + regNum + ", "
		}

		fmt.Println(strings.Trim(returnedStr, ", "))

	case "exit":
		os.Exit(0)
	default:
		fmt.Printf("Command '%s' not found!!\n", commandList[0])
	}
}

func InteractiveMode() {
	for {
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Err:", err.Error())
			os.Exit(1)

		}
		commandList := strings.Split(strings.Trim(command, "\n"), " ")
		executeCommandFunc(commandList)
	}
}

func FileReaderMode() {
	filename := os.Args[1:][0]

	fileExists := utils.CheckFileExists(filename)
	if !fileExists {
		os.Exit(1)
	}

	commands, err := utils.ReadFileContent(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)

	}
	for _, command := range commands {
		commandList := strings.Split(command, " ")
		executeCommandFunc(commandList)
	}

}

func main() {
	if len(os.Args[1:]) >= 1 {
		FileReaderMode()
	} else {
		InteractiveMode()

	}

}
