FROM ubuntu:14.04
MAINTAINER leodotcloud@gmail.com

RUN apt-get update && \
    apt-get install -y curl && \
    rm -rf /var/lib/apt/lists/*

COPY getLogLevel setLogLevelDebug setLogLevelInfo setLogLevelError log-example /usr/bin/

CMD log-example
