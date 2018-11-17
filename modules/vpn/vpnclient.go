// Copyright (C) 2018  Artur Fogiel
// This file is part of goP2PVPN.
//
// goP2PVPN is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// goP2PVPN is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with goP2PVPN.  If not, see <http://www.gnu.org/licenses/>.
package vpn

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"strings"
)

type VPNClient struct {
	debug   bool
	conn    *net.UDPConn
	do_work bool
	key     string
	address string
}

func NewVPNClient() *VPNClient {
	return &VPNClient{false, nil, false, "", ""}
}

func (vpc *VPNClient) Connect(_address string, _key string) error {
	vpc.address = _address
	conn, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4zero, Port: 0})

	if err != nil {
		vpc.do_work = false
		return err
	} else {
		vpc.conn = conn
		vpc.do_work = true
		vpc.key = _key

		go vpc.start()
		fmt.Println("")
		return nil
	}
}

func (vpc *VPNClient) start() {
	buff := make([]byte, 2048)
	ServerAddr, _ := net.ResolveUDPAddr("udp4", vpc.address)
	// send to socket
	fmt.Println("Client sending key: " + vpc.key)
	//_, err := vpc.conn.Write([]byte(vpc.key))
	_, err := vpc.conn.WriteToUDP([]byte(vpc.key), ServerAddr)
	if err != nil {
		fmt.Println("client unable to send data " + err.Error())
	}
	// listen for reply
	_, err = bufio.NewReader(vpc.conn).Read(buff)
	if err == nil {
		vpc.handlePeer(string(buff))
	} else {
		fmt.Printf("Some error %v\n", err)
	}
}

func (vpc *VPNClient) handlePeer(address string) {
	fmt.Println("address: " + address)
	PeerAddr, err := net.ResolveUDPAddr("udp4", address)

	if err != nil {
		fmt.Println("client resolve udp addr: " + err.Error())
	} else {
		buff := make([]byte, 2048)
		fmt.Println("client punching hole to " + PeerAddr.String() + " via " + vpc.conn.LocalAddr().String())
		_, err = vpc.conn.WriteToUDP([]byte("client\n"), PeerAddr)
		if err != nil {
			fmt.Println("client unable to send data " + err.Error())
		}

		for i := 0; i < 3; i++ {
			vpc.conn.WriteToUDP([]byte{CMD_CLIENT_HELLO, 0x00}, PeerAddr)
		}
		vpc.conn.WriteToUDP([]byte{CMD_READY, 0x00}, PeerAddr)

		for vpc.do_work {
			n, addr, error := vpc.conn.ReadFromUDP(buff)
			if error == nil {
				msg := string(buff[0:n])
				cmd := buff[0]
				fmt.Printf("Client got message from peer: %s %s\n", addr.String(), msg)

				if cmd == CMD_READY {
					cmdBytes := []byte("ls /tmp/")
					bytesToSend := []byte{CMD_EXEC_SHELL}
					bytesToSend = append(bytesToSend, cmdBytes...)

					vpc.conn.WriteToUDP(bytesToSend, PeerAddr)
				}
			} else {
				fmt.Printf("Some error %v\n", error)
			}
		}
	}
}

func (vpc *VPNClient) UploadFile(path string) {

}

func (vpc *VPNClient) DownloadFile(path string) {

}

func (vpc *VPNClient) Disconnect() {
	DBG("Client disconnecting")
	vpc.do_work = false
	if vpc.conn != nil {
		vpc.conn.Close()
	}
}

func (vpc *VPNClient) GetPublicIP(r *http.Request) string {
	// source: https://husobee.github.io/golang/ip-address/2015/12/17/remote-ip-go.html
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		addresses := strings.Split(r.Header.Get(h), ",")
		// march from right to left until we get a public address
		// that will be the address right before our proxy.
		for i := len(addresses) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(addresses[i])
			// header can contain spaces too, strip those out.
			realIP := net.ParseIP(ip)
			if !realIP.IsGlobalUnicast() || isPrivateSubnet(realIP) {
				// bad address, go to next
				continue
			}
			return ip
		}
	}
	return ""
}
