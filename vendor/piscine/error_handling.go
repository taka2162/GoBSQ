package piscine

// import . "fmt"

// mapgen.go内でのエラーハンドリング
func BadLength(Map []string) bool {
	Length := len(Map[0])
	if len(Map) < 2 || Length < 4 || !IsNumeric(Map[0][:Length-3]) {
		return true
	} else {
		return false
	}
}

func IsError(MapData MapInfo) bool {
	if Duplicate(MapData.Empty, MapData.Obstacle, MapData.Full) ||
		!Friend(MapData.Lines, MapData.Empty, MapData.Obstacle, MapData.Full) {
		return true
	}
	for i := 1; i < len(MapData.Lines)-1; i++ {
		if MapData.Row != len(MapData.Lines[i]) {
			return true
		}
	}
	return false
}

func Duplicate(Empty, Obstacle, Full byte) bool {
	if Obstacle == Full || Obstacle == Empty || Full == Empty {
		return true
	}
	return false
}

func Friend(Map []string, Obstacle, Full, Empty byte) bool {
	for i := 1; i < len(Map)-1; i++ {
		for j := 0; j < len(Map[i]); j++ {
			if (Map[i][j] != Obstacle && Map[i][j] != Full && Map[i][j] != Empty) ||
				!IsPrintable(Map[i][j]) {
				return false
			}
		}
	}
	return true
}

func IsPrintable(r byte) bool {
	if r <= '~' && r >= ' ' {
		return true
	}
	return false
}

func IsNumeric(s string) bool {
	for _, b := range s {
		if !(b >= '0' && b <= '9') {
			return false
		}
	}
	return true
}
