.PHONY: all generate build

all: generate build

generate:
	go run tools/gen_initall.go

build:
	go build -o regtool main.go