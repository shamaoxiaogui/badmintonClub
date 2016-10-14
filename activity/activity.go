// Package activity implements the struct used to record

package activity

import (
	"errors"
	"fmt"
	"time"
)

// The callback function calculate the number of yards need to rent
// When it return 0, the process(for one Activity) can be terminated
type YardNumberFunc func(int) int

// Time period with different cost
type PriceLUT [2][]int

type Activity struct {
	date    string
	begin   int // time to start
	end     int
	num     int // peoples
	income  int
	payment int
	profit  int
	// yf is order strategy, calculate the number of yard need to order
	yf YardNumberFunc
	// timeLUT is an 2-D array containing time points which constructs time period with different cost
	timeLUT PriceLUT
	// moneyLUT is an 2-D array containing corresponding price
	moneyLUT PriceLUT
	// Both the 2-D arrays' first slice represent weekdays' price, and the second one represent weekends's
}

func NewActivity(yf YardNumberFunc, tLUT PriceLUT, mLUT PriceLUT) *Activity {
	ret := &Activity{}
	ret.yf = yf
	ret.timeLUT = tLUT
	ret.moneyLUT = mLUT
	return ret
}

// Parser string to activity, return err if input is in wrong format
// Input Format expected is "yyyy-MM-dd HH:mm~HH:mm [num]"
// This method is public so that we can use it to "re-fill" a Activity
func (a *Activity) Parser(input string) error {
	// defer func() {
	//     if err := recover(); err != nil {
	//         fmt.Println(err)
	//         fmt.Println("input format error")
	//     }
	// }()
	// index := strings.Index(input, " ")
	// a.date = input[:index]
	// a.begin, _ = strconv.Atoi(input[index+1 : index+3])
	// a.end, _ = strconv.Atoi(input[index+7 : index+9])
	// a.num, _ = strconv.Atoi(input[index+13:])
	if _, err := fmt.Sscanf(input, "%s %2d:00~%2d:00 %d",
		&a.date, &a.begin, &a.end, &a.num); err != nil {
		*a = Activity{}
		return err
	} else if a.begin < 9 || a.begin > 22 || a.end < 9 ||
		a.end > 22 || a.begin > a.end || a.num < 0 {
		*a = Activity{}
		return errors.New("invailed input: with minus???")
	} else {
		const shortForm = "2006-01-02" //...fuck golang...
		if _, err := time.Parse(shortForm, a.date); err != nil {
			*a = Activity{}
			return err
		}
		// re-initialize
		a.income = 0
		a.payment = 0
		a.profit = 0
	}

	return nil
}

func (a Activity) String() string {
	str := fmt.Sprintf("%s %02d:00~%02d:00 +%d -%d ", a.date,
		a.begin, a.end, a.income, a.payment)

	if a.profit > 0 {
		str = str + "+"
	}
	return str + fmt.Sprintf("%d", a.profit)
}

// Calculate the income, payment, and profit.
// O(n) where n is the number of time period with different cost
func (a *Activity) CalcProfit() {
	numYard := a.yf(a.num)

	if numYard != 0 {
		// Calc Income
		a.income = a.num * 30

		//Calc Payment
		const shortForm = "2006-01-02"
		t, _ := time.Parse(shortForm, a.date)
		index := 0 //default Mon.-Fri.
		if wd := t.Weekday(); wd == time.Sunday || wd == time.Saturday {
			index = 1
		}
		i := 0
		for ; i < len(a.timeLUT[index])-1; i++ {
			if a.begin >= a.timeLUT[index][i] && a.begin < a.timeLUT[index][i+1] {
				//find the lower boundary
				break
			}
		}
		i++
		j := a.begin
		// sum the payment
		for ; a.timeLUT[index][i] < a.end; i++ {
			a.payment += numYard * a.moneyLUT[index][i-1] * (a.timeLUT[index][i] - j)
			j = a.timeLUT[index][i]
		}
		// there are maybe a tail left
		a.payment += numYard * a.moneyLUT[index][i-1] * (a.end - j)

		//Calc profit
		a.profit = a.income - a.payment
	} // else the activity is terminated, income=0, payment=0, profit=0
}

// Getter
func (a Activity) Income() int {
	return a.income
}
func (a Activity) Payment() int {
	return a.payment
}
func (a Activity) Profit() int {
	return a.profit
}
