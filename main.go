package main

import (
	"ft"
	"os"
	"piscine"
)

const BUFFERSIZE = 1024

func main() {
	Argc := len(os.Args)
	if Argc == 1 {
		HandleStdInput()
    } else if Argc > 1 {
		HandleFileInputs(Argc)
	}
}

func HandleStdInput() {
	buf := make([]byte, BUFFERSIZE)
	var input []byte
	for {
		n, err := os.Stdin.Read(buf)
		if n > 0 {
			input = append(input, buf[:n]...)
		}
		if err != nil {
			break
		}
	}
	ProcessMapData(input)
}

func HandleFileInputs(Argc int) {
	for i := 1; i < Argc; i++ {
		buf := piscine.OpenRead(i)
		if buf == nil {
			piscine.PrintError()
			if i != Argc-1 {
				ft.PrintRune('\n')
			}
			continue
		}
		ProcessMapData(buf)
		if i != Argc-1 {
			ft.PrintRune('\n')
		}
	}
}

func ProcessMapData(input []byte) {
	MapData := piscine.MapSetup(input)
	if piscine.IsError(MapData) || MapData.Error {
		piscine.PrintError()
		return
	}
	ShowAnswer(MapData)
}

func ShowAnswer(MapData piscine.MapInfo) {
	MAP := piscine.InitializeMap(MapData)
	Answer := piscine.AnswerMap(MAP, MapData)
	piscine.PrintAnswer(Answer)
}
