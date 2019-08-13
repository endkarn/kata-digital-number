package number

import "testing"
import "github.com/stretchr/testify/assert"

func TestConvertIntegerIntoArrayTextDigiFormat_By_1_Should_Be_1(t *testing.T) {
	expected := []string{
		"|",
		"|",
		"|",
		"|",
	}

	actual := ConvertIntegerIntoArrayTextDigiFormat(1)

	assert.Equal(t, expected, actual)
}

func TestSplitNumberToArray_By_12_Shoud_Be_1_2(t *testing.T) {
	expected := []int{1,2}

	actual := SplitNumberToArray(12)

	assert.Equal(t, expected, actual)
}