package attributes

import (
	"encoding/binary"
)

//Password atrribute
type Password struct {
	Password string
}

//NewPassword create a Password attribute
func NewPassword() *Password {
	attr := &Password{}
	return attr
}

//Decode decode Password attribute
func (a *Password) Decode(src []byte) (int, error) {
	//binary.BigEndian.Uint16(src[0:2]) type
	len := binary.BigEndian.Uint16(src[2:4]) //length
	a.Password = string(src[4 : 4+len])
	return 4 + int(len), nil
}

//Encode encode Password message
func (a *Password) Encode() (buf []byte, err error) {
	total := 2 + 2 + a.Length()
	buf = make([]byte, total)
	binary.BigEndian.PutUint16(buf[0:2], uint16(a.Type()))
	binary.BigEndian.PutUint16(buf[2:4], uint16(a.Length()))
	copy(buf[4:], a.Password)
	return buf, nil
}

//Length get len of atrribute (tlv)
func (a *Password) Length() int {
	return len([]byte(a.Password))
}

//Type get attribute type
func (a *Password) Type() AttributeType {
	return PASSWORD
}

func (a *Password) String() string {
	return a.Password
}
