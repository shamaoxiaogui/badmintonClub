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
