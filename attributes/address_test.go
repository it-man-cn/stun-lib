package attributes

import (
	"fmt"
	"testing"
)

func TestArray2Address(t *testing.T) {
	ip := []byte{192, 168, 1, 1}
	address := Array2Address(ip)
	fmt.Println(address.String())
	if address.String() != "192.168.1.1" {
		t.Fail()
	}
}

func TestString2Address(t *testing.T) {
	address := String2Address("192.168.1.1")
	if address.FirstOctet != 192 ||
		address.SencondOctet != 168 ||
		address.ThirdOctet != 1 ||
		address.FourthOctet != 1 {
		t.Fail()
	}
}

func TestAddress2Array(t *testing.T) {
	address := String2Address("192.168.1.1")
	ip := Address2Array(address)
	if address.FirstOctet != int(ip[0]) ||
		address.SencondOctet != int(ip[1]) ||
		address.ThirdOctet != int(ip[2]) ||
		address.FourthOctet != int(ip[3]) {
		t.Fail()
	}
}
