package main

import (
	"fmt"
	"io/ioutil"
)

func cat(fileName string) {

	data, _ := ioutil.ReadFile(fileName)

	fmt.Println(string(data))

}
