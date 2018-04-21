package main

import (
	"./modules/cli"
	"./modules/vpn"
	"sync"
)

func main() {
	shell := cli.Shell{}
	shell.Exec("ls")

	var wg sync.WaitGroup

	vps := vpn.NewVPNServer(&wg)
	vps.Connect("127.0.0.1:8081", vps.GenKey())

	vpc := vpn.NewVPNClient(&wg)
	vpc.Connect("127.0.0.1:8081", vps.GetKey())
	wg.Wait()

	vps.Disconnect()
	vpc.Disconnect()
}
