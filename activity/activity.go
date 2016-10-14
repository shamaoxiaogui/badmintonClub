// Package activity implements the struct used to record

package activity

import (
	"errors"
	"fmt"
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
func (a Activity) YardNumber() int {
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
