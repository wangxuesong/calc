package main
import "fmt"
import "os"
import "read"

func main(){
    args := read.Checkinput(os.Args)
    if "" == args {
        fmt.Println("usage: ./calc <expression>")
        return
    }
    fmt.Println("input string: ",args);
}
