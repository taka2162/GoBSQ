package main

import (
	"os"
	"piscine"
)

func main() {
	Argc := len(os.Args)
	if Argc == 1 {
		piscine.HandleStdInput()
	} else if Argc > 1 {
		piscine.HandleFileInputs(Argc)
	}
}
