package piscine

import (
	"ft"
	"os"
)

const BUFFERSIZE = 1024

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
		buf := OpenRead(i)
		if buf == nil {
			PrintError()
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
	MapData := MapSetup(input)
	if IsError(MapData) || MapData.Error {
		PrintError()
		return
	}
	ShowAnswer(MapData)
}

func ShowAnswer(MapData MapInfo) {
	MAP := InitializeMap(MapData)
	Answer := AnswerMap(MAP, MapData)
	PrintAnswer(Answer)
}
