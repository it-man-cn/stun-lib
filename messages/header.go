package messages

import (
	"encoding/binary"
	attr "github.com/it-man-cn/stun-lib/attributes"
	"math/rand"
	"time"
)

type header struct {
	Mtype         MessageType
	Len           uint16
	TransactionID []byte //16byte
	Attributes    map[attr.AttributeType]attr.Attribute
}

func (h *header) encode() ([]byte, error) {
	buf := make([]byte, 20)
	binary.BigEndian.PutUint16(buf[0:2], uint16(h.Mtype))
	binary.BigEndian.PutUint16(buf[2:4], h.Len)
	copy(buf[4:20], h.TransactionID)
	return buf, nil
}

func (h *header) decode(src []byte) (int, error) {
	h.Mtype = MessageType(binary.BigEndian.Uint16(src[0:2]))
	h.Len = binary.BigEndian.Uint16(src[2:4])
	h.TransactionID = src[4:20]
	return 20, nil
}
func (h *header) SetAttributes(atrributes map[attr.AttributeType]attr.Attribute) {
	h.Attributes = atrributes
}

func (h *header) AddAttribute(atrribute attr.Attribute) {
	h.Attributes[atrribute.Type()] = atrribute
}

func (h *header) GetAttribute(attribTyte attr.AttributeType) attr.Attribute {
	return h.Attributes[attribTyte]
}

func (h *header) GenerateTransactionID() {
	h.TransactionID = make([]byte, 16)
	rand.Seed(time.Now().UnixNano())
	binary.BigEndian.PutUint16(h.TransactionID[0:2], uint16(rand.Intn(65536)))
	binary.BigEndian.PutUint16(h.TransactionID[2:4], uint16(rand.Intn(65536)))
	binary.BigEndian.PutUint16(h.TransactionID[4:6], uint16(rand.Intn(65536)))
	binary.BigEndian.PutUint16(h.TransactionID[6:8], uint16(rand.Intn(65536)))
	binary.BigEndian.PutUint16(h.TransactionID[8:10], uint16(rand.Intn(65536)))
	binary.BigEndian.PutUint16(h.TransactionID[10:12], uint16(rand.Intn(65536)))
	binary.BigEndian.PutUint16(h.TransactionID[12:14], uint16(rand.Intn(65536)))
	binary.BigEndian.PutUint16(h.TransactionID[14:16], uint16(rand.Intn(65536)))
}
