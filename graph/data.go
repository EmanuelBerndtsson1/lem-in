package graph

type Room struct {
	Name     string
	X        int
	Y        int
	AntCount int
	Ant      int
}

type Farm struct {
	FileData string
	Ants     int
	Start    string
	End      string
	Rooms    map[string]*Room
	Links    map[string][]string
}

type Path struct {
	Rooms []string
	Ants  int
}

func CreateFarm(fileData string) Farm {
	return Farm{Ants: 0, Start: "", End: "", FileData: fileData, Rooms: make(map[string]*Room), Links: make(map[string][]string)}
}

func (p Path) Length() int {
	return len(p.Rooms) - 2
}

func (p Path) ToString() string {
	var res string
	for i, room := range p.Rooms {
		if i != len(p.Rooms)-1 {
			res += room + " "
		} else {
			res += room
		}
	}
	return res
}
