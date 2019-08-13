package number

import "testing"
import "github.com/stretchr/testify/assert"

func Test_ConvertIntegerToDigitalNumber_By_1_Should_Be_1(t *testing.T) {
	expected := "|\n|"

	actual := ConvertIntegerToDigitalNumber(1)

	assert.Equal(t, expected, actual)
}
