export APP=visitor
export LDFLAGS="-w -s"

run:
	go run -ldflags $(LDFLAGS) ./cmd server

build:
	CGO_ENABLED=1 go build -ldflags $(LDFLAGS) ./cmd

install:
	CGO_ENABLED=1 go install -ldflags $(LDFLAGS) ./cmd
