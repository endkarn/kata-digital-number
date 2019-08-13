package number

import (
	"strconv"
	"strings"
)

func SplitNumberToArray(num int) []int {
	var intArray []int
	stringNum := strconv.Itoa(num)
	stringNumArray := strings.Split(stringNum, "")
	for _ , aStringNum := range stringNumArray{
		aIntNum,_ := strconv.Atoi(aStringNum)
		intArray = append(intArray, aIntNum)
	}
	return intArray
}


func ConvertIntegerIntoArrayTextDigiFormat(num int) []string{
	if num == 1 {
		return []string{
			"|",
			"|",
			"|",
			"|",
		}
	}
	if num == 2 {
		return []string{
			"---",
			" _|",
			"|  ",
			"---",
		}
	}
	if num == 3 {
		return []string{
			"---",
			"  |",
			"--|",
			"---",
		}
	}
	return []string{}
}