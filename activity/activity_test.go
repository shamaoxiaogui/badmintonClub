package activity

import (
	"fmt"
	"testing"
)

func TestActivityParse(t *testing.T) {
	var act Activity
	ret := Activity{
		date:    "2016-06-03",
		begin:   9,
		end:     12,
		num:     14,
		income:  0,
		payment: 0,
		profit:  0,
	}
	act.Parser("2016-06-03 09:00~12:00 14")
	if act != ret {
		t.Error("Activity parser('2016-06-02 20:00~22:00 7') failed.\n" +
			fmt.Sprintf("Got date:%s begin:%d end:%d num:%d income:%d payment:%d profit:%d",
				act.date, act.begin, act.end, act.num, act.income, act.payment, act.profit))
	}
}

func TestActivityString1(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 14")
	act.profit = 3
	ret := "2016-06-03 09:00~12:00 +0 -0 +3"
	if str := act.String(); str != ret {
		t.Error("Activity String() failed. Got %s, should be %s", str, ret)
	}
}
func TestActivityString2(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 14")
	act.profit = -3
	ret := "2016-06-03 09:00~12:00 +0 -0 -3"
	if str := act.String(); str != ret {
		t.Error("Activity String() failed. Got %s, should be %s", str, ret)
	}
}
func TestActivityString3(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 14")
	ret := "2016-06-03 09:00~12:00 +0 -0 0"
	if str := act.String(); str != ret {
		t.Error("Activity String() failed. Got %s, should be %s", str, ret)
	}
}
func TestActivityYardNumber1(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 0")
	if n := act.YardNumber(); n != 0 {
		t.Error("Activity YardNumber() failed. Got %d, should be %d",
			n, 0)
	}
}
func TestActivityYardNumber2(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 3")
	if n := act.YardNumber(); n != 0 {
		t.Error("Activity YardNumber() failed. Got %d, should be %d",
			n, 0)
	}
}
func TestActivityYardNumber3(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 4")
	if n := act.YardNumber(); n != 1 {
		t.Error("Activity YardNumber() failed. Got %d, should be %d",
			n, 1)
	}
}
func TestActivityYardNumber4(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 7")
	if n := act.YardNumber(); n != 2 {
		t.Error("Activity YardNumber() failed. Got %d, should be %d",
			n, 2)
	}
}
func TestActivityYardNumber5(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 14")
	if n := act.YardNumber(); n != 2 {
		t.Error("Activity YardNumber() failed. Got %d, should be %d",
			n, 2)
	}
}
func TestActivityYardNumber6(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 16")
	if n := act.YardNumber(); n != 3 {
		t.Error("Activity YardNumber() failed. Got %d, should be %d",
			n, 3)
	}
}
func TestActivityYardNumber7(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 25")
	if n := act.YardNumber(); n != 4 {
		t.Error("Activity YardNumber() failed. Got %d, should be %d",
			n, 4)
	}
}
