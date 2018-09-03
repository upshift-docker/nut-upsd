#!/bin/sh -ex

if [ -z "$API_PASSWORD" ]
then
   API_PASSWORD=$(dd if=/dev/urandom bs=18 count=1 2>/dev/null | base64)
fi

if [ -z "$ADMIN_PASSWORD" ]
then
   ADMIN_PASSWORD=$(dd if=/dev/urandom bs=18 count=1 2>/dev/null | base64)
fi

cat >/etc/nut/ups.conf <<EOF
[$UPS_NAME]
	desc = "$UPS_DESC"
	driver = $UPS_DRIVER
	port = $UPS_PORT
EOF

cat >/etc/nut/upsd.conf <<EOF
LISTEN 0.0.0.0 3493
EOF

cat >/etc/nut/upsd.users <<EOF
[admin]
	password = $ADMIN_PASSWORD
	actions = set
	actions = fsd
	instcmds = all

[monitor]
	password = $API_PASSWORD
	upsmon master
EOF

cat >/etc/nut/upsmon.conf <<EOF
MONITOR $UPS_NAME@localhost 1 monitor $API_PASSWORD master
SHUTDOWNCMD "$SHUTDOWN_CMD"
EOF

chgrp -R nut /etc/nut /dev/bus/usb
chmod -R o-rwx /etc/nut

/usr/sbin/upsdrvctl start
/usr/sbin/upsd
exec /usr/sbin/upsmon -D
