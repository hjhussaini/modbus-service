FROM golang:1.17.2 AS build

WORKDIR /app
COPY . /app

RUN go mod tidy --compat=1.17
RUN go build -o modbus-service

RUN mkdir -p /etc/
RUN cp etc/modbus.service.yaml /etc/

EXPOSE 502
CMD ["/app/modbus-service"]
