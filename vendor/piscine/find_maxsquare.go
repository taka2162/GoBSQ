package piscine

func WriteMaxSize(Map [][]int, MaxX, MaxY int) {
	size := Abs(Map[MaxX][MaxY])
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

func Abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return n * -1
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
			if Map[i][j] == 0 || Map[i][j] == -3 {
				switch Map[i][j] {
				case  -3:
					Map[i][j] = Calculate(Map, i, j) * -1
				default:
					Map[i][j] = Calculate(Map, i, j)
				}
				if max < Abs(Map[i][j]) {
					max = Abs(Map[i][j])
					MaxX = i
					MaxY = j
				}
				if -3 < Map[i][j] && Map[i][j] < 0 {
					Map[i][j] = -3
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
				if -3 < Map[i+y][j+x] && Map[i+y][j+x] < 0 {
					return size - 1
				}
				x++
			}
			y++
		}
		size++
	}
}
