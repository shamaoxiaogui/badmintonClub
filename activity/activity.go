// Package activity implements the struct used to record

package activity

import (
	"errors"
	"fmt"
	"time"
)

type Activity struct {
	date    string
	begin   int // time to start
	end     int
	num     int // peoples
	income  int
	payment int
	profit  int
}

// Parser string to activity, return err if input is in wrong format
// Input Format expected is "yyyy-MM-dd HH:mm~HH:mm [num]"
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
		return err
	} else if a.begin < 0 || a.end < 0 || a.num < 0 {
		return errors.New("invailed input: with minus???")
	}

	// re-initialize
	a.income = 0
	a.payment = 0
	a.profit = 0
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

// Calculate the number of yards need to rent
// When it return 0, the process(for one Activity) can be terminated
func (a Activity) yardNumber() int {
	t, x := a.num/6, a.num%6
	var ret int
	switch t {
	case 0:
		if x < 4 {
			ret = 0
		} else {
			ret = 1
		}
	case 1:
		ret = 2
	case 2:
		fallthrough
	case 3:
		if x >= 4 {
			ret = t + 1
		} else {
			ret = t
		}
	default:
		ret = t
	}
	return ret
}

var timeLUT = [2][]int{
	{9, 12, 18, 20, 22}, //Mon.-Fri.
	{9, 12, 18, 22},     //Sat.-Sun.
}
var moneyLUT = [2][]int{
	{30, 50, 80, 60}, //Mon.-Fri.
	{40, 50, 60},     //Sat.-Sun.
}

func (a *Activity) CalcProfit() {
	numYard := a.yardNumber()

	if numYard != 0 {
		// Calc Income
		a.income = a.num * 30

		//Calc Payment
		const shortForm = "2016-10-14"
		t, _ := time.Parse(shortForm, a.date)
		index := 0 //default Mon.-Fri.
		if wd := t.Weekday(); wd == time.Sunday || wd == time.Saturday {
			index = 1
		}
		i := 0
		for ; i < len(timeLUT[index])-1; i++ {
			if a.begin >= timeLUT[index][i] && a.begin < timeLUT[index][i+1] {
				break
			}
		}
		i++
		j := a.begin
		for ; timeLUT[index][i] < a.end; i++ {
			a.payment += numYard * moneyLUT[index][i-1] * (timeLUT[index][i] - j)
			j = timeLUT[index][i]
		}
		a.payment += numYard * moneyLUT[index][i-1] * (a.end - j)

		//Calc profit
		a.profit = a.income - a.payment
	} // else the activity is terminated, income=0, payment=0, profit=0
}
