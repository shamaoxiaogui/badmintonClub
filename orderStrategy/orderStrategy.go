// Generate functions which represents the order strategy, simply change the
// OSFactory to modifiy the strategy function
package orderStrategy

import (
	act "github.com/shamaoxiaogui/badmintonClub/activity"
)

func OSFactory() act.YardNumberFunc {
	return func(num int) int {
		t, x := num/6, num%6
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
}
