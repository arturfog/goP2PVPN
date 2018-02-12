package modules
import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Shell struct {
	In string
}

func (sh* Shell) SetInput(input string) {
	sh.In = input
}

func (sh* Shell) Exec(command string, args ...string) string {
    cmd := exec.Command(command, "")
    cmd.Args = args

    cmd.Stdin = strings.NewReader(sh.In)

    var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("in all caps: %q\n", out.String())

    return out.String()
}
