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
		t.Error(fmt.Sprintf("Activity String() failed. Got %s, should be %s", str, ret))
	}
}
func TestActivityString2(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 14")
	act.profit = -3
	ret := "2016-06-03 09:00~12:00 +0 -0 -3"
	if str := act.String(); str != ret {
		t.Error(fmt.Sprintf("Activity String() failed. Got %s, should be %s", str, ret))
	}
}
func TestActivityString3(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 14")
	ret := "2016-06-03 09:00~12:00 +0 -0 0"
	if str := act.String(); str != ret {
		t.Error(fmt.Sprintf("Activity String() failed. Got %s, should be %s", str, ret))
	}
}
func TestActivityyardNumber1(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 0")
	if n := act.yardNumber(); n != 0 {
		t.Error(fmt.Sprintf("Activity yardNumber() failed. Got %d, should be %d",
			n, 0))
	}
}
func TestActivityyardNumber2(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 3")
	if n := act.yardNumber(); n != 0 {
		t.Error(fmt.Sprintf("Activity yardNumber() failed. Got %d, should be %d",
			n, 0))
	}
}
func TestActivityyardNumber3(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 4")
	if n := act.yardNumber(); n != 1 {
		t.Error(fmt.Sprintf("Activity yardNumber() failed. Got %d, should be %d",
			n, 1))
	}
}
func TestActivityyardNumber4(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 7")
	if n := act.yardNumber(); n != 2 {
		t.Error(fmt.Sprintf("Activity yardNumber() failed. Got %d, should be %d",
			n, 2))
	}
}
func TestActivityyardNumber5(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 14")
	if n := act.yardNumber(); n != 2 {
		t.Error(fmt.Sprintf("Activity yardNumber() failed. Got %d, should be %d",
			n, 2))
	}
}
func TestActivityyardNumber6(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 16")
	if n := act.yardNumber(); n != 3 {
		t.Error(fmt.Sprintf("Activity yardNumber() failed. Got %d, should be %d",
			n, 3))
	}
}
func TestActivityyardNumber7(t *testing.T) {
	var act Activity
	act.Parser("2016-06-03 09:00~12:00 25")
	if n := act.yardNumber(); n != 4 {
		t.Error(fmt.Sprintf("Activity yardNumber() failed. Got %d, should be %d",
			n, 4))
	}
}

var testInput = [...]string{
	"2016-06-02 20:00~22:00 7",
	"2016-06-03 09:00~12:00 14",
	"2016-06-04 14:00~17:00 22",
	"2016-06-05 19:00~22:00 3",
	"2016-06-06 12:00~15:00 15",
	"2016-06-07 15:00~17:00 12",
	"2016-06-08 10:00~13:00 19",
	"2016-06-09 16:00~18:00 16",
	"2016-06-10 20:00~22:00 5",
	"2016-06-11 13:00~15:00 11",
}
var testExpected = [...]string{
	"2016-06-02 20:00~22:00 +210 -240 -30",
	"2016-06-03 09:00~12:00 +420 -180 +240",
	"2016-06-04 14:00~17:00 +660 -600 +60",
	"2016-06-05 19:00~22:00 +0 -0 0",
	"2016-06-06 12:00~15:00 +450 -300 +150",
	"2016-06-07 15:00~17:00 +360 -200 +160",
	"2016-06-08 10:00~13:00 +570 -330 +240",
	"2016-06-09 16:00~18:00 +480 -300 +180",
	"2016-06-10 20:00~22:00 +150 -120 +30",
	"2016-06-11 13:00~15:00 +330 -200 +130",
}

func TestActivityCalcProfit1(t *testing.T) {
	var act Activity
	for i := 0; i < len(testInput); i++ {
		act.Parser(testInput[i])
		act.CalcProfit()
		if str := act.String(); str != testExpected[i] {
			t.Error(fmt.Sprintf("Activity String() failed. Got %s, should be %s", str, testExpected[i]))
		}
	}
}
