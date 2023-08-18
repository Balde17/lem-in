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

	validPaths, _ := deleteDuplicatePaths(allPaths, startRoom, endRoom)
	fmt.Println(validPaths)

	for _, path := range allPaths {
		fmt.Println("Chemin trouvé:", path)
	}

}
