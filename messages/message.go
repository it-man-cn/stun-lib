package messages

import (
	"fmt"
	attr "github.com/it-man-cn/stun-lib/attributes"
)

//MessageType is the stun message types.
type MessageType uint16

const (
	//BINDINGREQUEST request binding
	BINDINGREQUEST MessageType = 0x0001
	//BINDINGRESPONSE  binding response
	BINDINGRESPONSE MessageType = 0x0101
	//BINDINGERRORRESPONSE binding error
	BINDINGERRORRESPONSE MessageType = 0x0111
)

//Message interface
type Message interface {
	Encode() ([]byte, error)
	Decode([]byte) (int, error)
	SetAttributes(atrributes map[attr.AttributeType]attr.Attribute)
	AddAttribute(atrribute attr.Attribute)
	GetAttribute(attribTyte attr.AttributeType) attr.Attribute
}

// New creates a new Message based on the Message type.
func (m MessageType) New() (Message, error) {
	switch m {
	case BINDINGREQUEST:
		return NewBindingRequest(), nil
	case BINDINGRESPONSE:
		return NewBindingResponse(), nil
	}
	return nil, fmt.Errorf("msgtype/NewMessage: Invalid message type %d", m)
}
