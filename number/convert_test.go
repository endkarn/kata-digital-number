package number

import "testing"
import "github.com/stretchr/testify/assert"

func Test_GetDigitalNumber_By_1_Should_Be_1(t *testing.T) {
	expected := "|\n|"

	actual := ConvertNumberToDigitalNumberString(1)

	assert.Equal(t, expected, actual)
}
