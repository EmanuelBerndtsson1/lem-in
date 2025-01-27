package graph

import (
	"fmt"
	"os"
)

func PrintAntFarm(farm Farm) {
	fmt.Println()
	s := farm.Rooms[farm.Start]
	e := farm.Rooms[farm.End]

	fmt.Print("Ants: ", farm.Ants, "\n")
	fmt.Print("Start: ", *s, "\n")
	fmt.Print("End: ", *e, "\n")
	fmt.Print("Rooms: ")
	for _, r := range farm.Rooms {
		fmt.Print(*r)
		fmt.Print(" ")
	}
	fmt.Println()
	fmt.Print("Links: ")
	fmt.Println(farm.Links)
	fmt.Println()
}

func PrintPaths(allPaths map[string][][]string) {
	fmt.Println("\t\t\tPaths")
	for k, p := range allPaths {
		fmt.Println(k + ": ")
		for _, a := range p {
			fmt.Print(a)
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintShortestPath(shortest []string) {
	fmt.Println("\t\t\tShortest")
	fmt.Println(shortest)
	fmt.Println()
}

func PrintValidPaths(validPaths [][]Path) {
	fmt.Println("\t\t\tValid")
	for _, p := range validPaths {
		fmt.Print(p, "\n")
	}
	fmt.Println()
}

func PrintFileData(farm Farm) {
	fmt.Println(farm.FileData)
	fmt.Println()
}

func PrintError(errorMsg string) {
	fmt.Println("ERROR: " + errorMsg)
	os.Exit(1)
}
