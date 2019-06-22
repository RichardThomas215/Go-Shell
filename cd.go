package main

import (
	"os"
)

func changeDirectory(command []string) {

	if len(command) == 1 {

	} else {
		if command[1] == "." {

		} else {
			//create new directory path
			newDirectory := command[1]

			//	fmt.Println(os.Getenv("HOME"))
			os.Chdir(newDirectory)

		}

	}

}
