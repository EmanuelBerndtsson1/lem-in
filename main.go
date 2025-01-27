package main

import (
	g "lem-in/graph"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		g.PrintError("invalid number of arguments\nUsage: go run . <filename>")
	}
	fileContent, err := os.ReadFile(os.Args[1])
	if err != nil {
		g.PrintError(" could not read file " + os.Args[1])
	}
	farm := g.CreateFarm(string(fileContent))
	g.ReadAntData(&farm)
	g.ReadRoomData(&farm)
	farm.Rooms[farm.Start].AntCount = farm.Ants
	_, validPaths := g.FindAllPaths(farm.Start, farm.End, farm.Links)
	g.MoveAnts(farm, validPaths)
}
