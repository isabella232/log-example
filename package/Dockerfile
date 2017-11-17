FROM ubuntu:14.04
MAINTAINER leodotcloud@gmail.com

RUN apt-get update && \
    apt-get install -y curl && \
    rm -rf /var/lib/apt/lists/*

ADD loglevel.tar.gz /usr/bin

COPY log-example /usr/bin/

CMD log-example
