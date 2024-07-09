package piscine

func Strlen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func Atoi(s string) (int, bool) {
	if Strlen(s) == 0 {
		return 0, false
	}
	num := 0
	i := 0
	sign := 1
	if rune(s[i]) == '+' {
		i++
	} else if rune(s[i]) == '-' {
		sign = -1
		i++
	}
	for i < Strlen(s) {
		check := int(s[i]) - '0'
		if check >= 0 && check <= 9 {
			num = 10*num + check
			i++
		} else {
			return 0, false
		}
	}
	return num * sign, true
}
