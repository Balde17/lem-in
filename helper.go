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
		if countainSpace(scanner.Text()) && boolean == false {
			tampon := strings.Split(scanner.Text(), " ")
			roomtampon := Room{}
			if len(tampon) == 3 {
				roomtampon.name = tampon[0]
				roomtampon.coord_x, _ = strconv.Atoi(tampon[1])
				roomtampon.coord_y, _ = strconv.Atoi(tampon[2])

			}
			data.rooms = append(data.rooms, roomtampon)
		}
		if countainTiret(scanner.Text()) && boolean == false {
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

func deleteDuplicatePaths(allPaths [][]string, startRoom, endRoom string) ([][]string, error) {
	answer := [][]string{}
	if len(allPaths) != 0 {

		temp := allPaths[0]
		for j := 0; j < len(allPaths); j++ {
			for i := 0; i < len(allPaths); i++ {
				fmt.Println(intersectionPaths(temp, allPaths[i]))
				if intersectionPaths(temp, allPaths[i]) {

					answer = append(answer, temp)
					temp = allPaths[i]
				}
			}
		}
	}

	return answer, nil
}

func intersectionPaths(temp, path []string) bool {

	for i := 1; i < len(temp)-1; i++ {
		for j := 1; j < len(path)-1; j++ {
			if path[j] == temp[i] && j != i {
				return true
			}
		}
	}
	return false
}
