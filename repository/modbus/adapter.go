package modbus

import (
    "github.com/goburrow/modbus"
)

type Adapter interface {
    Connect(protocol string, address string) error
    Close() error
}

type ModbusAdapter struct {
    rtuHandler  modbus.RTUClientHandler
    tcpHandler  modbus.TCPClientHandler
    client      modbus.Client
}

func (adapter *Modpackage) Connect(protocol string, address string) error {

    return nil
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

func (adapter *ModbusAdapter) Close() error {
    return nil
}
