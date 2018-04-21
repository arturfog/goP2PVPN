package cli

type Shell struct {

}

func (sh *Shell) Exec(cmd string) {
	cl := CLI{}
	switch cmd {
	case "ls":
		cl.LS("/tmp")
	}
}