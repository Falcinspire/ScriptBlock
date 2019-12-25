package symbol

import "fmt"

type AddressType int

const (
	UNIT           = 0
	PARAMETER      = 1
	CAPTURE        = 2
	IMPORTED_UNIT  = 3
	FREE_PARAMETER = 4
)

type UnitAddress struct {
	Module string
	Unit   string
	Name   string
}

// func (this *UnitAddress) String() string {
// 	return fmt.Sprintf("unit_address:\"%s:%s:%s\"", this.Module, this.Unit, this.Name)
// }

func NewUnitAddress(module string, unit string, name string) *UnitAddress {
	return &UnitAddress{module, unit, name}
}

func NewUnitAddressBox(module string, unit string, name string) *AddressBox {
	return NewAddressBox(UNIT, NewUnitAddress(module, unit, name))
}

type ParameterAddress struct {
	FunctorDepth int
	Name         string
}

func NewParameterAddress(closureDepth int, name string) *ParameterAddress {
	return &ParameterAddress{closureDepth, name}
}

func NewParameterAddressBox(closureDepth int, name string) *AddressBox {
	return NewAddressBox(PARAMETER, NewParameterAddress(closureDepth, name))
}

func (this *ParameterAddress) String() string {
	return fmt.Sprintf("parameter:\"%s\"", this.Name)
}

type CaptureAddress struct {
	Name string
}

func (this *CaptureAddress) String() string {
	return fmt.Sprintf("capture:\"%s\"", this.Name)
}

func NewCaptureAddress(name string) *CaptureAddress {
	address := new(CaptureAddress)
	address.Name = name
	return address
}

func NewCaptureAddressBox(name string) *AddressBox {
	return NewAddressBox(CAPTURE, NewCaptureAddress(name))
}

type AddressBox struct {
	Type AddressType
	Data interface{}
}

func (this *AddressBox) String() string {
	return fmt.Sprintf("(%d)%s", this.Type, this.Data)
}

func NewAddressBox(addressType AddressType, data interface{}) *AddressBox {
	address := new(AddressBox)
	address.Type = addressType
	address.Data = data
	return address
}
