package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    server := &http.Server{
        Addr:   ":3128",
    }

    errs := make(chan error, 2)
    go func() {
        if err := server.ListenAndServe(); err != nil {
            errs <-err

            return
        }
        errs <-nil
    }()
    log.Printf("Started HTTP server on \n")

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
        log.Fatalf("Failed to shut server down: %v", err)
    }
    log.Println("Stopped HTTP server")
}
