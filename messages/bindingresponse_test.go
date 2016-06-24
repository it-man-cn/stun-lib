package messages

import (
	"encoding/hex"
	"fmt"
	attr "github.com/it-man-cn/stun-lib/attributes"
	"testing"
)

func TestBindingResponse(t *testing.T) {
	fmt.Println("encode...")
	response := NewBindingResponse()
	response.GenerateTransactionID()

	buf, err := response.Encode()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(response.TransactionID))
	fmt.Println(hex.EncodeToString(buf))
	fmt.Println("decode...")
	resp := NewBindingResponse()
	_, err = resp.Decode(buf)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(resp.TransactionID))
	if string(response.TransactionID) != string(resp.TransactionID) {
		t.Fail()
	}
	/*
		ma := attr.NewMappedAddress()
		ca := attr.NewChangedAddress()
		sa := attr.NewSourceAddress()
	*/
}

func TestMappedAddress(t *testing.T) {
	fmt.Println("encode...")
	response := NewBindingResponse()
	response.GenerateTransactionID()
	ma := attr.NewMappedAddress()
	ma.Port = 80
	ma.Address = attr.String2Address("192.168.1.1")
	response.AddAttribute(ma)
	buf, err := response.Encode()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(buf))
	fmt.Println("decode...")
	req := NewBindingResponse()
	_, err = req.Decode(buf)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(req.TransactionID))
	a := req.GetAttribute(attr.MAPPEDADDRESS)
	if a == nil {
		t.Fail()
	}
	ma1 := a.(*attr.MappedAddress)
	fmt.Println(ma1)
	if ma.String() != ma1.String() {
		t.Fail()
	}
}

func TestChangedAddress(t *testing.T) {
	fmt.Println("encode...")
	response := NewBindingResponse()
	response.GenerateTransactionID()
	ca := attr.NewChangedAddress()
	ca.Port = 80
	ca.Address = attr.String2Address("192.168.1.1")
	response.AddAttribute(ca)
	buf, err := response.Encode()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(buf))
	fmt.Println("decode...")
	req := NewBindingResponse()
	_, err = req.Decode(buf)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(req.TransactionID))
	a := req.GetAttribute(attr.CHANGEDADDRESS)
	if a == nil {
		t.Fail()
	}
	ca1 := a.(*attr.ChangedAddress)
	fmt.Println(ca1)
	if ca.String() != ca1.String() {
		t.Fail()
	}
}

func TestSourceAddress(t *testing.T) {
	fmt.Println("encode...")
	response := NewBindingResponse()
	response.GenerateTransactionID()
	sa := attr.NewSourceAddress()
	sa.Port = 80
	sa.Address = attr.String2Address("192.168.1.1")
	response.AddAttribute(sa)
	buf, err := response.Encode()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(buf))
	fmt.Println("decode...")
	req := NewBindingResponse()
	_, err = req.Decode(buf)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(req.TransactionID))
	a := req.GetAttribute(attr.SOURCEADDRESS)
	if a == nil {
		t.FailNow()
	}
	sa1 := a.(*attr.SourceAddress)
	fmt.Println(sa1)
	if sa.String() != sa1.String() {
		t.FailNow()
	}
}
