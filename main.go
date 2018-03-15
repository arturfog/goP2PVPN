package main

import "./modules/cli"

func main() {
	cl := cli.CLI{}
	cli.Exec(&cl, "ls", "", false)
}
