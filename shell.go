package main

import ("fmt"
		"bufio"
		"os"
		"strings")

func main() {

	scanner := bufio.NewScanner(os.Stdin)

 	fmt.Print("Enter Command? ")

	scanner.Scan()
	 
	fmt.Print( scanner.Text(), "\n" )

	commandArray := strings.Fields(scanner.Text())

	switch [0]commandArray {
	case "ls":
		ls()
	}


	for _, v:= range commandArray{
		fmt.Println(v)
	}
	bye()




}