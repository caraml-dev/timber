FROM alpine:latest

RUN apk update && apk add git make bash unzip libtool gcc musl-dev ca-certificates dumb-init tzdata

RUN mkdir -p /opt/observation_service

COPY . /opt/observation_service/
COPY ./bin/* /opt/observation_service/

WORKDIR /opt/observation_service
ENTRYPOINT [ "./observation-service" ]
