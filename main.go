package main

import (
	"fmt"
	lemIn "lem/functions"
	"os"
)

func main() {

	if len(os.Args) == 2 {
		if lemIn.IsLemInFileCorect(os.Args[1]) {
			err := lemIn.CheckDuplicateRooms(lemIn.FilePath + os.Args[1])
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			Tab := lemIn.Tourn(os.Args[1])
			lemIn.DiplayFile(os.Args[1])

			lemIn.DisplayPaths(Tab)
		} else {
			fmt.Println("ERROR: invalid data format")
			return
		}

	} else {
		fmt.Println("ERROR: invalid arguments,only 2 arguments!!!")
		return
	}

}
