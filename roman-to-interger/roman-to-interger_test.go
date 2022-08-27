package romantointerger

import "testing"

func TestPalindromeNumber(t *testing.T) {
	t.Run("should return convert roman(III) to Int(3)", func(t *testing.T) {
		actual := romanToInt("III")
		expected := 3

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})

	t.Run("should return convert roman(LVIII) to Int(58)", func(t *testing.T) {
		actual := romanToInt("LVIII")
		expected := 58

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})

	t.Run("should return convert roman(MCMXCIV) to Int(1994)", func(t *testing.T) {
		actual := romanToInt("MCMXCIV")
		expected := 1994

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})
}
