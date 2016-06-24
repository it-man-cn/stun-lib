package attributes

import (
	"encoding/binary"
)

//ConnectionRequestBinding atrribute
type ConnectionRequestBinding struct {
	Value string
}

//NewConnectionRequestBinding create a ConnectionRequestBinding attribute
func NewConnectionRequestBinding() *ConnectionRequestBinding {
	attr := &ConnectionRequestBinding{}
	return attr
}

//Decode decode ConnectionRequestBinding attribute
func (a *ConnectionRequestBinding) Decode(src []byte) (int, error) {
	//binary.BigEndian.Uint16(src[0:2]) type
	len := binary.BigEndian.Uint16(src[2:4]) //length
	a.Value = string(src[4 : 4+len])
	return 4 + int(len), nil
}

//Encode encode ConnectionRequestBinding message
func (a *ConnectionRequestBinding) Encode() (buf []byte, err error) {
	total := 2 + 2 + a.Length()
	buf = make([]byte, total)
	binary.BigEndian.PutUint16(buf[0:2], uint16(a.Type()))
	binary.BigEndian.PutUint16(buf[2:4], uint16(a.Length()))
	copy(buf[4:], a.Value)
	return buf, nil
}

//Length get len of atrribute (tlv)
func (a *ConnectionRequestBinding) Length() int {
	return len([]byte(a.Value))
}

//Type get attribute type
func (a *ConnectionRequestBinding) Type() AttributeType {
	return CONNECTIONREQUESTBINDING
}

func (a *ConnectionRequestBinding) String() string {
	return a.Value
}
