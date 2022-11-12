all: build

build:
	go build -o shutdown-server main.go
deploy:
	./deploy $(filter-out $@,$(MAKECMDGOALS))
.PHONY: deploy
