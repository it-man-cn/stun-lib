package attributes

import (
	"encoding/binary"
	"fmt"
)

/*
 *  0                   1                   2                   3
 *  0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
 * +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
 * |x x x x x x x x|    Family     |           Port                |
 * +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
 * |                             Address                           |
 * +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
 */

//ResponseAddress atrribute
type ResponseAddress struct {
	Port    int
	Address Address
}

//NewResponseAddress create a ResponseAddress attribute
func NewResponseAddress() *ResponseAddress {
	attr := &ResponseAddress{}
	return attr
}

//Decode decode ResponseAddress attribute
func (a *ResponseAddress) Decode(src []byte) (int, error) {
	//binary.BigEndian.Uint16(src[0:2]) type
	//len := binary.BigEndian.Uint16(src[2:4]) length
	//src[4:5]  x x x x x x x x
	//src[5:6]	Family
	a.Port = int(binary.BigEndian.Uint16(src[6:8]))
	//src[8:12] //Address
	a.Address = Array2Address(src[8:12])
	return 12, nil
}

//Encode encode ResponseAddress message
func (a *ResponseAddress) Encode() (buf []byte, err error) {
	total := 2 + 2 + 8
	buf = make([]byte, total)
	binary.BigEndian.PutUint16(buf[0:2], uint16(a.Type()))
	binary.BigEndian.PutUint16(buf[2:4], uint16(a.Length()))
	buf[4] = 0x00
	buf[5] = 0x01                                        //family
	binary.BigEndian.PutUint16(buf[6:8], uint16(a.Port)) //port
	copy(buf[8:12], Address2Array(a.Address))
	return buf, nil
}

//Length get len of atrribute (tlv)
func (a *ResponseAddress) Length() int {
	return 8
}

//Type get attribute type
func (a *ResponseAddress) Type() AttributeType {
	return RESPONSEADDRESS
}

func (a *ResponseAddress) String() string {
	return fmt.Sprintf("%s:%d", a.Address.String(), a.Port)
}
