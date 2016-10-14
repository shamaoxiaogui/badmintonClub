package priceStrategy

import (
	act "github.com/shamaoxiaogui/badmintonClub/activity"
)

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
