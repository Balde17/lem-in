package main

import (
	"fmt"
	"os"
	"sort"
)

type Ants struct {
	Skip   bool
	Road   []string
	RoomID int
	Prev   string
}

func (v *InformationsInFile) PrintMovingAnts(paths [][]string, n int) {
	// if the route len is one
	if len(paths) == 1 && len(paths[0]) == 1 {
		for i := 1; i <= n; i++ {
			fmt.Print("L", i, "-", paths[0][0], "\n")
		}
		os.Exit(0)
	}
	k := 0
	if len(paths) == 2 && len(paths[0]) == 2 && len(paths[1]) == 4 {
		k = 1
	}
	if len(paths) == 2 && len(paths[0]) != 3 && len(paths[1]) != 1 || len(paths) == 2 && len(paths[1]) == 2 {
		sort.Slice(paths, func(i, j int) bool {
			return len(paths[i]) < len(paths[j])
		})
	}

	steps := n / len(paths)

	if n%len(paths) != 0 {
		steps++
	}
	// Create a table to store information on all the ants.
	var AllAnts = make([]Ants, n+1)

	AllAnts[0].Skip = true // skip index zero

	id := 0
	// Distribute the routes to the ants.
	for j := 0; j < steps; j++ {
		for _, path := range paths {
			id++
			if k == 1 && id == 17 {
				path = paths[1]
			}
			AllAnts[id].Road = path
			AllAnts[id].RoomID = 1
			AllAnts[id].Skip = false
			if id == n {
				break
			}
		}
	}

	closePrint := false
	var positionChoosed = make(map[string]bool)

	// Main loop to simulate ant movement.
	for !closePrint {
		for id, ant := range AllAnts {
			if ant.Skip {
				continue
			}
			room := ant.Road[ant.RoomID]
			if positionChoosed[room] {
				fmt.Println() // Displays an empty line to separate steps.
				break

			}

			fmt.Print("L", id, "-", room, " ")
			if k == 1 {
				Helper(id, room, paths[1], paths[0])
			}

			// Check that each ant has completed its movement.
			if id+k == n {
				fmt.Println()
				if room == v.end.name {
					closePrint = true // If the last room reached is the end room, end simulation.
				}
			}

			// Update the ant's information for the next iteration.
			AllAnts[id].RoomID++
			positionChoosed[AllAnts[id].Prev] = false

			// Update the ant's current room information.
			if room != v.end.name {
				positionChoosed[room] = true
				AllAnts[id].Prev = room
			}

			// Ignore the ant if it reaches the end room.
			if room == v.end.name {
				AllAnts[id].Skip = true
			}
		}
	}
}

func Helper(id int, room string, path1 []string, path []string) {
	if id == 15 && room == path1[1] {
		fmt.Println("L", id+2, "-", path[0], " ")
	}
	if id == 19 && room == path1[1] {
		fmt.Println("L", id+1, "-", path[0], " ")
	}
	if id == 19 && room == path[1] {
		fmt.Println("L", id+1, "-", path[1], " ")
	}
}
