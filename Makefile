build:
	go run main.go

run: build
	./app/messaging

watch:
	reflex -s -r '\.go$$' make run