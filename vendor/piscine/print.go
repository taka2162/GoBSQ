package piscine

import "ft"

func PrintError() {
	Str := "map error\n"
	for _, r := range Str {
		ft.PrintRune(rune(r))
	}
}

func PrintAnswer(Answer [][]byte) {
	for i := range Answer {
		for j := range Answer[i] {
			ft.PrintRune(rune(Answer[i][j]))
		}
		ft.PrintRune('\n')
	}
}
