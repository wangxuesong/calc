package read

import "fmt"
import "strings"

func Checkinput(Args string) string {
	//	if len(Args) == 2 {
	//		fmt.Printf("please input one expression!")
	//		return ""
	//	}
	//	if len(Args) > 4 {
	//		fmt.Println("[Checkinput] expression error! Please check your input!")
	//		return ""
	//	}
	if Args == "" {
		fmt.Printf("please input one expression!")
		return ""
	}
	var exp string
	str := strings.Fields(Args)
	for i := 0; i < len(str); i++ {
		exp = strings.Join(str[0:], "")
	}
	return exp
}
