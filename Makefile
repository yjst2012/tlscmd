BINARY=tserver
CLIENT_BINARY=tclient

VERSION=1.0.0
BUILD=`git rev-parse HEAD`

# ToDo: set verions stuffs in files
# Setup the -ldflags option for go build here, interpolate the variable values
# LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"


build:
	cd server && go get -t ./...
	go build -o ${BINARY} server/main.go
	cd client && go get -t ./...
	go build -o ${CLIENT_BINARY} client/main.go


run:
	./${BINARY}

install:
	go install

clean:
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi
	if [ -f ${CLIENT_BINARY} ]; then rm ${CLIENT_BINARY}; fi

.PHONY: build run test install clean