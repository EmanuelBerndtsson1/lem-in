package graph

import (
	"fmt"
	"sort"
)

func MoveAnts(farm Farm, validPaths [][]Path) {
	//chosenPaths, maxPerRoom := choosePath(farm, validPaths)
	chosenPaths, _ := choosePath(farm, validPaths)

	PrintFileData(farm)

	//onlyOnePath(farm, chosenPaths)
	//moreThanOnePath(farm, maxPerRoom, chosenPaths)

	//moveAntsThroughPaths(farm, chosenPaths)

	moreThanOnePathX(farm, 1, chosenPaths)
}

// Chooses the best path to use for the number of ants specified
/* func choosePath(farm Farm, validPaths [][]Path) ([]Path, int) {
	var shortest [3]int
	for i, paths := range validPaths {

		totalLength := 0
		for _, path := range paths {
			totalLength += path.Length()
		}
		pathData := [3]int{i, (farm.Ants + totalLength) / len(paths), (farm.Ants + totalLength) % len(paths)}

		if i == 0 || pathData[1] < shortest[1] || (pathData[1] == shortest[1] && pathData[2] < shortest[2]) {
			shortest = pathData
		}

	}
	chosenPaths := validPaths[shortest[0]]
	sort.Slice(chosenPaths, func(i, j int) bool {
		return chosenPaths[i].Length() < chosenPaths[j].Length()
	})

	return chosenPaths, shortest[1]
} */

func choosePath(farm Farm, validPaths [][]Path) ([]Path, int) {
	var shortestGroupLenght int = 10000000 // Find something better than this :)
	var shortestGroupIndex int

	for i, paths := range validPaths {

		nOfAnts := farm.Ants

		for nOfAnts > 0 {
			shortestNow := 0
			// find index of best path for this ant
			for j, p := range paths {
				if len(p.Rooms)+p.Ants < len(paths[shortestNow].Rooms)+paths[shortestNow].Ants {
					shortestNow = j
				}
			}
			paths[shortestNow].Ants++
			nOfAnts--
		}

		// Establish a number to compare to
		longestNow := 0
		for j, p := range paths {
			if p.Ants != 0 {
				longestNow = j
				break
			}
		}

		// What the index of the longest taking path in this group is
		for j, p := range paths {
			if (len(p.Rooms)+p.Ants > len(paths[longestNow].Rooms)+paths[longestNow].Ants) && p.Ants != 0 {
				longestNow = j
			}
		}

		// compare to currently stored best group, and update if needed
		if len(paths[longestNow].Rooms)+paths[longestNow].Ants < shortestGroupLenght {
			shortestGroupLenght = len(paths[longestNow].Rooms) + paths[longestNow].Ants
			shortestGroupIndex = i
		}
	}

	chosenPaths := validPaths[shortestGroupIndex]

	sort.Slice(chosenPaths, func(i, j int) bool {
		return chosenPaths[i].Length() < chosenPaths[j].Length()
	})

	return chosenPaths, 1
}

// Prints the ant's movement
func printAntMovement(ant int, room string) {
	fmt.Printf("L%d-%s ", ant, room)
}

// There is only on path
/* func onlyOnePath(farm Farm, chosenPaths []Path) {
	if len(chosenPaths) > 1 {
		return
	}

	for farm.Rooms[farm.End].AntCount < farm.Ants {
		// Move ants already on path
		for _, path := range chosenPaths {
			var ant int = 0
			for i := len(path.Rooms) - 1; i > 0; i-- {
				room := path.Rooms[i]
				if room == farm.Start || room == farm.End {
					continue
				}
				if farm.Rooms[room].AntCount > 0 && farm.Rooms[room].Ant != ant {
					ant = farm.Rooms[room].Ant
					// Move ant to the next room
					farm.Rooms[path.Rooms[i+1]].AntCount++
					// Assign which ant is in the room
					farm.Rooms[path.Rooms[i+1]].Ant = ant
					// Decrement the number of ants in the room
					farm.Rooms[room].AntCount--
					// Print the ant's movement
					printAntMovement(ant, path.Rooms[i+1])
				}
			}
		}
		if farm.Rooms[farm.Start].AntCount > 0 {
			// Add ant to next room
			farm.Rooms[chosenPaths[0].Rooms[1]].AntCount++
			// Assign which ant is in the room
			farm.Rooms[chosenPaths[0].Rooms[1]].Ant = farm.Ants - farm.Rooms[farm.Start].AntCount + 1
			// Print the ant's movement
			printAntMovement(farm.Rooms[chosenPaths[0].Rooms[1]].Ant, chosenPaths[0].Rooms[1])
			// Increment the number of ants in the path
			chosenPaths[0].Ants++
			// Decrement the number of ants in the start room
			farm.Rooms[farm.Start].AntCount--
		}
		fmt.Println()
	}
} */

