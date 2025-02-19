.DEFAULT_GOAL :=  run

.PHONY: build run

build:
	go build
run: build
	./currency-converter
