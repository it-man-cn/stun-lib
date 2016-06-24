package messages

import (
	"bytes"
	"fmt"
	attr "github.com/it-man-cn/stun-lib/attributes"
)

//BindingResponse message
type BindingResponse struct {
	header
}

//NewBindingResponse create a BindingResponse message
func NewBindingResponse() *BindingResponse {
	msg := &BindingResponse{}
	msg.Mtype = BINDINGRESPONSE
	msg.Attributes = make(map[attr.AttributeType]attr.Attribute)
	return msg
}

//Encode encode BindingResponse message
func (m *BindingResponse) Encode() (buf []byte, err error) {
	length := 0
	//header := m.header.encode()
	//body := make([]byte)
	out := bytes.Buffer{}
	if m.Attributes != nil {
		for _, attribute := range m.Attributes {
			buf, err = attribute.Encode()
			if err != nil {
				fmt.Println(err)
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
		fmt.Println(err)
		return buf, err
	}
	buf = append(buf, out.Bytes()...)
	return buf, nil
}

//Decode decode BindingResponse message
func (m *BindingResponse) Decode(src []byte) (n int, err error) {
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
