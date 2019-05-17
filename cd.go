package main

import (
	"os"
)

func changeDirectory(Directory string) {

	Directory = "../" + Directory
	os.Chdir(Directory)

}
