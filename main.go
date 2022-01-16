package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    "modbus-service/repository"

    "github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("/etc/modbus.service.yaml")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Failed to read config file: %v", err)
    }

    adapter := repository.New()

    mqttProtocol := viper.GetString("modbus.protocol")
    mqttAddress := viper.GetString("modbus.address")
    if err := adapter.Modbus().Connect(mqttProtocol, mqttAddress); err != nil {
            log.Fatalf("Failed to connect to %s: %v", mqttAddress, err)
    }
    defer func() {
        if err := adapter.Modbus().Close(); err != nil {
            log.Fatalf("Failed to close %s: %v", mqttAddress, err)
        }
    }()

    address := viper.GetString("server.address")
    server := &http.Server{
        Addr:   address,
    }

    errs := make(chan error, 2)
    go func() {
        if err := server.ListenAndServe(); err != nil {
            errs <-err

            return
        }
        errs <-nil
    }()
    log.Printf("Started Modbus service on %s\n", address)

    go func() {
        stop := make(chan os.Signal, 1)
        signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

        <-stop

        timeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
        defer cancel()
        if err := server.Shutdown(timeout); err != nil {
            errs <-err

            return
        }

        errs <-timeout.Err()
    }()

    if err := <-errs; err != nil {
        log.Fatal(err)
    }
    log.Println("Stopped Modbus service")
}
