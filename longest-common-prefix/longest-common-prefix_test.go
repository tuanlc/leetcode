package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	t.Run("should return an empty string if there is no common prefix", func(t *testing.T) {
		input := []string{"dog", "racecar", "car"}
		actual := longestCommonPrefix(input)
		expected := ""

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})

	t.Run("should return an empty string if the input is an empty array", func(t *testing.T) {
		input := []string{}
		actual := longestCommonPrefix(input)
		expected := ""

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})

	t.Run("should return an the first string if the input has only one item", func(t *testing.T) {
		input := []string{"car"}
		actual := longestCommonPrefix(input)
		expected := "car"

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})

	t.Run("should return an the common prefix", func(t *testing.T) {
		input := []string{"flower", "flow", "flight"}
		actual := longestCommonPrefix(input)
		expected := "fl"

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})
}
