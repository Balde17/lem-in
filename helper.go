package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RecuperationInFile(filename string) InformationsInFile {
	var data InformationsInFile
	file, err := os.Open(filename)
	if err != nil {
		println(err)
	}
	defer file.Close()
	var boolean bool
	counter := 0
	scanner := bufio.NewScanner(file)
	indexS, indexE := 0, 0
	for scanner.Scan() {

		if counter == 0 {
			num, _ := strconv.Atoi(scanner.Text())
			data.number_of_ants = num
		}
		if scanner.Text() == "##start" {
			indexS = counter
			boolean = true
		} else {
			boolean = false
		}
		if counter == indexS+1 {
			tampon := strings.Split(scanner.Text(), " ")
			if len(tampon) == 3 {
				data.start.name = tampon[0]
				data.start.coord_x, _ = strconv.Atoi(tampon[1])
				data.start.coord_y, _ = strconv.Atoi(tampon[2])

			}

		}

		if scanner.Text() == "##end" {
			indexE = counter
			boolean = true
		} else {
			boolean = false
		}
		if counter == indexE+1 {
			tampon := strings.Split(scanner.Text(), " ")
			if len(tampon) == 3 {
				data.end.name = tampon[0]
				data.end.coord_x, _ = strconv.Atoi(tampon[1])
				data.end.coord_y, _ = strconv.Atoi(tampon[2])

			}

		}
		if countainSpace(scanner.Text()) && !boolean {
			tampon := strings.Split(scanner.Text(), " ")
			roomtampon := Room{}
			if len(tampon) == 3 {
				roomtampon.name = tampon[0]
				roomtampon.coord_x, _ = strconv.Atoi(tampon[1])
				roomtampon.coord_y, _ = strconv.Atoi(tampon[2])

			}
			data.rooms = append(data.rooms, roomtampon)
		}
		if countainTiret(scanner.Text()) && !boolean {
			tampon := strings.Split(scanner.Text(), "-")
			linktampon := LinksInRooms{}
			if len(tampon) == 2 {
				linktampon.room1 = tampon[0]
				linktampon.room2 = tampon[1]
			}
			data.links = append(data.links, linktampon)

		}
		counter++
	}

	return data
}

func countainSpace(s string) bool {
	count := 0
	for _, let := range s {
		if let == ' ' {
			count++
		}
	}
	return count == 2
}

func countainTiret(s string) bool {
	count := 0
	for _, let := range s {
		if let == '-' {
			count++
		}
	}
	return count == 1

}
func findPaths(filename, currentRoom, endRoom string, visited map[string]bool, currentPath []string, allPaths *[][]string) {
	visited[currentRoom] = true
	currentPath = append(currentPath, currentRoom)

	if currentRoom == endRoom {
		*allPaths = append(*allPaths, append([]string{}, currentPath...))
	} else {

		for _, neighbor := range Association(currentRoom, filename) {
			if !visited[neighbor] {
				findPaths(filename, neighbor, endRoom, visited, currentPath, allPaths)
			}
		}
	}

	delete(visited, currentRoom)
	// currentPath = currentPath[:len(currentPath)-1]
}

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

func removeCrossingPaths(allPaths [][]string, startRoom, endRoom string) [][]string {
	// Créez une nouvelle liste vide pour stocker les chemins filtrés
	filteredPaths := make([][]string, 0)
	// Parcourez tous les chemins dans allPaths
	for _, path := range allPaths {
		
		// Si le chemin ne se croise pas avec les chemins déjà filtrés
		if !isCrossing(path, filteredPaths, startRoom, endRoom) {
			// Ajoutez ce chemin à la liste des chemins filtrés
			filteredPaths = append(filteredPaths, path)
		}
	}

	// Renvoyez la liste des chemins filtrés
	return filteredPaths
}

func isCrossing(path []string, existingPaths [][]string, startRoom, endRoom string) bool {
	// Parcourez tous les chemins existants
	for _, existingPath := range existingPaths {
		// Si le chemin actuel se croise avec l'un des chemins existants
		if isPathCrossing(path, existingPath, startRoom, endRoom) {
			// Renvoyez true pour indiquer qu'il y a un croisement
			fmt.Println("cossing on ",path)

			return true
		}
	}
	fmt.Println("No cross on ", path)
	// Si le chemin ne se croise avec aucun chemin existant
	return false
}

func isPathCrossing(pathA, pathB []string, startRoom, endRoom string) bool {
	//verifier si il nya pas de chevauchement
	// if pathA[0] == pathB[0] || pathA[len(pathA)-1] == pathB[len(pathB)-1] {
	// 	return true
	// }
	// Parcourez les salles du premier chemin
	for _, roomA := range pathA {
		if roomA == startRoom || roomA == endRoom{
			continue
		}
		// Parcourez les salles du deuxième chemin
		for _, roomB := range pathB {
			// Si une salle commune est trouvée, cela signifie que les chemins se croisent
			if roomA == roomB {
				return true
			}
		}
	}
	// Si aucune salle commune n'est trouvée, les chemins ne se croisent pas
	return false
}

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

func stringPathToRoomPath(allPaths [][]string) [][]Room{
	var pathByRooms []Room
	var allPathsByRooms [][] Room

	for _, path := range allPaths{
		for _, room := range path{
			element := Room{
				name: room,
				Visited: false,
			}
			pathByRooms = append(pathByRooms, element)
		}
		allPathsByRooms = append(allPathsByRooms, pathByRooms)
		pathByRooms = []Room{}
	}
	return allPathsByRooms
}


// type Ant struct {
// 	ID       int
// 	Position *Room
// }

// func simulateAntsOptimized(paths [][]*Room, antsCount int) {
// 	// Initialisation des fourmis
// 	ants := make([]Ant, antsCount)
// 	for i := 0; i < antsCount; i++ {
// 		ants[i] = Ant{
// 			ID:       i + 1,
// 			Position: paths[0][0], // Toutes les fourmis commencent à la première salle du premier chemin
// 		}
// 		ants[i].Position.Visited = true // Marquer la pièce comme visitée par une fourmi
// 	}

// 	// Simulation du déplacement des fourmis
// 	for _, path := range paths {
// 		// Boucle pour gérer chaque étape de la salle
// 		for step := 1; step < len(path); step++ {
// 			// Boucle pour chaque fourmi
// 			for i := range ants {
// 				if ants[i].Position == path[step-1] && ants[i].Position != path[step] {
// 					if ants[i].Position.Visited {
// 						// Libérer la pièce précédente
// 						ants[i].Position.Visited = false
// 					}
// 					if !path[step].Visited {
// 						// Si la pièce actuelle n'a pas déjà une fourmi, la marquer comme visitée
// 						ants[i].Position = path[step]
// 						ants[i].Position.Visited = true
// 					}
// 				}
// 			}
// 		}
// 		// Afficher l'état des fourmis à chaque étape
// 		printAntsState(ants)
// 	}
// }

// func printAntsState(ants []Ant) {
// 	// Afficher l'état de chaque fourmi
// 	for _, ant := range ants {
// 		fmt.Printf("Ant %d is in room %s\n", ant.ID, ant.Position.Name)
// 	}
// 	fmt.Println() // Ligne vide pour séparer les étapes
// }

