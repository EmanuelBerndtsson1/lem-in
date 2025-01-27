package graph

import (
	"sort"
	"strconv"
	"strings"
)

//Reads if the number of ants in the file is valid & returns the number of ants
func ReadAntData(farm *Farm) {
	lines := strings.Split(farm.FileData, "\n")
	ants, err := strconv.Atoi(lines[0])
	if err != nil {
		PrintError("Invalid file format.\nCould not find ants")
	}
	if ants <= 0 {
		PrintError("Invalid file format.\nInvalid number of ants")
	}
	farm.Ants = ants
}

//Reads room data from file & returns the data if everything is valid
func ReadRoomData(farm *Farm) {
	lines := strings.Split(farm.FileData, "\n")[1:]
	var n []string
	var x []int
	var y []int
	for i := 0; i < len(lines); i++ {
		if isStart(lines[i]) {
			if len(farm.Start) > 0 {
				PrintError("Invalid file format.\nFound multiple start rooms")
			}
			for !isRoom(lines[i]) {
				i++
				if i >= len(lines) {
					break
				}
			}
			if i < len(lines) {
				r := parseRoom(lines[i], &n, &x, &y)
				farm.Start = r.Name
				farm.Rooms[r.Name] = r
			}
			continue
		} else if isRoom(lines[i]) {
			r := parseRoom(lines[i], &n, &x, &y)
			farm.Rooms[r.Name] = r
			continue
		} else if isLink(lines[i]) {
			link1, link2 := parseLink(lines[i])

			if link1 == link2 {
				PrintError("Invalid file format.\nRoom cannot be linked to itself")
			}

			if !checkLink(farm.Links, link1) {
				farm.Links[link1] = make([]string, 0)
			}
			if !checkLink(farm.Links, link2) {
				farm.Links[link2] = make([]string, 0)
			}

			if Contains(farm.Links[link1], link2) || Contains(farm.Links[link2], link1) {
				PrintError("Link alredy exists")
			}

			farm.Links[link1] = append(farm.Links[link1], link2)
			farm.Links[link2] = append(farm.Links[link2], link1)
			continue
		} else if isEnd(lines[i]) {
			if len(farm.End) > 0 {
				PrintError("Invalid file format.\nFound multiple end rooms")
			}
			for !isRoom(lines[i]) {
				i++
				if i >= len(lines) {
					break
				}
			}
			if i < len(lines) {
				r := parseRoom(lines[i], &n, &x, &y)
				farm.End = r.Name
				farm.Rooms[r.Name] = r
			}
			continue
		}
	}

	checkStartEnd(farm)

	// Sort end link to be first in link list for room
	for _, l := range farm.Links {
		sort.Slice(l, func(i, j int) bool {
			return l[i] == farm.End || l[j] == farm.End
		})
	}
}

//Checks if the string is start
func isStart(line string) bool {
	return strings.Compare(line, "##start") == 0
}

//Checks if the string is end
func isEnd(line string) bool {
	return strings.Compare(line, "##end") == 0
}

//Checks if the string is a room
func isRoom(str string) bool {
	room := strings.Split(str, " ")
	if len(room) != 3 || room[0][0] == 'L' || room[0][0] == '#' {
		return false
	}
	if _, err := strconv.Atoi(room[1]); err != nil {
		return false
	}
	if _, err := strconv.Atoi(room[2]); err != nil {
		return false
	}
	return true
}

//Checks if the string is a link
func isLink(str string) bool {
	link := strings.Split(str, "-")
	if len(link) != 2 || link[0][0] == '#' {
		return false
	}

	return true
}

//Checks if links between rooms are valid
func checkLink(links map[string][]string, link string) bool {
	_, found := links[link]

	return found
}

//Parses room data feom string
func parseRoom(line string, n *[]string, x *[]int, y *[]int) *Room {
	room := strings.Split(line, " ")
	roomName := room[0]
	roomX, _ := strconv.Atoi(room[1])
	roomY, _ := strconv.Atoi(room[2])

	r := &Room{Name: roomName, X: roomX, Y: roomY, AntCount: 0}

	checkRoomData(*r, *n, *x, *y)

	*n = append(*n, roomName)
	*x = append(*x, roomX)
	*y = append(*y, roomY)

	return r
}

//Parses links
func parseLink(line string) (string, string) {
	link := strings.Split(line, "-")
	return link[0], link[1]
}

//Checks if start and end rooms are valid
func checkStartEnd(farm *Farm) {
	s, foundStart := farm.Rooms[farm.Start]
	e, foundEnd := farm.Rooms[farm.End]

	if !foundStart && !foundEnd {
		PrintError("Invalid file format.\nCould not find start and end")
	} else if !foundStart && foundEnd {
		PrintError("Invalid file format.\nCould not find start")
	} else if foundStart && !foundEnd {
		PrintError("Invalid file format.\nCould not find end")
	} else if s == e {
		PrintError("Invalid file format.\nStart and end are the same room")
	} else if s.Name == e.Name {
		PrintError("Invalid file format.\nStart and end have the same name")
	} else if s.X == e.X && s.Y == e.Y {
		PrintError("Invalid file format.\nStart and end have the same coordinates")
	}
}

//Checks if room data is valid
func checkRoomData(r Room, n []string, x []int, y []int) {
	for i := 0; i < len(n); i++ {
		if r.Name == n[i] {
			PrintError("Invalid file format.\nDuplicate room name")
		} else if r.X == x[i] && r.Y == y[i] {
			PrintError("Invalid file format.\nDuplicate room coordinates")
		}
	}
}
