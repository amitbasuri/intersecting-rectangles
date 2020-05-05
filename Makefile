build: @build

run: @build
	./bin/interesting-rectangles -file=sample_input.json

@build:
	go build -o bin/interesting-rectangles

test:
	go test ./...
