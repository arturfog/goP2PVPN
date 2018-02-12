package main
import "fmt"
import "./modules"

func main() {
    fmt.Println("hello world")
    ex := modules.Shell{In:  ""}
    ex.SetInput("/tmp")
    ex.Exec("wc",  "-c")
}
