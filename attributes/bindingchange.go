package attributes

import (
	"encoding/binary"
)

//BindingChange atrribute
type BindingChange struct {
	data []byte
}

//NewBindingChange create a BindingChange attribute
func NewBindingChange() *BindingChange {
	attr := &BindingChange{}
	return attr
}

//Decode decode BindingChange attribute
func (a *BindingChange) Decode(src []byte) (int, error) {
	//binary.BigEndian.Uint16(src[0:2]) type
	len := binary.BigEndian.Uint16(src[2:4]) //length
	a.data = src[4:len]
	return 4 + int(len), nil
}

//Encode encode BindingChange message
func (a *BindingChange) Encode() (buf []byte, err error) {
	total := 2 + 2 + a.Length()
	buf = make([]byte, total)
	binary.BigEndian.PutUint16(buf[0:2], uint16(a.Type()))
	binary.BigEndian.PutUint16(buf[2:4], uint16(a.Length()))
	copy(buf[4:], a.data)
	return buf, nil
}

//Length get len of atrribute (tlv)
func (a *BindingChange) Length() int {
	return len(a.data)
}

//Type get attribute type
func (a *BindingChange) Type() AttributeType {
	return CONNECTIONREQUESTBINDING
}

func (a *BindingChange) String() string {
	return string(a.data)
}
