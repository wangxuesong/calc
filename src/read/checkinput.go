package read

import "fmt"

func Checkinput(Args []string) string {
	if len(Args) == 1 {
		fmt.Printf("please input one expression!")
		return ""
	}
	if len(Args) > 2 {
		fmt.Println("Dont't input more than one expression!")
		return ""
	}
	return Args[1]
}
