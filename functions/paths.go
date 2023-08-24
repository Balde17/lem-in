package lemIn

//if the paths are crossing
func isPathCrossing(pathA, pathB []string, startRoom, endRoom string) bool {
	
	for _, roomA := range pathA {
		if roomA == startRoom || roomA == endRoom {
			continue
		}
		for _, roomB := range pathB {
			
			if roomA == roomB {
				return true
			}
		}
	}
	// if there is no crossing
	return false
}

//Sort all paths by length
func TriAllPaths(allPaths [][]string) [][]string {
	for i := 0; i < len(allPaths); i++ {
		for j := 0; j < len(allPaths); j++ {
			if len(allPaths[i]) < len(allPaths[j]) {
				temp := allPaths[i]
				allPaths[i] = allPaths[j]
				allPaths[j] = temp
			}
		}
	}
	return allPaths
}

// convert String Path To Room Path
func StringPathToRoomPath(allPaths [][]string) [][]Room {
	var pathByRooms []Room
	var allPathsByRooms [][]Room

	for _, path := range allPaths {
		for _, room := range path {
			element := Room{
				name:    room,
				Visited: false,
			}
			pathByRooms = append(pathByRooms, element)
		}
		allPathsByRooms = append(allPathsByRooms, pathByRooms)
		pathByRooms = []Room{}
	}
	return allPathsByRooms
}

//Remove crossing paths after sort
func RemoveCrossingPaths(allPaths [][]string, startRoom, endRoom string) [][]string {
	
	filteredPaths := make([][]string, 0)
	for _, path := range allPaths {

		if !isCrossing(path, filteredPaths, startRoom, endRoom) {
			
			filteredPaths = append(filteredPaths, path)
		}
	}

	return filteredPaths
}

//Checking if the paths are cross
func isCrossing(path []string, existingPaths [][]string, startRoom, endRoom string) bool {
	
	for _, existingPath := range existingPaths {
		if isPathCrossing(path, existingPath, startRoom, endRoom) {
			return true
		}
	}
	
	return false
}

//search all paths
func FindPaths(filename, currentRoom, endRoom string, visited map[string]bool, currentPath []string, allPaths *[][]string) {
	visited[currentRoom] = true
	currentPath = append(currentPath, currentRoom)

	if currentRoom == endRoom {
		*allPaths = append(*allPaths, append([]string{}, currentPath...))
	} else {

		for _, neighbor := range Association(currentRoom, filename) {
			if !visited[neighbor] {
				FindPaths(filename, neighbor, endRoom, visited, currentPath, allPaths)
			}
		}
	}

	delete(visited, currentRoom)
	// currentPath = currentPath[:len(currentPath)-1]
}

//return all rooms these are connected to the current room
func Association(current string, filename string) []string {
	el := []string{}
	data := RecuperationInFile(filename)
	for _, name := range data.links {
		if name.room1 == current {
			el = append(el, name.room2)
		} else if name.room2 == current {
			el = append(el, name.room1)

		}
	}
	return el
}
