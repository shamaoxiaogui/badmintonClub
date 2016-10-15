package badmintonClub

import (
	"fmt"
	"strings"

	act "github.com/shamaoxiaogui/badmintonClub/activity"
	"github.com/shamaoxiaogui/badmintonClub/strategy"
)

func GenerateSummary(input string) (output string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("runtime error, empty the output")
			fmt.Println(err)
			output = ""
		}
	}()
	var inputs []string
	var n int
	// split the 'big' string to strings
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
	// generate strategy
	calc := strategy.NewStrategy()
	// inject the strategies into the new Activity class
	a := act.NewActivity(calc)
	output += fmt.Sprintf("[Summary]\n\n")
	// Calculate every activity parered from input string, and append them after a single output string
	for i := 0; i < len(inputs); i++ {
		a.Parser(inputs[i])
		a.CalcProfit()
		income += a.Income()
		payment += a.Payment()
		profit += a.Profit()
		output += fmt.Sprintln(a)
		// fmt.Println("debug: " + inputs[i])
	}
	output += fmt.Sprintf("\nTotal Income: %d\n", income)
	output += fmt.Sprintf("Total Payment: %d\n", payment)
	output += fmt.Sprintf("Profit: %d\n", profit)
	return
}
