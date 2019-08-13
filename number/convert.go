package number

import (
	"strconv"
	"strings"
)

func ConvertIntegerToDigitalNumber(num int) string {
	var result string
	for _ , aNumber := range SplitNumberToStringArray(num){
		result += "\n"
		if aNumber == "1" {
			result+= "|\n|"
		}
		if aNumber == "2" {
			result+= " _\n  |\n -\n|_"
		}
		if aNumber == "3"{
			result+= " _\n  |\n -|\n _"
		}
	}
	return result
}

func SplitNumberToStringArray(num int) []string {
	stringNum := strconv.Itoa(num)
	return strings.Split(stringNum, "")
}
