package piscine

func Len(sRune []rune) int {
	count := 0
	for range sRune {
		count++
	}
	return count
}

func sLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func Index(s, Tofind string) int {
	sRune := []rune(s)
	Findr := []rune(Tofind)
	FindLen := Len(Findr)

	for i := range sRune {
		if i+FindLen > Len(sRune) {
			return -1
		}
		if string(sRune[i:i+FindLen]) == string(Findr) {
			return i
		}
	}
	return -1
}

func Split(s, sep string) []string {
	var strArray []string
	sepLen := sLen(sep)
	tmpStr := s

	for {
		indx := Index(tmpStr, sep)
		if indx == -1 {
			strArray = append(strArray, tmpStr)
			break
		}
		strArray = append(strArray, tmpStr[:indx])
		tmpStr = tmpStr[indx+sepLen:]
	}
	return strArray
}
