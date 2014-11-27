package testCalc

import "testing"
import "eval"

func TestMapping(t *testing.T) {
	var str string = "1+1"
	ret := eval.Mapping(str)
	if 2 != ret {
		t.Errorf("eval.Mapping(\"1+1\") failed. Get %d, except 2", ret)
	}
}
func TestMapping1(t *testing.T) {
	var str string = "1+2"
	ret := eval.Mapping(str)
	if 3 != ret {
		t.Errorf("eval.Mapping(\"1+2\") failed. Get %d, except 3", ret)
	}
}
func TestMapping2(t *testing.T) {
	var str string = "2+1"
	ret := eval.Mapping(str)
	if 3 != ret {
		t.Errorf("eval.Mapping(\"2+1\") failed. Get %d, except 3", ret)
	}
}
func TestMapping3(t *testing.T) {
	var str string = "2+2"
	ret := eval.Mapping(str)
	if 4 != ret {
		t.Errorf("eval.Mapping(\"2+2\") failed. Get %d, except 4", ret)
	}
}
func TestMapping4(t *testing.T) {
	var str string = "testttt"
	ret := eval.Mapping(str)
	if -1 != ret {
		t.Errorf("eval.Mapping(\"testttt\") failed. Get %d, except -1", ret)
	}
}
