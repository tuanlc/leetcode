package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	commonPrefix := strs[0]

	for i := 1; i < len(strs); i++ {
		commonPrefix = longestCommonPrefixBetweenTwoStrings(commonPrefix, strs[i])

		if commonPrefix == "" {
			break
		}
	}

	return commonPrefix
}

func longestCommonPrefixBetweenTwoStrings(str1 string, str2 string) string {
	commonPrefix := ""
	minLen := minLenBetweenTwoStrings(str1, str2)

	for i := 0; i < minLen; i++ {
		if str1[i] != str2[i] {
			break
		}

		commonPrefix = commonPrefix + string(str1[i])
	}

	return commonPrefix
}

func minLenBetweenTwoStrings(str1 string, str2 string) int {
	if len(str1) > len(str2) {
		return len(str2)
	}

	return len(str1)
}
