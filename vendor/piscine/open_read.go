package piscine

import "os"

func OpenRead(buf []byte, i int) []byte {
	f, err := os.Open(os.Args[i])
	if err != nil {
		return nil
	}
	defer f.Close()
	for {
		count, _ := f.Read(buf)
		if count == 0 {
			break
		}
	}
	return buf
}
