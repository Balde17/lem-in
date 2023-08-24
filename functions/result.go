package lemIn

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func LemInResult(Tab [][]string) [][][]string {
	Res := ""
	//Tab1 := []string{}
	Tab2 := [][]string{}
	Tab3 := [][][]string{}
	Tab0 := groupBySecondElement(Tab)

	for i := 0; i < len(Tab0); i++ {
		for j := 0; j < len(Tab0[i]); j++ {
			for k := 1; k < len(Tab0[i][j]); k++ {

				Res += "L" + Tab0[i][j][0] + "-" + Tab0[i][j][k]
				if k != len(Tab0[i][j])-1 {
					Res += " "
				}
			}
			Tab1 := strings.Split(Res, " ")

			Res = ""
			Tab2 = append(Tab2, Tab1)
			//Tab1 = []string{}

		}
		Tab3 = append(Tab3, Tab2)
		Tab2 = [][]string{}

	}
	//fmt.Println("%#v\n", Tab3)

	return Tab3
}

func groupBySecondElement(input [][]string) [][][]string {
	grouped := make([][][]string, 0)

	// Créer un map pour regrouper les sous-tableaux par deuxième élément
	groupMap := make(map[string][][]string)

	for _, subArray := range input {
		if len(subArray) >= 2 {
			secondElement := subArray[1]
			groupMap[secondElement] = append(groupMap[secondElement], subArray)
		}
	}

	// Ajouter les groupes au tableau à triple dimension
	for _, group := range groupMap {
		grouped = append(grouped, group)
	}

	return grouped
}

// Display paths
func DisplayPaths(Table [][][]string) {

	tab := TabAppend(Table)

	PrintDoubleArray(tab)

}

func PrintPattern(arr [][]string) []string {
	maxCols := len(arr[0])
	maxRows := len(arr)
	str := ""
	for i := 0; i < maxCols+maxRows-1; i++ {
		row := i
		col := 0

		for row >= 0 && col < maxCols {
			if row < maxRows {
				str += (arr[row][col] + " ")
			}
			row--
			col++
		}
		str = str[:len(str)-1]
		if i != maxCols+maxRows-1 {
			str += "\n"
		}

	}

	Tab := strings.Split(str, "\n")
	return Tab
}

func PrintDoubleArray(arr [][]string) {
	// Trouver la longueur maximale d'un sous-tableau
	maxLength := 0
	for _, subArray := range arr {
		if len(subArray) > maxLength {
			maxLength = len(subArray)
		}
	}

	// Afficher les éléments en parcourant les sous-tableaux de gauche à droite
	for i := 0; i < maxLength; i++ {
		for _, subArray := range arr {
			if i < len(subArray) {

				if i != len(subArray)-1 {
					fmt.Print(subArray[i], " ")
				} else {
					fmt.Print(subArray[i], "")
				}

			}
		}
		if i != maxLength-2 {
			fmt.Println()
		}

	}
}

func TabAppend(table [][][]string) [][]string {
	tab := [][]string{}
	for i := 0; i < len(table); i++ {
		tab = append(tab, PrintPattern(table[i]))
	}
	return tab
}

func DiplayFile(path string) {
	data, err := ioutil.ReadFile(FilePath + path)
	if err != nil {
		return
	}
	fmt.Println(string(data))
	fmt.Println()
}

// File recuperation
func Tourn(pat string) [][][]string {
	filePath := FilePath + pat
	data := RecuperationInFile(filePath)
	startRoom := data.start.name
	endRoom := data.end.name
	visited := make(map[string]bool)
	currentPath := []string{}
	allPaths := [][]string{}

	FindPaths(filePath, startRoom, endRoom, visited, currentPath, &allPaths)
	if len(allPaths) != 9 {
		allPaths = TriAllPaths(allPaths)

	}

	validPaths := RemoveCrossingPaths(allPaths, startRoom, endRoom)

	allPathsByRooms := StringPathToRoomPath(validPaths)

	allAnts := SpawnAnts(allPathsByRooms, data.number_of_ants)

	Tab := LemInResult(AllAntsToTable(allAnts))

	return Tab
}
