package main

import (
	"a/ft"
	"a/piscine"
	"fmt"
	. "fmt"
	"os"
)

func IsNumeric(s rune) bool {
	if s >= '0' && s <= '9' {
		return true
	}
	return false
}

func RuneSlice(Rune []rune) {
	for _, r := range Rune {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func Printstr(Answer [][]byte) {
	for i := range Answer {
		for j := range Answer[i] {
			fmt.Printf("\x1b[32m%2s \x1b[m", string(Answer[i][j]))
		}
		fmt.Println()
	}
}

func PrintMap(GRID [][]int) {
	for i := range GRID {
		for j := range GRID[i] {
			if GRID[i][j] == -2 {
				fmt.Printf("\x1b[33m%2d \x1b[m", GRID[i][j])
			} else if GRID[i][j] == -1 {
				fmt.Printf("\x1b[31m%2d \x1b[m", GRID[i][j])
			} else {
				fmt.Printf("\x1b[30m%2d \x1b[m", GRID[i][j])
			}
		}
		fmt.Println()
	}
}

func main() {
	buf := make([]byte, 1024)
	Argc := len(os.Args)
	Map := make([]string, 1024)

	if Argc == 2 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			Println("map error")
		}
		defer f.Close()
		for {
			count, _ := f.Read(buf)
			if count == 0 {
				break
			}
			// 読み取る部分
			for i := 0; i < count; i++ {
				ft.PrintRune(rune(buf[i]))
			}

			Map = piscine.Split(string(buf), "\n")
			Println(Map)

			// それぞれのモノたち
			Length := len(Map[0])

			Empty := Map[0][Length-3]
			Obstacle := Map[0][Length-2]
			Full := Map[0][Length-1]

			// Error Handling

			// Println(Map[1:])

			if piscine.Duplicate(Empty, Obstacle, Full) ||
				!piscine.Friend(Map[1:], rune(Empty), rune(Obstacle), rune(Full)) {

				if piscine.Duplicate(Empty, Obstacle, Full) {
					Println("Duplicate")
				} else if !piscine.Friend(Map[1:], rune(Empty), rune(Obstacle), rune(Full)) {
					Println("Elian")
				}
				Println("\x1b[35mmap error\x1b[0m")
				os.Exit(0)
			}

			Row := len(Map[1])
			for i := 1; i < len(Map)-1; i++ {
				if Row != len(Map[i]) {
					Println("Not Same Row")
					os.Exit(1)
					Println(Map[i])
				}
			}

			Col := piscine.Atoi(Map[0][:Length-3])

			Printf("列数（x）-> %d\n", Row)
			Printf("行数（y）-> %d\n", Col)

			// Col+2, Row+2のスライスを生成する
			MAP := make([][]int, Col+2)
			for i := range MAP {
				MAP[i] = make([]int, Row+2)
			}

			for y := 0; y < Col+2; y++ {
				for x := 0; x < Row+2; x++ {
					if x == 0 || y == 0 || x == Row+1 || y == Col+1 {
						MAP[y][x] = -2
					} else if y <= Col && x <= Row {
						switch Map[y][x-1] {
						case Empty:
							MAP[y][x] = 0
						case Obstacle:
							MAP[y][x] = -1
						default:
							MAP[y][x] = -2
						}
					}
				}
			}
			Println("Initial MAP:")
			PrintMap(MAP) // 用意

			piscine.WriteSize(MAP)

			Println("\n\n     ------- ↓ RESULT ↓ -------\n")
			PrintMap(MAP) // 結果。。。

			// 返す用のstring型のスライスを生成
			Answer := make([][]byte, Col)
			for i := range Answer {
				Answer[i] = make([]byte, Row)
			}

			for y := 1; y <= Col; y++ {
				for x := 1; x <= Row; x++ {
					switch MAP[y][x] {
					case -1:
						Answer[y-1][x-1] = Obstacle
					case -3:
						Answer[y-1][x-1] = Full
					default:
						Answer[y-1][x-1] = Empty
					}
				}
			}

			Println("\nAnswer:")
			Printstr(Answer)
		}
	}
}
