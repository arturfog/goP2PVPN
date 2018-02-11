import "fmt"

type Shell struct {}

func (sh *Shell) open() {

}

func (sh *Shell) exec(cmd string, args ...string) string {
    	cmd := exec.Command(cmd, args)
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
