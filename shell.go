package main

import ("fmt"
		"bufio"
		"os")

func main() {

	  scanner := bufio.NewScanner(os.Stdin)

 	 fmt.Print("Enter Command? ")

	 scanner.Scan()
	 
	fmt.Print( scanner.Text(), "\n" )

	bye()




}