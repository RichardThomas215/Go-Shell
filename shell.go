package main

import ("fmt"
		"bufio"
		"os"
		"strings")

func main() {


	for {

	
	scanner := bufio.NewScanner(os.Stdin)

 	fmt.Print("Go Shell> ")

	scanner.Scan() 

	commandArray := strings.Fields(scanner.Text())
	
	if commandArray[0] == "ls" {
		
		ls( commandArray )

	} else if commandArray[0] == "end" {

		break

	}else {

		fmt.Printf("Command '%s' not found \n", strings.Join(commandArray, " ") )
	}
		

}


}