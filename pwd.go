
package main

import(
	"os" 
	"fmt"
	"log"
  )
  
  func pwd() {
	dir, err := os.Getwd()
	  if err != nil {
		  log.Fatal(err)
	  }
	fmt.Println(dir)
  }