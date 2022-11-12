# Shutdown server, what is it ?

What if you have multiple hosts connected to the UPS? And want them to safely shutdown?

The shutdown server runs on host and waits for api call either it's shutdown or shutdown cancel command, when the electricity fails, nut-upsd executes shutdown client with specified parameters

Remember to deploy this service to all hosts you are running

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
