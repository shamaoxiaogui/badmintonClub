// Generate price related look up table, simply modify the LUTFactory to change the price or time period
package priceStrategy

import (
	act "github.com/shamaoxiaogui/badmintonClub/activity"
)

// the size of timeLUT's element is always equal to the size of moneyLUT's
// element plus one. For more information ,check the comment in package activity
func LUTFactory() (timeLUT act.PriceLUT, moneyLUT act.PriceLUT) {
	timeLUT = [2][]int{
		{9, 12, 18, 20, 22}, //Mon.-Fri.
		{9, 12, 18, 22},     //Sat.-Sun.
	}
	moneyLUT = [2][]int{
		{30, 50, 80, 60}, //Mon.-Fri.
		{40, 50, 60},     //Sat.-Sun.
	}
	return
}
