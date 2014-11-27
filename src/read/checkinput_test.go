package testCalc

import "read"
import "testing"

func TestCheckinput(t *testing.T) {
	args := []string{"./calc", "1+1"}
	ret := read.Checkinput(args)
	if ret != args[1] {
		t.Errorf("Checkinput failed. Get %s, except %s", ret, args[1])
	}
}
func TestCheckinput1(t *testing.T) {
	args := []string{"./calc"}
	ret := read.Checkinput(args)
	if ret != "" {
		t.Errorf("Checkinput failed. Get %s, except \"\" ", ret)
	}
}
func TestCheckinput2(t *testing.T) {
	args := []string{"./calc", "1+1", "ddd"}
	ret := read.Checkinput(args)
	if ret != "" {
		t.Errorf("Checkinput failed. Get %s, except \"\" ", ret)
	}
}
