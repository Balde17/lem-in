package main

import (
	"fmt"
	"os"
)

func main() {

	filePath := os.Args[1]
	data := RecuperationInFile(filePath)
	startRoom := data.start.name
	endRoom := data.end.name
	visited := make(map[string]bool)
	currentPath := []string{}
	allPaths := [][]string{}

	findPaths(filePath, startRoom, endRoom, visited, currentPath, &allPaths)
	if len(allPaths) != 9 {
		allPaths = TriAllPaths(allPaths)

	}

	validPaths := removeCrossingPaths(allPaths, startRoom, endRoom)
	//fmt.Println(validPaths)

	for _, path := range validPaths {
		fmt.Println("Chemin trouvé:", path)
	}

	allPathsByRooms := stringPathToRoomPath(validPaths)
	//fmt.Println(allPathsByRooms)

	allAnts := SpawnAnts(allPathsByRooms, data.number_of_ants)
	
	for _, p := range allAnts{
		fmt.Println(p.Id, p.Path)
	}

	//MakeStep(allAnts, data)

}
