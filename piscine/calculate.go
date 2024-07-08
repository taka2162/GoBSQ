package piscine

func WriteMaxSize(Map [][]int, i, j int) {
	size := Map[i][j]
	y := 0
	for y < size {
		x := 0
		for x < size {
			Map[i+y][j+x] = -3
			x++
		}
		y++
	}
}

func WriteSize(Map [][]int) {
	i := 1
	max := 0
	indx := make([]int, 2)

	for Map[i][1] != -2 {
		j := 1
		for Map[i][j] != -2 {
			if Map[i][j] == 0 {
				Map[i][j] += Caluculate(Map, i, j)
				if max < Map[i][j] {
					max = Map[i][j]
					indx[0] = i
					indx[1] = j
				}
			}
			j++
		}
		i++
	}
	WriteMaxSize(Map, indx[0], indx[1])
}

func Caluculate(Map [][]int, i int, j int) int {
	size := 1
	for {
		y := 0
		for y < size {
			x := 0
			for x < size {
				if Map[i+y][j+x] != 0 {
					return size - 1
				}
				x++
			}
			y++
		}
		size++
	}
	return size
}
