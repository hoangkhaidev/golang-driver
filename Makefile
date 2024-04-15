pro:
	docker build . -t go-containerized:latest
	docker compose up
local: 
	go run main.go