package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var revertNumber int

	tmp := x
	for tmp > 0 {
		revertNumber = revertNumber*10 + tmp%10
		tmp = tmp / 10
	}

	return revertNumber == x
}
