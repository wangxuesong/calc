package eval

import "fmt"
import "strconv"
import "errors"

func Mapping(operator string, operand []string) (result int, err error) {
	opd1, err1 := strconv.Atoi(operand[0])
	opd2, err2 := strconv.Atoi(operand[1])
	if err1 != nil || err2 != nil {
		fmt.Println("strconv.Atoi error!")
		err = errors.New("strconv.Atoi error")
		return
	}
	//var result int
	switch operator {
	case "+":
		result = opd1 + opd2
	case "-":
		result = opd1 - opd2
	default:
		fmt.Println("unknown expression!")
		err = errors.New("unknown expression")
	}
	return
}
