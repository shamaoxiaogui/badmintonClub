package strategy

import (
	"fmt"
	"testing"
)

func TestStrategyYardNumber1(t *testing.T) {
	var ds demoStrategy
	ret := 0
	if out := ds.yardNumber(0); out != ret {
		t.Error(fmt.Sprintf("demoStrategy yardNumber() test failed. Got %d, should be %d", out, ret))
	}
}
func TestStrategyYardNumber2(t *testing.T) {
	var ds demoStrategy
	ret := 0
	if out := ds.yardNumber(3); out != ret {
		t.Error(fmt.Sprintf("demoStrategy yardNumber() test failed. Got %d, should be %d", out, ret))
	}
}
func TestStrategyYardNumber3(t *testing.T) {
	var ds demoStrategy
	ret := 1
	if out := ds.yardNumber(4); out != ret {
		t.Error(fmt.Sprintf("demoStrategy yardNumber() test failed. Got %d, should be %d", out, ret))
	}
}
func TestStrategyYardNumber4(t *testing.T) {
	var ds demoStrategy
	ret := 2
	if out := ds.yardNumber(7); out != ret {
		t.Error(fmt.Sprintf("demoStrategy yardNumber() test failed. Got %d, should be %d", out, ret))
	}
}
func TestStrategyYardNumber5(t *testing.T) {
	var ds demoStrategy
	ret := 2
	if out := ds.yardNumber(14); out != ret {
		t.Error(fmt.Sprintf("demoStrategy yardNumber() test failed. Got %d, should be %d", out, ret))
	}
}
func TestStrategyYardNumber6(t *testing.T) {
	var ds demoStrategy
	ret := 3
	if out := ds.yardNumber(16); out != ret {
		t.Error(fmt.Sprintf("demoStrategy yardNumber() test failed. Got %d, should be %d", out, ret))
	}
}
func TestStrategyYardNumber7(t *testing.T) {
	var ds demoStrategy
	ret := 4
	if out := ds.yardNumber(25); out != ret {
		t.Error(fmt.Sprintf("demoStrategy yardNumber() test failed. Got %d, should be %d", out, ret))
	}
}

type testIterm struct {
	date     string
	begin    int
	end      int
	num      int
	ePayment int
}

var weekdayTests = [...]testIterm{
	{"2016-06-02", 9, 22, 4, 670},
	{"1900-01-01", 10, 14, 4, 160},
	{"2000-02-29", 13, 15, 4, 100},
	{"2050-12-30", 12, 19, 4, 380},
	{"2016-06-03", 10, 18, 4, 360},
	{"2016-06-07", 9, 22, 25, 2680},
}
var weekendTests = [...]testIterm{
	{"2016-06-04", 9, 22, 4, 660},
	{"1900-01-06", 10, 14, 4, 180},
	{"2004-02-29", 19, 21, 4, 120},
	{"2050-12-31", 12, 19, 4, 360},
	{"2016-06-05", 10, 18, 4, 380},
	{"2016-06-05", 9, 22, 25, 2640},
}

func TestStrategy1(t *testing.T) {
	var ds demoStrategy
	for _, iterm := range weekdayTests {
		if _, out := ds.Strategy(iterm.date, iterm.begin, iterm.end, iterm.num); out != iterm.ePayment {
			t.Error(fmt.Sprintf("demoStrategy Strategy() weekdaytest failed. Got %d, should be %d", out, iterm.ePayment))

		}
	}
}
func TestStrategy2(t *testing.T) {
	var ds demoStrategy
	for _, iterm := range weekendTests {
		if _, out := ds.Strategy(iterm.date, iterm.begin, iterm.end, iterm.num); out != iterm.ePayment {
			t.Error(fmt.Sprintf("demoStrategy Strategy() weekendtest failed. Got %d, should be %d", out, iterm.ePayment))

		}
	}
}
