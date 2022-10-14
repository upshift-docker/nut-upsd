all: build

build:
	docker build -t nut-upsd:latest -f Dockerfile --no-cache=true .
upload:
	docker tag nut-upsd:latest nulldevil/nut-upsd:latest
	docker push nulldevil/nut-upsd:latest
up:
	docker compose up
