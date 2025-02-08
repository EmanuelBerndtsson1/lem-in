package graph

import (
	"sort"
)

// Function to find all the paths that are not the shortest path even if the dont have an end room
func FindAllPaths(startRoom string, endRoom string, links map[string][]string) (map[string][][]string, [][]Path) {
	paths := make(map[string][][]string)
	keys := make([]string, 0)
	valid := false

	for _, l := range links[startRoom] {
		paths[l] = make([][]string, 0)
		findAllPaths2(paths, l, l, endRoom, links, []string{startRoom, l}, []string{startRoom, l})
		keys = append(keys, l)
	}

	for _, p := range paths {
		if len(p) > 0 {
			valid = true
		}
	}
	if !valid {
		PrintError("Invalid file format.\nNo paths from start to end found")
	}
	//Sort paths based on size
	for _, p := range paths {
		sort.Slice(p, func(i, j int) bool {
			return len(p[i]) < len(p[j])
		})
	}

	validPaths := make([][]Path, 0)
	shortest := Path{Rooms: make([]string, 0), Ants: 0}
	for _, p := range paths {
		for _, v := range p {
			if len(shortest.Rooms) < 1 || len(v) < len(shortest.Rooms) {
				shortest.Rooms = v
			}
			validPaths = append(validPaths, narrowDown(Path{Rooms: v, Ants: 0}, paths, keys, startRoom, endRoom))
		}
	}
	validPaths = append(validPaths, []Path{shortest})

	for i := 0; i < len(validPaths); i++ {
		for j := 0; j < len(validPaths[i]); j++ {
			for k := 0; k < len(validPaths[i]); k++ {
				if j != k && validPaths[i][j].ToString() == validPaths[i][k].ToString() {
					validPaths[i] = append(validPaths[i][:j], validPaths[i][k+1:]...)
					continue
				}
			}
		}
	}

	sort.Slice(validPaths, func(i, j int) bool {
		return len(validPaths[i]) < len(validPaths[j])
	})

	return paths, validPaths
}

// Function to find all the paths that are not the shortest path even if the dont have an end room
func findAllPaths2(paths map[string][][]string, root string, startRoom string, endRoom string, links map[string][]string, path []string, visited []string) {
	for _, v := range links[startRoom] {
		if !Contains(visited, v) {
			visited = append(visited, v)
			path = append(path, v)
			findAllPaths2(paths, root, v, endRoom, links, path, visited)
			path = path[:len(path)-1]
			visited = visited[:len(visited)-1]
		}
	}
	if path[len(path)-1] == endRoom {
		paths[root] = append(paths[root], fixAppend(path))
	}
}

// Function to narrows down paths that can be used
func narrowDown(firstPath Path, paths map[string][][]string, keys []string, start string, end string) []Path {
	valid := make([]Path, 0)
	valid = append(valid, firstPath)

	for _, k1 := range keys {
		//Has no paths
		if len(paths[k1]) < 1 {
			continue
		}
		//Loop trough paths for k1
		for _, p1 := range paths[k1] {
			currentPath := p1
			if len(currentPath) > 0 && !resemblesPathArray(valid, currentPath, start, end) {
				valid = append(valid, Path{Rooms: currentPath, Ants: 0})
			}
		}
	}

	return valid
}

// Function to check if a slice contains a specific element
func Contains(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

// Function to check if an array resemles another array
func resemblesArrays(a1 []string, a2 []string, start string, end string) bool {
	for _, v1 := range a1 {
		if v1 == start || v1 == end {
			continue
		}
		for _, v2 := range a2 {
			if v2 == start || v2 == end {
				continue
			}
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

// Function to check if an array resemles another array in the valid map
func resemblesPathArray(v []Path, a2 []string, start string, end string) bool {
	for _, a1 := range v {
		if resemblesArrays(a2, a1.Rooms, start, end) {
			return true
		}
	}
	return false
}

func fixAppend(a []string) []string {
	cpy := make([]string, len(a))

	copy(cpy, a)

	return cpy
}

/*
func moveAntsThroughPaths(farm Farm, paths []Path) {
    antNumber := 1
    totalAnts := farm.Ants

    // Initialize the ant count in the start room
    farm.Rooms[farm.Start].AntCount = totalAnts

    // Move ants until all ants reach the end room
    for farm.Rooms[farm.End].AntCount < totalAnts {
        // Move ants already on paths
        for _, path := range paths {
            for i := len(path.Rooms) - 1; i > 0; i-- {
                currentRoom := path.Rooms[i]
                previousRoom := path.Rooms[i-1]

                if farm.Rooms[previousRoom].AntCount > 0 {
                    // Move ant to the current room
                    farm.Rooms[currentRoom].AntCount++
                    farm.Rooms[previousRoom].AntCount--

                    // Print the ant's movement
                    printAntMovement(farm.Rooms[currentRoom].Ant, currentRoom)
                }
            }
        }

        // Move new ants from the start room to the paths
        for _, path := range paths {
            if farm.Rooms[farm.Start].AntCount > 0 {
                nextRoom := path.Rooms[1]

                // Move ant to the next room on the path
                farm.Rooms[nextRoom].AntCount++
                farm.Rooms[farm.Start].AntCount--

                // Assign the ant number to the room
                farm.Rooms[nextRoom].Ant = antNumber
                antNumber++

                // Print the ant's movement
                printAntMovement(farm.Rooms[nextRoom].Ant, nextRoom)
            }
        }

        // Print a new line after each round of movements
        fmt.Println()
    }
}
*/
