package main

import "os"

func mkdir(name string) {

	path := pwd()

	path = path + name
	os.MkdirAll(name, 0771)

}
