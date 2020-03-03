# Go-Jek Parking Lot Assignment (Golang)

## 1. Problem Statement
Design a Parking lot which can hold `n` number of Cars. Every time a car is arrived at the parking, system needs to successfully assign the car a nearest parking slot.
Also, the system should be capable of returning queries such as:
- Registration numbers of all cars of a particular colour.
- Slot number in which a car with a given registration number is parked.
- Slot numbers of all slots where a car of a particular colour is parked.

## 2. Supported Commands

- `create_parking_lot` <`n`>   
To create a Parking lot of size `n`.

- `park` <`registration_number`> <`colour`>   
To park the car in the parking lot, Where `registration_number` is given registration number for the car and `colour` is given colour of the car.

- `leave` <`slot`>   
To unpark the car from a parking `slot`.

- `status`   
Displays the current status of all the parking lots.

- `slot_numbers_for_cars_with_colour` <`colour`>   
Displays the slot numbers in a parking lot for the given `color` of the car.

- `slot_number_for_registration_number` <`registration_number`>   
Displays the slot number in a parking lot for the given `registration_number` of the car.

- `registration_numbers_for_cars_with_colour` <`colour`>   
Displays the registration numbers of the cars for the given `colour` of the car.

## 3. Running Application(MAC OS)
Extract the `parking_lot.zip` file in your ${GOPATH}/src/github.com/ folder and `cd` to /src directory to run the below commands.

#### 3.1 Running the application in File mode:

```
make runinfilemode
```

#### 3.2 Running the application in Interactive mode:

```
make run
```

## 4. Running tests(MAC OS)

#### 4.1 For running the tests with code coverage

```
make test
```

## 5. Build Application and Run(MAC OS)

#### 5.1 To build the application

```
make build
```
Above command will create a binary `parking_lot` in `/src/bin` folder.

#### 5.2 To run from build in interactive Mode

```
./bin/parking_lot
```

#### 5.3 To run from build in file Mode
```
./bin/parking_lot <filename>
```

