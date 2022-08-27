package romantointerger

func getDictionary() map[byte]int {
	m := make(map[byte]int)

	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	return m
}

func getIntForRomanChar(char byte) int {
	dictionary := getDictionary()

	return dictionary[char]
}

func romanToInt(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && getIntForRomanChar(s[i+1]) > getIntForRomanChar(s[i]) {
			result += getIntForRomanChar(s[i+1]) - getIntForRomanChar(s[i])
			i++
		} else {
			result += getIntForRomanChar(s[i])
		}
	}

	return result
}
