include .env
export $(shell sed 's/=.*//' .env)

make run:
	go run cmd/main.go