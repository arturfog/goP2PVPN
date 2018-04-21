package proxy

import (
	"net"
	"bufio"
	"fmt"
)

type Proxy struct {

}

func (proxy *Proxy) start() {

}
// The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
// If the port in the address parameter is empty or "0", as in
// "127.0.0.1:" or "[::1]:0", a port number is automatically chosen.
func (proxy *Proxy) ListenAndServe(network, addr string) error {
	l, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	return proxy.Serve(l)
}

// Serve is used to serve connections from a listener
func (proxy *Proxy) Serve(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go proxy.ServeConn(conn)
	}
	return nil
}

// ServeConn is used to serve a single connection.
func (proxy *Proxy) ServeConn(conn net.Conn) error {
	defer conn.Close()
	bufConn := bufio.NewReader(conn)

	// Read the version byte
	version := []byte{0}
	if _, err := bufConn.Read(version); err != nil {
		s.config.Logger.Printf("[ERR] socks: Failed to get version byte: %v", err)
		return err
	}

	// Ensure we are compatible
	if version[0] != socks5Version {
		err := fmt.Errorf("Unsupported SOCKS version: %v", version)
		s.config.Logger.Printf("[ERR] socks: %v", err)
		return err
	}
}


func (proxy *Proxy) stop() {

}