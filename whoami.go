package main

import (
	"fmt"
	"os/user"
)

func whoami() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Current User
	fmt.Println(user.Username)

}
