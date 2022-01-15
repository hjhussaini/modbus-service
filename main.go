package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("/etc/modbus.service.yaml")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Failed to read config file: %v", err)
    }

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
    log.Printf("Started Modbus service on :502\n")

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
