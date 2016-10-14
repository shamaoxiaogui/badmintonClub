package badmintonClub

import (
	"fmt"
	"strings"

	act "github.com/shamaoxiaogui/badmintonClub/activity"
	osfunc "github.com/shamaoxiaogui/badmintonClub/orderStrategy"
	plut "github.com/shamaoxiaogui/badmintonClub/priceStrategy"
)

func GenerateSummary(input string) (string, error) {
	var inputs []string
	var ret string
	var n int
	for i := 0; i < len(input)-1 && n > -1; i += n + 1 {
		n = strings.IndexByte(input[i:], '\n')
		inputs = append(inputs, input[i:i+n])
	}
	// fmt.Println(inputs)
	var (
		income  int
		payment int
		profit  int
	)
	orderMethod := osfunc.OSFactory()
	timelut, pricelut := plut.LUTFactory()
	a := act.NewActivity(orderMethod, timelut, pricelut)
	ret += fmt.Sprintf("[Summary]\n\n")
	for i := 0; i < len(inputs); i++ {
		if err := a.Parser(inputs[i]); err != nil {
			// fmt.Println(err)
			return ret, err
		}
		a.CalcProfit()
		income += a.Income()
		payment += a.Payment()
		profit += a.Profit()
		ret += fmt.Sprintln(a)
		// fmt.Println("debug: " + inputs[i])
	}
	ret += fmt.Sprintf("\nTotal Income: %d\n", income)
	ret += fmt.Sprintf("Total Payment: %d\n", payment)
	ret += fmt.Sprintf("Profit: %d\n", profit)
	return ret, nil
}
