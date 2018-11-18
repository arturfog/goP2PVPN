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
	"crypto/rand"
	"fmt"
	"net"
	"strings"

	"../cli"
)

func DBG(msg string) {
	fmt.Println(msg)
}

type VPNServer struct {
	debug   bool
	conn    *net.UDPConn
	do_work bool
	key     string
	address string
	shell   cli.Shell
	Peer    *net.UDPAddr
}

func NewVPNServer() *VPNServer {
	return &VPNServer{false, nil, false, "", "", cli.Shell{}, nil}
}

func (vps *VPNServer) Connect(_address string, _key string) error {
	vps.address = strings.Trim(_address, "\x00")
	conn, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4zero, Port: 0})

	if err != nil {
		vps.do_work = false
		return err
	} else {
		vps.conn = conn
		vps.do_work = true
		vps.key = _key
		go vps.work()
		return nil
	}
}

func (vps *VPNServer) work() {
	buff := make([]byte, 2048)
	ServerAddr, _ := net.ResolveUDPAddr("udp4", vps.address)
	// send to socket
	fmt.Println("Server sending key: " + vps.key)
	_, err := vps.conn.WriteToUDP([]byte(vps.key), ServerAddr)
	if err != nil {
		fmt.Println("client unable to send data " + err.Error())
	}

	_, err = bufio.NewReader(vps.conn).Read(buff)
	if err == nil {
		vps.handlePeer(string(buff))
	} else {
		fmt.Printf("Some error %v\n", err)
	}
}

func (vps *VPNServer) handleMsg(code byte, msg string, peerAddr *net.UDPAddr) {
	if code == CMD_EXEC_SHELL {
		cmd := strings.Split(msg, " ")
		out := ""
		if len(cmd) > 1 {
			out = vps.shell.Exec(cmd[0], cmd[1])
		} else {
			out = vps.shell.Exec(cmd[0], "")
		}

		bytesToSend := []byte{CMD_READY}
		bytesToSend = append(bytesToSend, []byte(out)...)
		vps.conn.WriteToUDP(bytesToSend, vps.Peer)

		//fmt.Println("Output:" + out)
	}
}

func (vps *VPNServer) handlePeer(address string) {
	fmt.Println("address: " + address)
	peer, err := net.ResolveUDPAddr("udp4", strings.Trim(address, "\x00"))
	vps.Peer = peer

	if err != nil {
		fmt.Println("server resolve udp addr: " + err.Error())
	} else {
		buff := make([]byte, 2048)
		fmt.Println("server punching hole to " + vps.Peer.String() + " via " + vps.conn.LocalAddr().String())

		for i := 0; i < 3; i++ {
			vps.conn.WriteToUDP([]byte{CMD_SERVER_HELLO, 0x00}, vps.Peer)
		}
		vps.conn.WriteToUDP([]byte{CMD_READY, 0x00}, vps.Peer)

		for vps.do_work {
			n, addr, error := vps.conn.ReadFromUDP(buff)
			if error == nil {
				msg := string(buff[1:n])
				vps.handleMsg(buff[0], msg, vps.Peer)
				fmt.Printf("Server got Message from peer: %s %s\n", addr.String(), msg)
			} else {
				fmt.Printf("Some error %v\n", error)
			}
		}
	}
}

func (vps *VPNServer) Disconnect() {
	vps.do_work = false
	DBG("Server disconnecting")

}

func (vps *VPNServer) pseudo_uuid() (uuid string) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}

func (vps *VPNServer) GenKey() (key string) {
	return vps.pseudo_uuid()
}

func (vps *VPNServer) GetKey() (key string) {
	return vps.key
}
