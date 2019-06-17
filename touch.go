package main

import "os"

func touch(newFile string) {

	os.Create(newFile)

}
