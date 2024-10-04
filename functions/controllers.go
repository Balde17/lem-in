package lemIn

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Verify if the file format is correct
func IsLemInFileCorect(str string) bool {
	file, err := os.Open(FilePath + str)
	boolean := true
	if err != nil {
		boolean = false
	}
	StartEnd := []string{}
	Room := []string{}
	RoomLinks := []string{}
	Number := 0
	NumberOfAnts := ""
	
	fileScaner := bufio.NewScanner(file)
	count := -1
	TabLinkRooms := []string{}
	for fileScaner.Scan() {
		line := fileScaner.Text()

		if Number == 0 {
			NumberOfAnts = line
			Number++
		}

		if count == 0 || count == 1 {
			StartEnd = append(StartEnd, strings.Split(line, " ")[0])
			count = -1
		}
		if line == "##start" {
			count = 0
		} else if line == "##end" {
			count = 1
		}

		//if it is a link
		if strings.Count(line, "-") == 1 {
			TabLinkRooms = append(TabLinkRooms, strings.Split(line, "-")[0]+"*"+strings.Split(line, "-")[1])
			
			RoomLinks = append(RoomLinks, strings.Split(line, "-")[0])
			RoomLinks = append(RoomLinks, strings.Split(line, "-")[1])
			if strings.Split(line, "-")[0] == strings.Split(line, "-")[1] {
				boolean = false
			}
		} else if strings.Count(line, "-") > 1 {
			boolean = false
		}
		//if it is a room
		if len(strings.Split(line, " ")) == 3 {
			Room = append(Room, strings.Split(line, " ")[0])
			_, er1 := strconv.Atoi(strings.Split(line, " ")[1])
			_, er2 := strconv.Atoi(strings.Split(line, " ")[2])
			if er1 != nil || er2 != nil {
				boolean = false
			}
			if string(strings.Split(line, " ")[0][0]) == "#" || string(strings.Split(line, " ")[0][0]) == "L" {
				boolean = false
			}

		}

	}
	n, _ := strconv.Atoi(NumberOfAnts)
	if n < 1 {
		boolean = false
	}
	for _, el := range RoomLinks {
		if !IsContain(Room, el) {
			boolean = false
		}
	}
	if len(StartEnd) != 2 {
		boolean = false
	}
	if len(StartEnd) == 2 {
		if !IsContain(RoomLinks, StartEnd[0]) || !IsContain(RoomLinks, StartEnd[1]) {
			boolean = false
		}
	}
	count0 := 0
	for j := 0; j < len(TabLinkRooms); j++ {
		for i := 0; i < len(TabLinkRooms); i++ {
			if TabLinkRooms[j] == TabLinkRooms[i] || TabLinkRooms[j] == strings.Split(TabLinkRooms[i], "*")[1]+"*"+strings.Split(TabLinkRooms[i], "*")[0] {
				count0++
			}
		}

		if count0 > 1 {
			boolean = false
		}
		count0 = 0
	}

	return boolean
}

// Verify if str is contained in Tab
func IsContain(Tab []string, str string) bool {
	for _, el := range Tab {
		if str == el {
			return true
		}
	}
	return false

}

//Checking duplicate rooms or coordinates
func CheckDuplicateRooms(filePath string) error {
	roomNames := make(map[string]int)
	roomCoordinates := make(map[string]int)

	file, err := os.Open(filePath)
	if err != nil {
		return errors.New("ERROR: invalid data format")
	}
	defer file.Close()

	begin := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if begin {
			begin = false
		}
		before, after, found := Spliter(line, " ")
		if found {
			//voir si la pièce existe, si oui incrementer de 1 sinon commencer à compter à 1
			_, ok := roomNames[before]
			if ok {
				roomNames[before] += 1
			} else {
				roomNames[before] = 1
			}
			//voir si la pièce existe, si oui incrementer de 1 sinon commencer à compter à 1
			_, exist := roomCoordinates[after]
			if exist {
				roomCoordinates[after] += 1
			} else {
				roomCoordinates[after] = 1
			}

		}
	}
	for _, v := range roomNames {
		if v > 1 {

			return errors.New("ERROR: invalid data format")
		}
	}
	for _, c := range roomCoordinates {
		if c > 1 {
			return errors.New("ERROR: invalid data format")
		}
	}

	return nil
}

func Spliter(s, sep string) (string, string, bool) {
	ind := strings.Index(s, sep)
	if ind > 0 {
		left := s[:ind]
		right := s[ind+len(sep):]

		return left, right, true
	}

	return s, "", false
}
