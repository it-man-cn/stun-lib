package attributes

import (
	"encoding/binary"
	"fmt"
)

/*  0                   1                   2                   3
 *  0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
 * +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
 * |0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 A B 0|
 * +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
 */

//ChangeRequest atrribute
type ChangeRequest struct {
	ChangeIP   bool
	ChangePort bool
}

//NewChangeRequest create a ChangeRequest attribute
func NewChangeRequest() *ChangeRequest {
	attr := &ChangeRequest{}
	return attr
}

//Decode decode ChangeRequest attribute
func (a *ChangeRequest) Decode(src []byte) (int, error) {
	//binary.BigEndian.Uint16(src[0:2]) type
	//len := binary.BigEndian.Uint16(src[2:4]) //length
	//skip 3 byte
	switch src[7] {
	case 0:
		break
	case 2:
		a.ChangePort = true
	case 4:
		a.ChangeIP = true
	case 6:
		a.ChangeIP = true
		a.ChangePort = true
	}
	return 8, nil
}

//Encode encode ChangeRequest message
func (a *ChangeRequest) Encode() (buf []byte, err error) {
	total := 2 + 2 + 4
	buf = make([]byte, total)
	binary.BigEndian.PutUint16(buf[0:2], uint16(a.Type()))
	binary.BigEndian.PutUint16(buf[2:4], uint16(a.Length()))
	buf[4] = 0x00
	buf[5] = 0x00
	buf[6] = 0x00
	if a.ChangeIP && a.ChangePort {
		buf[7] = 6
	} else if a.ChangeIP {
		buf[7] = 4
	} else if a.ChangePort {
		buf[7] = 2
	} else {
		buf[7] = 0x00
	}
	return buf, nil
}

//Length get len of atrribute (tlv)
func (a *ChangeRequest) Length() int {
	return 4
}

//Type get attribute type
func (a *ChangeRequest) Type() AttributeType {
	return CHANGEREQUEST
}

func (a *ChangeRequest) String() string {
	return fmt.Sprintf("changeIP:%v,changePort:%v", a.ChangeIP, a.ChangePort)
}
