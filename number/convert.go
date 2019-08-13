package number

import (
	"strconv"
	"strings"
)

func ConvertIntegerToDigitalNumber(num int) string {
	var result string
	if num == 1 {
		return "|\n|"
	}
	if num == 2 {
		return " _\n  |\n -\n|_"
	}
	return result
}

func SplitNumberToStringArray(num int) []string {
	stringNum := strconv.Itoa(num)
	return strings.Split(stringNum, "")
}
