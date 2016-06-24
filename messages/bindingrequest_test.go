package messages

import (
	"encoding/hex"
	"fmt"
	attr "github.com/it-man-cn/stun-lib/attributes"
	"testing"
)

func TestBindingRequest(t *testing.T) {
	fmt.Println("encode...")
	request := NewBindingRequest()
	request.GenerateTransactionID()

	buf, err := request.Encode()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(request.TransactionID))
	fmt.Println(hex.EncodeToString(buf))
	fmt.Println("decode...")
	req := NewBindingRequest()
	_, err = req.Decode(buf)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(req.TransactionID))
	if string(request.TransactionID) != string(req.TransactionID) {
		t.Fail()
	}
	/*
		ra := request.GetAttribute(attr.RESPONSEADDRESS)
		cr := request.GetAttribute(attr.CHANGEREQUEST)
		crb := request.GetAttribute(attr.CONNECTIONREQUESTBINDING)
	*/
}

func TestPassword(t *testing.T) {
	fmt.Println("encode...")
	request := NewBindingRequest()
	request.GenerateTransactionID()
	password := attr.NewPassword()
	password.Password = "14586025"
	fmt.Println(password.String())
	request.AddAttribute(password)
	buf, err := request.Encode()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(buf))
	fmt.Println("decode...")
	req := NewBindingRequest()
	_, err = req.Decode(buf)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(req.TransactionID))
	passwd := req.GetAttribute(attr.PASSWORD)
	if passwd == nil {
		t.Fail()
	}
	pa := passwd.(*attr.Password)
	fmt.Println(pa.String())
	if pa.Password != password.Password {
		t.Fail()
	}

}

func TestResponseAddress(t *testing.T) {
	fmt.Println("encode...")
	request := NewBindingRequest()
	request.GenerateTransactionID()
	ra := attr.NewResponseAddress()
	ra.Port = 80
	ra.Address = attr.String2Address("192.168.1.1")
	fmt.Println(ra.String())
	request.AddAttribute(ra)
	buf, err := request.Encode()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(buf))
	fmt.Println("decode...")
	req := NewBindingRequest()
	_, err = req.Decode(buf)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(req.TransactionID))
	a := req.GetAttribute(attr.RESPONSEADDRESS)
	if a == nil {
		t.Fail()
	}
	fmt.Println(a)
	ra1 := a.(*attr.ResponseAddress)
	fmt.Println(ra1.String())
	if ra.Port != ra1.Port && ra.Address != ra1.Address {
		t.Fail()
	}
}

func TestChangeRequest(t *testing.T) {
	fmt.Println("encode...")
	request := NewBindingRequest()
	request.GenerateTransactionID()
	cr := attr.NewChangeRequest()
	cr.ChangeIP = true
	cr.ChangePort = false
	request.AddAttribute(cr)
	fmt.Println(cr.String())
	buf, err := request.Encode()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(buf))
	fmt.Println("decode...")
	req := NewBindingRequest()
	_, err = req.Decode(buf)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(req.TransactionID))
	a := req.GetAttribute(attr.CHANGEREQUEST)
	if a == nil {
		t.Fail()
	}
	cr1 := a.(*attr.ChangeRequest)
	fmt.Println(cr1)
	if cr1.ChangeIP != cr.ChangeIP && cr1.ChangePort != cr.ChangePort {
		t.Fail()
	}
}

func TestConnectionRequestBinding(t *testing.T) {
	fmt.Println("encode...")
	request := NewBindingRequest()
	request.GenerateTransactionID()
	crb := attr.NewConnectionRequestBinding()
	crb.Value = "test"
	request.AddAttribute(crb)
	fmt.Println(crb)
	buf, err := request.Encode()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(buf))
	fmt.Println("decode...")
	req := NewBindingRequest()
	_, err = req.Decode(buf)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(req.TransactionID))
	for t, v := range req.Attributes {
		fmt.Println(t, v)
	}
	a := req.GetAttribute(attr.CONNECTIONREQUESTBINDING)
	fmt.Println(a)
	if a == nil {
		t.Fail()
	}
	crb1 := a.(*attr.ConnectionRequestBinding)
	fmt.Println(crb1)
	if crb1.Value != crb.Value {
		t.Fail()
	}
}
