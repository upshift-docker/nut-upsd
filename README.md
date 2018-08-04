# Network UPS Tools server

Docker image for Network UPS Tools server.

## Usage

This image provides a complete UPS monitoring service (USB driver only).

Start the container:

```console
# docker run \
	--name nut-upsd \
	--detach \
	--publish 3493:3493 \
	--device /dev/bus/usb/xxx/yyy \
	--env SHUTDOWN_CMD="my-shutdown-command-from-container" \
	upshift/nut-upsd
```

## Auto configuration via environment variables

This image supports customization via environment variables.

### UPS_NAME

*Default value*: `ups`

The name of the UPS.

### UPS_DESC

*Default value*: `Eaton 5SC`

This allows you to set a brief description that upsd will provide to clients that ask for a list of connected equipment.

### UPS_DRIVER

*Default value*: `usbhid-ups`

This specifies which program will be monitoring this UPS.

### UPS_PORT

*Default vaue*: `auto`

This is the serial port where the UPS is connected.

### API_USER

*Default vaue*: `upsmon`

This is the username used for communication between upsmon and upsd processes.

### API_PASSWORD

*Default vaue*: `secret`

This is the password for the upsmon user.

### SHUTDOWN_CMD

*Default vaue*: `echo 'System shutdown not configured!'`

This is the command upsmon will run when the system needs to be brought down. The command will be run from inside the container.

