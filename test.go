package main

import (
	"a/ft"
	"a/piscine"
	"fmt"
	"os"
)

type MapInfo struct {
	Map []string
	Row int
	Col int
	Empty byte
	Obstacle byte
	Full byte
}

func MapSetup(ReadBuf []byte) MapInfo {
	Setup := MapInfo{}
	// Setup.Lines := piscine.Split(string(ReadBuf), "\n")

	Map := piscine.Split(string(ReadBuf), "\n")
	Length := len(Map[0])

	Setup.Map = Map
	Setup.Row = len(Map[1])
	Setup.Col = piscine.Atoi(Map[0][:Length-3])

	Setup.Empty = Map[0][Length-3]
	Setup.Obstacle = Map[0][Length-2]
	Setup.Full = Map[0][Length-1]

	return Setup
}

func PrintError() {
	Str := "map error\n"
	for _, r := range Str {
		ft.PrintRune(rune(r))
	}
}

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
			// fmt.Printf("\x1b[32m%2s \x1b[m", string(Answer[i][j]))
			ft.PrintRune(rune(Answer[i][j]))
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

	if Argc == 1 {
        for {
            _, err := os.Stdin.Read(buf)
            if err != nil {
                return
            }
            // PrintStr(string(buf))
        }
		} else if Argc > 1 {
		for i := 1; i < Argc; i++ {
			f, err := os.Open(os.Args[i])
			if err != nil {
				PrintError()
				continue
				// os.Exit(0)
			}
			defer f.Close()
			for {
				count, _ := f.Read(buf)
				if count == 0 {
					break
				}
			}
			MapData := MapSetup(buf)
			if piscine.Duplicate(MapData.Empty, MapData.Obstacle, MapData.Full) ||
				!piscine.Friend(Map, MapData.Empty, MapData.Obstacle, MapData.Full) {
				PrintError()
				continue
			}
			for i := 1; i < len(MapData.Map)-1; i++ {
				if MapData.Row != len(MapData.Map[i]) {
					PrintError()
					continue
				}
			}

			// Col+2, Row+2のスライスを生成する
			MAP := make([][]int, MapData.Col+2)
			for i := range MAP {
				MAP[i] = make([]int, MapData.Row+2)
			}

			for y := 0; y < MapData.Col+2; y++ {
				for x := 0; x < MapData.Row+2; x++ {
					if x == 0 || y == 0 || x == MapData.Row+1 || y == MapData.Col+1 {
						MAP[y][x] = -2
					} else if y <= MapData.Col && x <= MapData.Row {
						switch MapData.Map[y][x-1] {
						case MapData.Empty:
							MAP[y][x] = 0
						case MapData.Obstacle:
							MAP[y][x] = -1
						default:
							MAP[y][x] = -2
						}
					}
				}
			}
			// Println("Initial MAP:")
			// PrintMap(MAP) // 用意

			piscine.WriteSize(MAP)

			// Println("\n\n     ------- ↓ RESULT ↓ -------\n")
			// PrintMap(MAP) // 結果。。。

			// 返す用のstring型のスライスを生成
			Answer := make([][]byte, MapData.Col)
			for i := range Answer {
				Answer[i] = make([]byte, MapData.Row)
			}

			for y := 1; y <= MapData.Col; y++ {
				for x := 1; x <= MapData.Row; x++ {
					switch MAP[y][x] {
					case -1:
						Answer[y-1][x-1] = MapData.Obstacle
					case -3:
						Answer[y-1][x-1] = MapData.Full
					default:
						Answer[y-1][x-1] = MapData.Empty
					}
				}
			}
			Printstr(Answer)
		}
	}
}