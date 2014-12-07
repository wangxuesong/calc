package read

import "testing"

func TestCheckinput(t *testing.T) {
	args := "1+1"
	ret := Checkinput(args)
	if ret != "1+1" {
		t.Errorf("Checkinput failed. Get %s, except %s", ret, args[1])
	}
}
func TestCheckinput1(t *testing.T) {
	args := ""
	ret := Checkinput(args)
	if ret != "" {
		t.Errorf("Checkinput failed. Get %s, except \"\" ", ret)
	}
}
func TestCheckinput2(t *testing.T) {
	args := "1 + 1"
	ret := Checkinput(args)
	if ret != "1+1" {
		t.Errorf("Checkinput failed. Get %s, except \"1+1\" ", ret)
	}
}
func TestCheckinput3(t *testing.T) {
	args := "1 +1"
	ret := Checkinput(args)
	if ret != "1+1" {
		t.Errorf("Checkinput failed. Get %s, except \"1+1\" ", ret)
	}
}
func TestCheckinput4(t *testing.T) {
	args := "1+ 1"
	ret := Checkinput(args)
	if ret != "1+1" {
		t.Errorf("Checkinput failed. Get %s, except \"\" ", ret)
	}
}
