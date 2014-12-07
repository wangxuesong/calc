package main

import "fmt"
import "os"

import "read"
import "eval"

func main() {
	if len(os.Args) > 1 {
		fmt.Println("usage: ./calc <expression>")
		return
	}
	var exit bool = false
	fmt.Println("Please input expression:")
	for false == exit {
		exp := read.Readinput()
		fmt.Print("input string: ", exp)
		if "\n" == exp {
			exit = true
			fmt.Println("exiting...")
			return
		}
		args := read.Checkinput(exp)
		operator, operand := eval.Splitinput(args)
		if operator == "" {
			fmt.Println("can't get valid input. please check your input!\n")
			continue
		}
		var ret int = eval.Mapping(operator, operand)
		if -127 == ret {
			fmt.Println("please try again!\n")
			continue
		}
		fmt.Println(args, "=", ret, "\n")
		//		fmt.Println("=", ret)
	}
	return
}
