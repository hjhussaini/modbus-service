package modbus

import (
    "errors"
    "time"

    "github.com/goburrow/modbus"
)

type Adapter interface {
    Connect(protocol string, address string) error
    Close() error
}

type ModbusAdapter struct {
    rtuHandler  *modbus.RTUClientHandler
    tcpHandler  *modbus.TCPClientHandler
    client      modbus.Client
}

func (adapter *ModbusAdapter) Connect(protocol string, address string) error {
    if protocol == "rtu" {
        return adapter.rtuConnect(address)
    }

    if protocol == "tcp" {
        return adapter.tcpConnect(address)
    }

    return errors.New("Invalid protocol")
}

func (adapter *ModbusAdapter) rtuConnect(address string) error {
    adapter.rtuHandler = modbus.NewRTUClientHandler(address)
    adapter.rtuHandler.Timeout = 5 * time.Second
    adapter.rtuHandler.BaudRate = 115200
    adapter.rtuHandler.DataBits = 8
    adapter.rtuHandler.Parity = "N"
    adapter.rtuHandler.StopBits = 1

    if err := adapter.rtuHandler.Connect(); err != nil {
        return err
    }

    adapter.client = modbus.NewClient(adapter.rtuHandler)

    return nil
}

func (adapter *ModbusAdapter) tcpConnect(address string) error {
    adapter.tcpHandler = modbus.NewTCPClientHandler(address)
    adapter.tcpHandler.Timeout = 10 * time.Second
    adapter.tcpHandler.SlaveId = 0xFF

    if err := adapter.tcpHandler.Connect(); err != nil {
        return err
    }

    adapter.client = modbus.NewClient(adapter.tcpHandler)

    return nil
    
}

func (adapter *ModbusAdapter) Close() error {
    if adapter.rtuHandler != nil {
        return adapter.rtuHandler.Close()
    }

    if adapter.tcpHandler != nil {
        return adapter.tcpHandler.Close()
    }

    return nil
}
