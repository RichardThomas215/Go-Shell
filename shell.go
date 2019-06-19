package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {

		// new scanner
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Go Shell> ")

		scanner.Scan()

		// parse input into an array
		commandArray := strings.Fields(scanner.Text())

		if commandArray[0] == "ls" {

			ls(commandArray)

		} else if commandArray[0] == "pwd" {
			fmt.Println(pwd())

		} else if commandArray[0] == "zip" {

			zipUp(commandArray)

		} else if commandArray[0] == "cd" {

			changeDirectory(commandArray)

		} else if commandArray[0] == "mkdir" {

			mkdir(commandArray[1])

		} else if commandArray[0] == "touch" {

			touch(commandArray[1])

		} else if commandArray[0] == "rm" {

			rm(commandArray[1])

		} else if commandArray[0] == "end" {

			break

		} else {

			fmt.Printf("Command '%s' not found \n", strings.Join(commandArray, " "))

		}

	}

}
