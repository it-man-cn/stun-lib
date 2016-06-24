package attributes

import (
	"encoding/binary"
	"fmt"
)

// AttributeType is the stun attribute types.
type AttributeType uint16

const (
	//MAPPEDADDRESS MappedAddress
	MAPPEDADDRESS AttributeType = 0x0001
	//RESPONSEADDRESS ResponseAddress
	RESPONSEADDRESS AttributeType = 0x0002
	//CHANGEREQUEST ChangeRequest
	CHANGEREQUEST AttributeType = 0x0003
	//SOURCEADDRESS SourceAddress
	SOURCEADDRESS AttributeType = 0x0004
	//CHANGEDADDRESS ChangedAddress
	CHANGEDADDRESS AttributeType = 0x0005
	//USERNAME Username
	USERNAME AttributeType = 0x0006
	//PASSWORD Password
	PASSWORD AttributeType = 0x0007
	//MESSAGEINTEGRITY MessageIntegrity
	MESSAGEINTEGRITY AttributeType = 0x0008
	//ERRORCODE ErrorCode
	ERRORCODE AttributeType = 0x0009
	//UNKNOWNATTRIBUTE UnknownAttribute
	UNKNOWNATTRIBUTE AttributeType = 0x000a
	//REFLECTEDFROM ReflectedFrom
	REFLECTEDFROM AttributeType = 0x000b
	//CONNECTIONREQUESTBINDING ConnectionRequestBinding
	CONNECTIONREQUESTBINDING AttributeType = 0xC001
	//BINDINGCHANGE BindingChange
	BINDINGCHANGE AttributeType = 0xC002
	//DUMMY Dummy
	DUMMY AttributeType = 0x0000
)

//Attribute stun msg attribute
type Attribute interface {
	//getType() uint16
	//getLength() uint16
	//getVal() string
	//parse(buf []byte)
	Encode() ([]byte, error)
	Decode([]byte) (int, error)
	Type() AttributeType
	Length() int
	String() string
}

//ParseAttrs parse attributes
/*
func ParseAttrs(buf []byte) (attributes []Attribute) {
	var (
		attrType uint16
		length   uint16
	)
	if len(buf) > 0 {
		attributes = make([]Attribute, 1)
		attrType = encode.Binary.Uint16(buf[0:2])
		length = encode.Binary.Uint16(buf[2:4])
		switch attrType {
		case PASSWORD:
			passwd := &password{}
			passwd.parse(buf[4 : 4+length])
			attributes = append(attributes, passwd)
		case CONNECTIONREQUESTBINDING:
		}

	}
	fmt.Printf("in parse attrs len: %d\n", len(attributes))
	return
}
*/

// New creates a new attribute based on the attribute type.
func (a AttributeType) New() (Attribute, error) {
	switch a {
	case MAPPEDADDRESS:
		return NewMappedAddress(), nil
	case RESPONSEADDRESS:
		return NewResponseAddress(), nil
	case SOURCEADDRESS:
		return NewSourceAddress(), nil
	case CHANGEDADDRESS:
		return NewChangedAddress(), nil
	case PASSWORD:
		return NewPassword(), nil
	case CONNECTIONREQUESTBINDING:
		return NewConnectionRequestBinding(), nil
	case BINDINGCHANGE:
		return NewBindingChange(), nil
	case CHANGEREQUEST:
		return NewChangeRequest(), nil
	}
	return nil, fmt.Errorf("atrributetype/NewAttribute: Invalid attribute type %d", a)
}

//DecodeAttributes deode attribtues from bytes
func DecodeAttributes(src []byte) (attributes map[AttributeType]Attribute, n int, err error) {
	attributes = make(map[AttributeType]Attribute)
	index := 0
	for index < len(src) {
		mtype := AttributeType(binary.BigEndian.Uint16(src[index : index+2]))
		len := binary.BigEndian.Uint16(src[index+2 : index+4])
		attribute, err := mtype.New()
		if err != nil {
			fmt.Println(err)
			break
		}
		n, err = attribute.Decode(src[index : index+int(len)+4])
		if err != nil {
			fmt.Println(err)
			break
		}
		index += n
		attributes[mtype] = attribute
	}

	return attributes, index, err
}
