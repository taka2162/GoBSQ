package main

import (
	"piscine"
	"os"
)

func main() {
	Argc := len(os.Args)
	if Argc == 1 {
		HandleStdInput()
    } else if Argc > 1 {
		HandleFileInputs(Argc)
	}
}

func HandleStdInput() {
	buf := make([]byte, 1024)
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
		buf := piscine.OpenRead(make([]byte, 1024), i)
		if buf == nil {
			piscine.PrintError()
			continue
		}
		ProcessMapData(buf)
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
