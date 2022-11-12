all: build

build:
<<<<<<< HEAD
	docker build -t nut-upsd:aarch64 -f Dockerfile --no-cache=true .
upload:
	docker tag nut-upsd:aarch64 nulldevil/nut-upsd:aarch64
	docker push nulldevil/nut-upsd:aarch64
=======
	docker build -t nut-upsd:latest -f Dockerfile --no-cache=true .
upload:
	docker tag nut-upsd:latest nulldevil/nut-upsd:latest
	docker push nulldevil/nut-upsd:latest
>>>>>>> ddb1ae8e2dbd1ec5899433f850c5fcac268a01a8
up:
	docker compose up
