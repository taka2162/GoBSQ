package piscine

import (
	. "fmt"
)

func Duplicate(Empty, Obstacle, Full byte) bool {
	if Obstacle == Full || Obstacle == Empty || Full == Empty {
		return true
	}
	return false
}

func Friend(Map []string, Obstacle, Full, Empty rune) bool {
	for i, s := range Map {
		for j, r := range s {
			if (r != Obstacle && r != Full && r != Empty) || !IsPrintable(r) {

				Printf("i -> %d\nj -> %d\n", i, j)
				Println(r)
				return false
			}
		}
	}
	return true
}

func IsPrintable(r rune) bool {
	if r <= '~' && r >= ' ' {
		return true
	}
	return false
}
