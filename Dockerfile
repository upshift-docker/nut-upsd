FROM ubuntu:20.04


LABEL maintainer="e1z0"

ENV NUT_VERSION 2.7.4
ENV UPS_NAME="ups"
ENV UPS_DESC="UPS"
ENV UPS_DRIVER="usbhid-ups"
ENV UPS_PORT="auto"
ENV API_PASSWORD=""
ENV ADMIN_PASSWORD=""



ENV SHUTDOWN_CMD="echo 'System shutdown not configured!'"

RUN apt-get update && apt-get install -y build-essential wget libusb-dev curl && \
        # download and extract
        cd /tmp && \
        wget http://www.networkupstools.org/source/2.7/nut-$NUT_VERSION.tar.gz && \
        tar xfz nut-$NUT_VERSION.tar.gz && \
        cd nut-$NUT_VERSION && \
        ./configure \
                --prefix=/usr \
                --sysconfdir=/etc/nut \
                --disable-dependency-tracking \
                --enable-strip \
                --disable-static \
                --with-all=no \
                --with-usb=yes \
                --datadir=/usr/share/nut \
                --with-drvpath=/usr/share/nut \
                --with-statepath=/var/run/nut \
                --with-user=root \
                --with-group=root \
        && \
        make install && \
        # create nut user
        install -d -m 750 -o root -g root /var/run/nut && \
        rm -rf /tmp/nut-$NUT_VERSION.tar.gz /tmp/nut-$NUT_VERSION

COPY src/docker-entrypoint /usr/local/bin/
COPY src/flux_mon /usr/local/bin/
<<<<<<< HEAD
COPY shutdown-client/shutdown-client /usr/local/bin/
=======
>>>>>>> ddb1ae8e2dbd1ec5899433f850c5fcac268a01a8
ENTRYPOINT ["docker-entrypoint"]

WORKDIR /var/run/nut

EXPOSE 3493
