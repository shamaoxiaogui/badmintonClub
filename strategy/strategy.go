// Package strategy implements struct respondsible for calculate the income and payment in activity
// Simply redefine a struct that implement Strategy method or inherit the DemoStrategy struct with override
// method to add new strategy
package strategy

import (
	"time"
)

// use NewStrategy function to generte xxxStrategy struct
type demoStrategy struct {
}

func NewStrategy() *demoStrategy {
	ret := &demoStrategy{}
	return ret
}

// timeLUT is an 2-D array containing time points which constructs time period with different cost
// moneyLUT is an 2-D array containing corresponding price
// Both the 2-D arrays' first slice represent weekdays' price, and the second one represent weekends's
// the size of timeLUT's element is always equal to the size of moneyLUT's
var timeLUT [2][]int = [2][]int{
	{9, 12, 18, 20, 22}, //Mon.-Fri.
	{9, 12, 18, 22},     //Sat.-Sun.
}
var moneyLUT [2][]int = [2][]int{
	{30, 50, 80, 60}, //Mon.-Fri.
	{40, 50, 60},     //Sat.-Sun.
}

// template used to parse string to time(still...fuck golang...)
const shortForm = "2006-01-02"

//order strategy, calculate the number of yard need to order
func (d demoStrategy) yardNumber(num int) int {
	t, x := num/6, num%6
	ret := t
	if x != 0 {
		switch t {
		case 0:
			if x >= 4 {
				ret = 1
			}
		case 1:
			ret = 2
		case 2:
			fallthrough
		case 3:
			if x >= 4 {
				ret += 1
			}
		}
	}
	return ret
}

// Calculate the income, payment.
// O(n) where n is the number of time period with different cost
// This method must be implemented when you creat a new strategy struct
func (d demoStrategy) Strategy(date string, begin, end, num int) (income, payment int) {
	numYard := d.yardNumber(num)
	if numYard != 0 {
		// Calc Income
		income = num * 30

		//Calc Payment
		t, _ := time.Parse(shortForm, date)
		index := 0 //default Mon.-Fri.
		if wd := t.Weekday(); wd == time.Sunday || wd == time.Saturday {
			index = 1
		}
		i := 0
		for ; i < len(timeLUT[index])-1; i++ {
			if begin >= timeLUT[index][i] && begin < timeLUT[index][i+1] {
				//find the lower boundary
				break
			}
		}
		i++
		j := begin
		// sum the payment
		for ; timeLUT[index][i] < end; i++ {
			payment += numYard * moneyLUT[index][i-1] * (timeLUT[index][i] - j)
			j = timeLUT[index][i]
		}
		// there are maybe a tail left
		payment += numYard * moneyLUT[index][i-1] * (end - j)

		//Calc profit
		// a.profit = a.income - a.payment
	} // else the activity is terminated, income=0, payment=0, profit=0
	return
}
