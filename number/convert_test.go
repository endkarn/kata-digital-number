package number

import "testing"
import "github.com/stretchr/testify/assert"

func Test_ConvertIntegerToDigitalNumber_By_1_Should_Be_1(t *testing.T) {
	expected := "\n|\n|"

	actual := ConvertIntegerToDigitalNumber(1)

	assert.Equal(t, expected, actual)
}

func Test_SplitNumberToStringArray_By_12_Should_Be_1_2(t *testing.T) {
	expected := []string{"1","2"}

	actual := SplitNumberToStringArray(12)

	assert.Equal(t,expected,actual)

}

func Test_ConvertIntegerToDigitalNumber_By_2_Should_Be_2(t *testing.T) {
	expected := "\n _\n  |\n -\n|_"

	actual := ConvertIntegerToDigitalNumber(2)

	assert.Equal(t, expected, actual)
}


func Test_ConvertIntegerToDigitalNumber_By_12_Should_Be_12(t *testing.T) {
	expected := "\n|\n|\n _\n  |\n -\n|_"

	actual := ConvertIntegerToDigitalNumber(12)

	assert.Equal(t, expected, actual)
}

func Test_ConvertIntegerToDigitalNumber_By_3_Should_Be_3(t *testing.T) {
	expected := "\n _\n  |\n -|\n _"

	actual := ConvertIntegerToDigitalNumber(3)

	assert.Equal(t, expected, actual)
}
