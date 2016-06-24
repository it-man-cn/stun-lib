package attributes

import (
	"fmt"
	"strconv"
	"strings"
)

//Address ip address
type Address struct {
	FirstOctet   int
	SencondOctet int
	ThirdOctet   int
	FourthOctet  int
}

//Array2Address 4 bytes convert to an ipv4 address
func Array2Address(ip []byte) Address {
	address := Address{}
	address.FirstOctet = int(ip[0] & 0xFF)
	address.SencondOctet = int(ip[1] & 0xFF)
	address.ThirdOctet = int(ip[2] & 0xFF)
	address.FourthOctet = int(ip[3] & 0xFF)
	return address
}

//String2Address ip(like 192.168.1.1) conver to an ipv4 address
func String2Address(addr string) Address {
	ip := strings.Split(addr, ".")
	address := Address{}
	address.FirstOctet, _ = strconv.Atoi(ip[0])
	address.SencondOctet, _ = strconv.Atoi(ip[1])
	address.ThirdOctet, _ = strconv.Atoi(ip[2])
	address.FourthOctet, _ = strconv.Atoi(ip[3])
	return address
}

//Address2Array  an ipv4 address convert to 4 bytes
func Address2Array(address Address) []byte {
	ip := make([]byte, 4)
	ip[0] = byte(address.FirstOctet & 0xFf)
	ip[1] = byte(address.SencondOctet & 0xFf)
	ip[2] = byte(address.ThirdOctet & 0xFf)
	ip[3] = byte(address.FourthOctet & 0xFf)
	return ip
}

func (a Address) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a.FirstOctet, a.SencondOctet, a.ThirdOctet, a.FourthOctet)
}
