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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
package vpn

import (
	"fmt"
	"net"
	"bufio"
	"crypto/rand"
	"sync"
)

func DBG(msg string) {
	fmt.Println(msg)
}

type VPNServer struct {
	debug bool
	conn *net.UDPConn
	do_work bool
	key string
	waitGroup *sync.WaitGroup
	address string
}

func NewVPNServer(_wg *sync.WaitGroup) * VPNServer {
	return &VPNServer{false, nil, false, "", _wg,""}
}

func (vps *VPNServer) Connect(_address string, _key string) error{
	vps.address = _address
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 0})

	if err != nil {
		vps.do_work = false
		return err
	} else {
		vps.conn = conn
		vps.do_work = true
		vps.key = _key
		vps.waitGroup.Add(1)
		go vps.work()
		return nil
	}
}

func (vps *VPNServer) work() {
	buff :=  make([]byte, 2048)
	ServerAddr,_ := net.ResolveUDPAddr("udp",vps.address)
	// send to socket
	fmt.Println("Server sending key: " + vps.key)
	_, err := vps.conn.WriteToUDP([]byte(vps.key), ServerAddr)
	if err != nil {
		fmt.Println("client unable to send data " + err.Error())
	}
	//for vps.do_work {
	// listen for reply
	_, err = bufio.NewReader(vps.conn).Read(buff)
	if err == nil {
		vps.handlePeer(string(buff))
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	//}
	defer vps.waitGroup.Done()
}

func (vps *VPNServer) handlePeer(address string) {
	PeerAddr,_ := net.ResolveUDPAddr("udp",address)

	buff := make([]byte, 2048)
	fmt.Println("server punching hole to " + PeerAddr.String() + " via " + vps.conn.LocalAddr().String())
	vps.conn.WriteToUDP([]byte("server\n"), PeerAddr)
	vps.conn.WriteToUDP([]byte("server\n"), PeerAddr)
	vps.conn.WriteToUDP([]byte("server\n"), PeerAddr)
	fmt.Println("server waiting for messages from " + PeerAddr.String())
	for vps.do_work {
		n, addr, error := vps.conn.ReadFromUDP(buff)
		if error == nil {
			msg := string(buff[0:n])
			fmt.Printf("Message from peer: %s %s\n", addr.String(), msg)
		} else {
			fmt.Printf("Some error %v\n", error)
		}
	}

	defer vps.waitGroup.Done()
}

func (vps *VPNServer) Disconnect() {
	vps.do_work = false
	DBG("Server disconnecting")

}

func (vps *VPNServer) GetFile() {

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