// There are multiple paths
/* func moreThanOnePath(farm Farm, maxPerRoom int, chosenPaths []Path) {
	if len(chosenPaths) < 2 {
		return
	}

	// Move the first ants
	for i, path := range chosenPaths {
		// Add ant to next room
		farm.Rooms[path.Rooms[1]].AntCount++
		// Assign which ant is in the room
		farm.Rooms[path.Rooms[1]].Ant = farm.Ants - farm.Rooms[farm.Start].AntCount + 1
		// Increment the number of ants in the path
		chosenPaths[i].Ants++
		// Decrement the number of ants in the start room
		farm.Rooms[farm.Start].AntCount--
		// Print the ant's movement
		printAntMovement(farm.Rooms[path.Rooms[1]].Ant, path.Rooms[1])
	}
	fmt.Println()

	for farm.Rooms[farm.End].AntCount < farm.Ants {
		// Move ants already on path
		for _, path := range chosenPaths {
			var ant int = 0
			for i := len(path.Rooms) - 1; i > 0; i-- {
				room := path.Rooms[i]
				if room == farm.Start || room == farm.End {
					continue
				}
				if farm.Rooms[room].AntCount > 0 && farm.Rooms[room].Ant != ant {
					ant = farm.Rooms[room].Ant
					// Move ant to the next room
					farm.Rooms[path.Rooms[i+1]].AntCount++
					// Assign which ant is in the room
					farm.Rooms[path.Rooms[i+1]].Ant = farm.Rooms[room].Ant
					// Decrement the number of ants in the previous room
					farm.Rooms[room].AntCount--
					// Print the ant's movement
					printAntMovement(farm.Rooms[path.Rooms[i+1]].Ant, path.Rooms[i+1])
				}
			}
		}
		// Move ants from start to path
		for i, path := range chosenPaths {
			pathLength := path.Length() + path.Ants
			if pathLength >= maxPerRoom && chosenPaths[0].Rooms[0] == farm.Start && chosenPaths[0].Rooms[1] == farm.End {
				continue
			}
			if farm.Rooms[farm.Start].AntCount > 0 {
				// Add ant to next room
				farm.Rooms[path.Rooms[1]].AntCount++
				// Assign which ant is in the room
				farm.Rooms[path.Rooms[1]].Ant = farm.Ants - farm.Rooms[farm.Start].AntCount + 1
				// Increment the number of ants in the path
				chosenPaths[i].Ants++
				// Decrement the number of ants in the start room
				farm.Rooms[farm.Start].AntCount--
				// Print the ant's movement
				printAntMovement(farm.Rooms[path.Rooms[1]].Ant, path.Rooms[1])
			}
		}
		fmt.Println()
	}
} */

/* func moveAntsThroughPaths(farm Farm, paths []Path) {
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
} */

func moreThanOnePathX(farm Farm, maxPerRoom int, chosenPaths []Path) {
	totalAnts := farm.Ants

	// Move ants until all ants reach the end room
	for farm.Rooms[farm.End].AntCount < totalAnts {

		for i, p := range chosenPaths {
			moved := 0

			for j := len(p.Rooms) - 1; j > 0; j-- {
				currRoomName := p.Rooms[j]
				prevRoomName := p.Rooms[j-1]

				if farm.Rooms[prevRoomName].AntCount > 0 {
					farm.Rooms[prevRoomName].AntCount--
					farm.Rooms[currRoomName].AntCount++
					//fmt.Print(prevRoomName, "-", currRoomName, " | ")

					fmt.Print((i + moved + 1), "-", currRoomName, " | ")

					moved++
				}

				if moved >= p.Ants {
					break
				}
			}

		}

		fmt.Println()

	}

}
