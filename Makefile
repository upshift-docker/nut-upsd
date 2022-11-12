all: build

build:
	docker build -t nut-upsd:aarch64 -f Dockerfile --no-cache=true .
upload:
	docker tag nut-upsd:aarch64 nulldevil/nut-upsd:aarch64
	docker push nulldevil/nut-upsd:aarch64
up:
	docker compose up
