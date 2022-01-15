package repository

type Adapter interface {
    Modbus() modbus.Adapter
}

type adapter struct {
    modbusAdapter   modbus.ModbusAdapter
}

func (dao *adapter) Modbus() modbus.Adapter {
    return &dao.modbusAdapter
}

func New() Adapter {
    return &adapter{}
}
