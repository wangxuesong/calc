package read

import "testing"

func TestCheckinput(t *testing.T) {
	args := []string{"./calc", "1+1"}
	ret := Checkinput(args)
	if ret != args[1] {
		t.Errorf("Checkinput failed. Get %s, except %s", ret, args[1])
	}
}
func TestCheckinput1(t *testing.T) {
	args := []string{"./calc"}
	ret := Checkinput(args)
	if ret != "" {
		t.Errorf("Checkinput failed. Get %s, except \"\" ", ret)
	}
}
func TestCheckinput2(t *testing.T) {
	args := []string{"./calc", "1+1", "ddd"}
	ret := Checkinput(args)
	if ret != "" {
		t.Errorf("Checkinput failed. Get %s, except \"\" ", ret)
	}
}
