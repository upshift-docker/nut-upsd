# Shutdown server, what is it ?

What if you have multiple hosts connected to the same UPS and want them to safely shutdown, when electricity goes down?

The shutdown server runs on host (or many hosts) and waits for api call either it's shutdown or shutdown cancel command, when the electricity fails, nut-upsd daemon with call to these shutdown servers.

**Remember to deploy shutdown-server to all hosts you are running on ups**

# How to deploy ?


```
make deploy user@target
```

If you would like to use api key for security reasons, then copy settings file and re-deploy
```
cp shutdown-server.conf-example shutdown-server.conf
```

It uses sudo, so make sure it's setup correctly

Use this shutdown-server with project https://github.com/e1z0/nut-upsd, that docker container already has support for shutdown-client and shutdown-server, just use docker environment variable **SHUTDOWN_CMD**

Example.:
```
SHUTDOWN_CMD: "/usr/local/bin/shutdown-client -action shutdown -api_key abc -hosts node0,node3,node1,node2"
```

There is [docker-compose](../docker-compose.yml) file that shows this stuff in action.
