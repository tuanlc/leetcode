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
}
