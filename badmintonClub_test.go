package badmintonClub

import (
	"fmt"
	"testing"
)

var testInput = `2016-06-02 20:00~22:00 7
2016-06-03 09:00~12:00 14
2016-06-04 14:00~17:00 22
2016-06-05 19:00~22:00 3
2016-06-06 12:00~15:00 15
2016-06-07 15:00~17:00 12
2016-06-08 10:00~13:00 19
2016-06-09 16:00~18:00 16
2016-06-10 20:00~22:00 5
2016-06-11 13:00~15:00 11
`
var expectedOutput = `[Summary]

2016-06-02 20:00~22:00 +210 -240 -30
2016-06-03 09:00~12:00 +420 -180 +240
2016-06-04 14:00~17:00 +660 -600 +60
2016-06-05 19:00~22:00 +0 -0 0
2016-06-06 12:00~15:00 +450 -300 +150
2016-06-07 15:00~17:00 +360 -200 +160
2016-06-08 10:00~13:00 +570 -330 +240
2016-06-09 16:00~18:00 +480 -300 +180
2016-06-10 20:00~22:00 +150 -120 +30
2016-06-11 13:00~15:00 +330 -200 +130

Total Income: 3630
Total Payment: 2470
Profit: 1160
`

func TestGenerateSummary1(t *testing.T) {
	if out := GenerateSummary(testInput); out != expectedOutput {
		t.Error(fmt.Sprintf("GenerateSummary() failed. Got %s, should be %s",
			out, expectedOutput))
	}
}

var illegalInputs = [...]string{
	`2016-14-11 13:00~15:00 11
	`,
	`2016-11-11 1:00~15:00 11
	`,
	`2016-11-11 11
	`,
}

func TestGenerateSummary2(t *testing.T) {
	for _, iterm := range illegalInputs {
		if out := GenerateSummary(iterm); out != "" {
			t.Error(fmt.Sprintf("GenerateSummary() illegal input test failed. Got %s, should be \"\"",
				out))
		}
	}
}
