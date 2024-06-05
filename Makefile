MAKEFLAGS += --silent

MAIN_PACKAGE_PATH := ./cmd/goauth

build:
	go build ${MAIN_PACKAGE_PATH}

run:
	go run ${MAIN_PACKAGE_PATH}

tidy:
	go mod tidy

clean:
	rm *.exe