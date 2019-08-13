package number

import (
	"strconv"
	"strings"
)

const maxDigitalNumberRows  = 4

func RenderSetOfDigitalNumberToString(setNumber [][]string) string {
	var rendered string

	rendered += "|\t\n|\t\n|\t\n|\t\n"

	return rendered
}

func GetSetOfDigitalNumber(numbers []int) [][]string {
	var digiNumberList [][]string
	for _, number := range numbers{
		digiNumber := ConvertIntegerIntoArrayTextDigiFormat(number)
		digiNumberList = append(digiNumberList,digiNumber)
	}
	return digiNumberList
}

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