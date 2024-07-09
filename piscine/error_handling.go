package piscine

// import (
// 	. "fmt"
// )

func Duplicate(Empty, Obstacle, Full byte) bool {
	if Obstacle == Full || Obstacle == Empty || Full == Empty {
		return true
	}
	return false
}

func Friend(Map []string, Obstacle, Full, Empty byte) bool {
	for i := 1; i < len(Map)-1; i++ {
		for j := 0; j < len(Map[i]); j++ {
			if (Map[i][j] != Obstacle && Map[i][j] != Full && Map[i][j] != Empty) ||
				!IsPrintable(Map[i][j]) {
				return false
			}
		}
	}
	return true
}

func IsPrintable(r byte) bool {
	if r <= '~' && r >= ' ' {
		return true
	}
	return false
}
