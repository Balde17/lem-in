package lemIn

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//get all file informations
func RecuperationInFile(filename string) InformationsInFile {
	var data InformationsInFile
	file, err := os.Open(filename)
	if err != nil {
		println(err)
		os.Exit(0)
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
