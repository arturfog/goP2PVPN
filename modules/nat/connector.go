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

package main

import (
	"fmt"
	"net"
	"strconv"
)

type Connector struct {
	connectionsList map[string]*net.UDPAddr
}

func (con *Connector) init() {
	con.connectionsList = make(map[string]*net.UDPAddr)
}

func (con *Connector) addToList(key string, addr *net.UDPAddr) {
	if con.connectionsList[key] == nil {
		con.connectionsList[key] = addr
	}
}

func (con *Connector) isOnList(key string) bool {
	if con.connectionsList[key] != nil {
		return true
	}
	return false
}

func (con *Connector) exchangeIP(key string, clientAddr *net.UDPAddr, serverConn *net.UDPConn) {
	server := con.connectionsList[key]
	client := clientAddr

	serverConn.WriteToUDP([]byte(server.String()), client)
	serverConn.WriteToUDP([]byte(client.String()), server)
}

func (con *Connector) listen(port int) {
	ServerAddr, _ := net.ResolveUDPAddr("udp4", ":"+strconv.Itoa(port))
	ServerConn, _ := net.ListenUDP("udp4", ServerAddr)
	buf := make([]byte, 1024)
	fmt.Println("Listenning on port: " + strconv.Itoa(port))
	for {
		n, addr, error := ServerConn.ReadFromUDP(buf)
		if error == nil {
			key := string(buf[0:n])

			fmt.Println("Key received:", key, "from address: ", addr)
			if con.isOnList(key) == false {
				con.addToList(key, addr)
			} else {
				con.exchangeIP(key, addr, ServerConn)
			}
		}
	}
}

func (con *Connector) ban(address string) {

}

func main() {
	c := Connector{}
	c.init()
	c.listen(8081)
}
