package piscine

type MapInfo struct {
	Lines    []string
	Row      int
	Col      int
	Empty    byte
	Obstacle byte
	Full     byte
	Error    bool
}

func MapSetup(ReadBuf []byte) MapInfo {
	Setup := MapInfo{}
	Map := Split(string(ReadBuf), "\n")
	Length := len(Map[0])

	if BadLength(Map) {
		Setup.Error = true
	} else {
		Setup.Lines = Map
		Setup.Row = len(Map[1])
		col, err := Atoi(Map[0][:Length-3])
		if err == false {
			Setup.Error = true
			return Setup
		}
		Setup.Col = col

		Setup.Empty = Map[0][Length-3]
		Setup.Obstacle = Map[0][Length-2]
		Setup.Full = Map[0][Length-1]

		//最初の障害物とかと改行？
		if Setup.Col != len(Map)-2 || Setup.Row == 0 ||
			Setup.Col == 0 {
			Setup.Error = true
		}
	}
	return Setup
}

func InitializeMap(MapData MapInfo) [][]int {
	MAP := make([][]int, MapData.Col+2)
	for i := range MAP {
		MAP[i] = make([]int, MapData.Row+2)
	}

	for y := 0; y < MapData.Col+2; y++ {
		for x := 0; x < MapData.Row+2; x++ {
			if x == 0 || y == 0 || x == MapData.Row+1 || y == MapData.Col+1 {
				MAP[y][x] = -2
			} else if y <= MapData.Col && x <= MapData.Row {
				switch MapData.Lines[y][x-1] {
				case MapData.Empty:
					MAP[y][x] = 0
				case MapData.Full:
					MAP[y][x] = 0
				case MapData.Obstacle:
					MAP[y][x] = -1
				default:
					MAP[y][x] = -2
				}
			}
		}
	}
	WriteSize(MAP)
	return MAP
}

func AnswerMap(MAP [][]int, MapData MapInfo) [][]byte {
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
	return Answer
}
