package messages

import (
	"bytes"
	"fmt"
	attr "github.com/it-man-cn/stun-lib/attributes"
)

//BindingRequest message
type BindingRequest struct {
	header
}

//NewBindingRequest create a BindingRequest message
func NewBindingRequest() *BindingRequest {
	msg := &BindingRequest{}
	msg.Mtype = BINDINGREQUEST
	msg.Attributes = make(map[attr.AttributeType]attr.Attribute)
	return msg
}

//Encode encode BindingRequest message
func (m *BindingRequest) Encode() (buf []byte, err error) {
	length := 0
	//header := m.header.encode()
	//body := make([]byte)
	out := bytes.Buffer{}
	if m.Attributes != nil {
		for _, attribute := range m.Attributes {
			buf, err = attribute.Encode()
			if err != nil {
				return buf, err
			}
			length += len(buf)
			out.Write(buf)
		}
	}
	m.Len = uint16(length)
	//buf = make([]byte, 20+length)
	buf, err = m.header.encode()
	if err != nil {
		return buf, err
	}
	buf = append(buf, out.Bytes()...)
	return buf, nil
}

//Decode decode BindingRequest message
func (m *BindingRequest) Decode(src []byte) (n int, err error) {

	n, err = m.header.decode(src[0:20])
	if err != nil {
		fmt.Println(err)
		return n, err
	}
	m.Attributes, n, err = attr.DecodeAttributes(src[20:])
	if err != nil {
		fmt.Println(err)
		return n + 20, err
	}
	return n + 20, nil
}

/*
func (msg *bindingRequest) Decode(buf []byte) {
	//fmt.Printf("len%d\n", len(buf))
	fmt.Println("request decode :" + string(buf))
	msg.Type = encode.Binary.Int16(buf[0:2])
	fmt.Println(msg.Type)

	msg.Length = encode.Binary.Int16(buf[2:4])
	fmt.Printf("msg length %d\n", msg.Length)
	msg.ID = buf[4:20]
	//msg.Id = string(buf[4:20])
	fmt.Println("msg id:" + hex.EncodeToString(msg.ID))
	msg.Attributes = ParseAttrs(buf[20:msg.Length])
	fmt.Printf("out parse attrs len: %d\n", len(msg.Attributes))
	password := msg.Attributes[0].(*password)
	fmt.Println(password.Password)
}
*/
