package lemIn

import (
	"fmt"
	"math"
	"reflect"
)

// Makes list of ants with own path, current room and id
func SpawnAnts(paths [][]Room, numberOfAnt int) []Ant {
	var result []Ant
	for i := 1; i < numberOfAnt+1; i++ {
		var antToAppend Ant

		antToAppend.Id = i
		antToAppend.CurrentRoom = paths[0][0]
		antToAppend.Path = getIdealPath(paths, result)

		result = append(result, antToAppend)

	}

	return result
}

//get the ideal path by (number of rooms + number of ants these affected on the path)
func getIdealPath(paths [][]Room, result []Ant) []Room {
	counter := 0
	mapPath := make(map[int][]Room)
	for _, path := range paths {
		counter = len(path)
		for _, ant := range result {
			if reflect.DeepEqual(ant.Path, path) {
				counter++
			}
		}
		mapPath[counter] = path
	}

	min := math.MaxInt32
	for number := range mapPath {
		if number < min {
			min = number
		}
	}
	return mapPath[min]
}

//put ants in table to print the result
func AllAntsToTable(allAnts []Ant) [][]string {
	Res := [][]string{}
	for _, ant := range allAnts {
		Tab := []string{}
		for i := 1; i < len(ant.Path); i++ {
			if i == 1 {

				Tab = append(Tab, fmt.Sprintf("%d", ant.Id))
			}
			if i == len(ant.Path)-1 {
				Tab = append(Tab, fmt.Sprintf("%v", ant.Path[i].name))
			} else {
				Tab = append(Tab, fmt.Sprintf("%v", ant.Path[i].name))
			}

		}
		Res = append(Res, Tab)

	}
	return Res
}
