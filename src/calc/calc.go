package main
import "fmt"
import "os"
import "read"
import "eval"

func main(){
    args := read.Checkinput(os.Args)
    if "" == args {
        fmt.Println("usage: ./calc <expression>")
        return
    }
    fmt.Println("input string: ",args);

    var ret int = eval.Mapping(args)
    if (-1 == ret){
        fmt.Println("please try again!")
        return
    }
    fmt.Println(os.Args[1],"=",ret)
}
