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

func (adapter *ModbusAdapter) Close() error {
    return nil
}
