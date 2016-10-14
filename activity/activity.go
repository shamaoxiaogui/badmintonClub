// Package activity implements the struct used to record

package activity

import (
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
