package main

import (
	"fmt"
	attr "github.com/it-man-cn/stun-lib/attributes"
	msg "github.com/it-man-cn/stun-lib/messages"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	PrimaryIP   string
	PrimaryPort int
)

func main() {
	var (
		addr *net.UDPAddr
		conn *net.UDPConn
		err  error
	)
	Debug := true
	PrimaryIP = "localhost"
	PrimaryPort = 3478
	bind := fmt.Sprintf("%s:%d", PrimaryIP, PrimaryPort)
	if addr, err = net.ResolveUDPAddr("udp4", bind); err != nil {
		fmt.Printf("net.ResolveUDPAddr(\"udp4\", \"%s\") error(%v)\n", bind, err)
		return
	}

	if conn, err = net.ListenUDP("udp4", addr); err != nil {
		fmt.Printf("net.ListenUDP(\"udp4\", \"%v\") error(%v)\n", addr, err)
		return
	}
	defer conn.Close()

	if Debug {
		fmt.Printf("start udp listen: \"%s\"\n", bind)
	}

	//N core accept
	for i := 0; i < 1; i++ {
		go udpAccept(conn)
	}

	//wait
	InitSignal()
}

func udpAccept(conn *net.UDPConn) {
	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("ReadFromUdp(\"%v\") error(%v) nbytes(%d)\n", addr, err, n)
		}
		if n > 0 {
			request := msg.NewBindingRequest()
			_, err := request.Decode(buf[:n])
			if err != nil {
				fmt.Printf("error:%v\n", err)
				continue
			}
			//request attributes
			ra := request.GetAttribute(attr.RESPONSEADDRESS)
			cr := request.GetAttribute(attr.CHANGEREQUEST)
			crb := request.GetAttribute(attr.CONNECTIONREQUESTBINDING)

			//create response
			resp := msg.NewBindingResponse()
			resp.TransactionID = request.TransactionID
			//reseponse attributes
			ma := attr.NewMappedAddress()
			ma.Address = attr.String2Address(addr.IP.String())
			ma.Port = addr.Port
			resp.AddAttribute(ma)
			ca := attr.NewChangedAddress()
			ca.Address = attr.String2Address(PrimaryIP)
			ca.Port = PrimaryPort
			resp.AddAttribute(ca)

			if crb != nil {
				//source address
				sa := attr.NewSourceAddress()
				sa.Address = attr.String2Address(PrimaryIP)
				sa.Port = PrimaryPort
				resp.AddAttribute(sa)
				data, err := resp.Encode()
				if err != nil {
					fmt.Printf("error:%v\n", err)
					continue
				}
				if ra != nil {
					// response address,echo to another client address
					respAddr := ra.(*attr.ResponseAddress)
					remoteAddr := &net.UDPAddr{}
					remoteAddr.Port = respAddr.Port
					remoteAddr.IP = attr.Address2Array(respAddr.Address)
					n, err = conn.WriteToUDP(data, remoteAddr)
					if err != nil {
						fmt.Printf("WriteToUDP(\"%v\") error(%v) nbytes(%d)\n", addr, err, n)
					}
				} else {
					// echo to the same client address
					n, err = conn.WriteToUDP(data, addr)
					if err != nil {
						fmt.Printf("WriteToUDP(\"%v\") error(%v) nbytes(%d)\n", addr, err, n)
					}
				}
			} else if cr != nil {
				//TODO
				// need echo from different IP or port
			} else {
				//Unknown
				fmt.Printf("Unknown\n")
			}
		}

	}
}

func InitSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		s := <-c
		fmt.Printf("stund get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			return
		case syscall.SIGHUP:
			return
		default:
			return
		}
	}
}
