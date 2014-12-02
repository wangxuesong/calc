package main

import "fmt"
import "os"
import "read"
import "eval"

func main() {
	args := read.Checkinput(os.Args)
	if "" == args {
		fmt.Println("usage: ./calc <expression>")
		return
	}
	fmt.Println("input string: ", args)
	operator, operand := eval.Splitinput(args)
	if operator == "" {
		fmt.Println("can't get valid input. please check your input!")
		return
	}
	var ret int = eval.Mapping(operator, operand)
	if -1 == ret {
		fmt.Println("please try again!")
		return
	}
	fmt.Println(args, "=", ret)
}
