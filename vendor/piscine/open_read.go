package piscine

import "os"

func OpenRead(i int) []byte {
	var input []byte
	buf := make([]byte, 1024)

	f, err := os.Open(os.Args[i])
	if err != nil {
		return nil
	}
	defer f.Close()
	for {
		count, err := f.Read(buf)
		if count > 0 {
			input = append(input, buf[:count]...)
		}
		if err != nil {
			break
		}
	}
	return input
}
