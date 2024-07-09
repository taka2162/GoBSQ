package piscine

func WriteMaxSize(Map [][]int, MaxX, MaxY int) {
	size := Map[MaxX][MaxY]
	y := 0
	for y < size {
		x := 0
		for x < size {
			Map[MaxX+y][MaxY+x] = -3
			x++
		}
		y++
	}
}

func WriteSize(Map [][]int) {
	i := 1
	max := 0
	MaxX := 0
	MaxY := 0

	for Map[i][1] != -2 {
		j := 1
		for Map[i][j] != -2 {
			if Map[i][j] == 0 {
				Map[i][j] += Calculate(Map, i, j)
				if max < Map[i][j] {
					max = Map[i][j]
					MaxX = i
					MaxY = j
				}
			}
			j++
		}
		i++
	}
	WriteMaxSize(Map, MaxX, MaxY)
}

func Calculate(Map [][]int, i int, j int) int {
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
}
