package eval

import "fmt"

func Mapping (inputstr string) int{
    var result int
    switch inputstr {
        case "1+1":
            result = 2
        case "1+2":
            result = 3
        case "2+1":
            result = 3
        case "2+2":
            result = 4
        default:
            fmt.Println("unknown expression!")
            result = -1
    }
    return result
}
