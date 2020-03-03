GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=parking_lot

all: test buildmake

build: test
	${GOBUILD} -o ./bin/${BINARY_NAME} main.go

test:
	${GOTEST} -v ./... -cover

run: test
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v main.go
	./bin/$(BINARY_NAME)

runinfilemode:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v main.go
	./bin/$(BINARY_NAME) file_inputs.txt