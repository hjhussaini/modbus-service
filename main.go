package main

import (
    "os"
    "os/signal"
    "syscall"
)

func main() {
    stop := make(chan os.Signal)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

    <- stop
}
