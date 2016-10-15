// Package activity implements the struct used to record

package activity

import (
	"fmt"
	"time"
)

// Strategy interface
type CalcStrategy interface {
	Strategy(date string, begin, end, num int) (int, int)
}

// force client to use NewActivity to create activity struct
type activity struct {
	date    string
	begin   int // time to start
	end     int
	num     int // peoples
	income  int
	payment int
	profit  int

	calc CalcStrategy
}

// Factory function to generate activity struct
// User must use this function insteady of activity{} (and you can't do that...)
// the calc Startegy need be transfered
func NewActivity(calc CalcStrategy) *activity {
	if calc == nil {
		panic("NewActivity:no Strategy injected")
	}
	ret := &activity{}
	ret.calc = calc
	return ret
}

// Parser string to activity, return err if input is in wrong format
// Input Format expected is "yyyy-MM-dd HH:mm~HH:mm [num]"
// This method is public so that we can use it to "re-fill" a activity
func (a *activity) Parser(input string) {
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
		panic(err)
	} else if a.begin < 9 || a.begin > 22 || a.end < 9 ||
		a.end > 22 || a.begin > a.end || a.num < 0 {
		panic("activity.Parser: invailed input: with minus???")
	} else {
		const shortForm = "2006-01-02" //...fuck golang...
		if _, err := time.Parse(shortForm, a.date); err != nil {
			*a = activity{}
			panic(err)
		}
		// re-initialize
		a.income = 0
		a.payment = 0
		a.profit = 0
	}
}

//Format Output
func (a activity) String() string {
	str := fmt.Sprintf("%s %02d:00~%02d:00 +%d -%d ", a.date,
		a.begin, a.end, a.income, a.payment)

	if a.profit > 0 {
		str = str + "+"
	}
	return str + fmt.Sprintf("%d", a.profit)
}

// Calculate the income, payment, and profit.
// Use Strategy interface do the calculation.
func (a *activity) CalcProfit() {
	a.income, a.payment = a.calc.Strategy(a.date, a.begin, a.end, a.num)
	a.profit = a.income - a.payment
}

// Getter
func (a activity) Income() int {
	return a.income
}
func (a activity) Payment() int {
	return a.payment
}
func (a activity) Profit() int {
	return a.profit
}
