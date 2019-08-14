package number

import "testing"
import "github.com/stretchr/testify/assert"

func TestConvertIntegerIntoArrayTextDigiFormat_By_1_Should_Be_ArrayString_1(t *testing.T) {
	expected := [4]string{
		"|",
		"|",
		"|",
		"|",
	}

	actual := ConvertIntegerIntoArrayTextDigiFormat(1)

	assert.Equal(t, expected, actual)
}

func TestSplitNumberToArray_By_12_Shoud_Be_1_2(t *testing.T) {
	expected := []int{1, 2}

	actual := SplitNumberToArray(12)

	assert.Equal(t, expected, actual)
}

func TestGetSetOfDigitalNumber_By_1_2_Should_Be_ArrayString_1_2(t *testing.T) {
	expected := [][4]string{
		[4]string{
			"|",
			"|",
			"|",
			"|",
		}, [4]string{
			"---",
			" _|",
			"|  ",
			"---",
		},
	}

	actual := GetSetOfDigitalNumber([]int{1, 2})

	assert.Equal(t, expected, actual)
}

func TestRenderSetOfDigitalNumberToString_By_1_Should_Be_LCD_1(t *testing.T) {
	expected := "|\t\n|\t\n|\t\n|\t\n"

	actual := RenderSetOfDigitalNumberToString([][4]string{
		[4]string{
			"|",
			"|",
			"|",
			"|",
		},
	})

	assert.Equal(t, expected, actual)
}

func TestRenderSetOfDigitalNumberToString_By_12_Should_Be_LCD_12(t *testing.T) {
	expected := "|\t---\t\n|\t _|\t\n|\t|  \t\n|\t---\t\n"

	actual := RenderSetOfDigitalNumberToString([][4]string{
		[4]string{
			"|",
			"|",
			"|",
			"|",
		},
		[4]string{
			"---",
			" _|",
			"|  ",
			"---",
		},
	})

	assert.Equal(t, expected, actual)
}
