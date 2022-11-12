# Network UPS Tools server

Docker image for Network UPS Tools server. Docker image have been remade using ubuntu 20.04 as the base, it links systemd inside container and enables shutdown functionality. Also there are some reporting to InfluxDB added. See docker compose for configuration.

**NOTE! New shutdown function have been introduced, look at [shutdown-server](../master/shutdown-server/README.md)**

## Usage

This image provides a complete UPS monitoring service (USB driver only).

Start the container:

```console
docker-compose up
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

*Default value*: `auto`

This is the serial port where the UPS is connected.

### API_USER

*Default value*: `upsmon`

This is the username used for communication between upsmon and upsd processes.

### API_PASSWORD

*Default value*: `secret`

This is the password for the upsmon user.

### SHUTDOWN_CMD

*Default value*: `echo 'System shutdown not configured!'`

You should use: `/usr/local/bin/shutdown-client -action shutdown -api_key abcd -hosts node1,node2,node3`

This is the command upsmon will run when the system needs to be brought down. The command will be run from inside the container.

### ENABLE_INFLUX

*Default value*: `false`

This will enable reporting to InfluxDB. Only v1 of http api is supported at the time of writing.

### INFLUX_DBUSER

*Default value*: `none`

InfluxDB Database user

### INFLUX_DBPASS

*Default value*: `none`

InfluxDB Database password

### INFLUX_HOST

*Default value*: `none`

InfluxDB Host

### INFLUX_PORT

*Default value*: `none`

InfluxDB Port

### INFLUX_DBNAME

*Default value*: `none`

InfluxDB Database name

### INFLUX_DISPLAY_HOST

*Default value*: `none`

There you can specifiy your server hostname or any other keyword, it defines the tag for InfluxDB measurement table.

### INFLUX_UPDATE_INTERVAL

*Default value*: `none`

How frequently write statistics to the InfluxDB, I'm advice you to use 10 (seconds) or more here.
