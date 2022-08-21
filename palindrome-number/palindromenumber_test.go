package palindromenumber

import "testing"

func TestPalindromeNumber(t *testing.T) {
	t.Run("should return true if the input is a palindrome number", func(t *testing.T) {
		actual := isPalindrome(121)
		expected := true

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})

	t.Run("should return false if the input is not a palindrome number", func(t *testing.T) {
		actual := isPalindrome(122)
		expected := false

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})

	t.Run("should return false if the input is not a palindrome number", func(t *testing.T) {
		actual := isPalindrome(12221)
		expected := true

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})

	t.Run("should return false if the input is not a palindrome number", func(t *testing.T) {
		actual := isPalindrome(1221)
		expected := true

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})

	t.Run("should return false if the input is not a negative number", func(t *testing.T) {
		actual := isPalindrome(-121)
		expected := false

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})
}
