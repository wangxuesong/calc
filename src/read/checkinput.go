package read

import "fmt"
import "strings"

func Checkinput(Args []string) string {
	if len(Args) == 1 {
		fmt.Printf("please input one expression!")
		return ""
	}
	if len(Args) > 4 {
		fmt.Println("[Checkinput] expression error! Please check your input!")
		return ""
	}
	var exp string = ""
	for i := 0; i < len(Args); i++ {
		exp = strings.Join(Args[1:], "")
	}
	return exp
}
