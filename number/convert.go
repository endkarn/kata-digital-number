package number

import (
	"strconv"
	"strings"
)

const maxDigitalNumberRows = 4

func RenderLcdNumber(number int) string {
	spltNumbers := SplitNumberToArray(number)
	digitalLcdSet := GetSetOfDigitalNumber(spltNumbers)
	return RenderSetOfDigitalNumberToString(digitalLcdSet)
}

func RenderSetOfDigitalNumberToString(setNumber [][4]string) string {
	var rendered string
	//rows loop
	for row := 0; row < maxDigitalNumberRows; row++ {
		for _, number := range setNumber {
			rendered += number[row] + "\t"
		}
		rendered += "\n"
	}
	return rendered
}

func GetSetOfDigitalNumber(numbers []int) [][4]string {
	var digiNumberList [][4]string
	for _, number := range numbers {
		digiNumber := ConvertIntegerIntoArrayTextDigiFormat(number)
		digiNumberList = append(digiNumberList, digiNumber)
	}
	return digiNumberList
}

func SplitNumberToArray(num int) []int {
	var intArray []int
	stringNum := strconv.Itoa(num)
	stringNumArray := strings.Split(stringNum, "")
	for _, aStringNum := range stringNumArray {
		aIntNum, _ := strconv.Atoi(aStringNum)
		intArray = append(intArray, aIntNum)
	}
	return intArray
}

func ConvertIntegerIntoArrayTextDigiFormat(num int) [4]string {
	var numberSet [10][4]string
	numberSet[0] = [4]string{
		"----",
		"|  |",
		"|  |",
		"----",
	}
	numberSet[1] = [4]string{
		"|",
		"|",
		"|",
		"|",
	}
	numberSet[2] = [4]string{
		"---",
		" _|",
		"|  ",
		"---",
	}
	numberSet[3] = [4]string{
		"---",
		"  |",
		"--|",
		"---",
	}
	numberSet[4] = [4]string{
		"|  |",
		"|  |",
		"----",
		"   |",
	}
	numberSet[5] = [4]string{
		"----",
		"|___",
		"    |",
		"----",
	}
	numberSet[6] = [4]string{
		"----",
		"|   ",
		"|---",
		"|___|",
	}
	numberSet[7] = [4]string{
		"----",
		"   |",
		"   |",
		"   |",
	}
	numberSet[8] = [4]string{
		"----",
		"|  |",
		"----",
		"|__|",
	}
	numberSet[9] = [4]string{
		"----",
		"|  |",
		"----",
		"   |",
	}
	return numberSet[num]
}
