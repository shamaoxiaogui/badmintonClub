package main

import (
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/shamaoxiaogui/badmintonClub"
)

func main() {
	temp, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(temp)
	// fmt.Print(input)
	output, err := GenerateSummary(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(output)
}
