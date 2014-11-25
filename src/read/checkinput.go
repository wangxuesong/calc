package read

import "fmt"

func Checkinput(Args []string) string {
    if len(Args) == 1{
        fmt.Println("please input one expression!")
        return ""
    }
    if len(Args) > 2{
        fmt.Println("Don't input more than one expression!")
        return ""
    }
    return Args[1]
}
