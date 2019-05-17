package main

import (
	"fmt"

	//	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"syscall"
	//"time"
)

func ls(command []string) {

	if len(command) <= 2 {

		if command[1] != "-al" {

			displayFileInfo(command[1])

		}

	} else {

		for i := 1; i < len(command); i++ {

			displayFileInfo(command[i])
		}
	}

}

func displayFileInfo(fileName string) {

	//layoutISO := "2019-04-15"
	//layoutUS := "January 2, 2006"

	info, _ := os.Stat(fileName)

	fmt.Print(info.Mode(), " ")

	nlink := uint64(0)
	if sys := info.Sys(); sys != nil {
		if stat, ok := sys.(*syscall.Stat_t); ok {
			nlink = uint64(stat.Nlink)
		}
	}

	fmt.Print(nlink, " ")

	getGroupandUserNames()
	//getFileSize(info)

	fmt.Print(info.Size(), " ")

	// fmt.Print(info.ModTime(), " ")
	t := info.ModTime()
	//parseTime, _ := time.Parse(layoutUS, t)
	fmt.Print(t.Format("Apr 15 21:01:05"), " ")

	fmt.Println(info.Name())

	//fmt.Println();

}

func getGroupandUserNames() {

	groupCharactheristics, err := user.LookupGroupId(strconv.Itoa(os.Getegid()))
	userCharacteristics, err := user.LookupId(strconv.Itoa(os.Geteuid()))
	if err != nil {

		fmt.Println(err)

	}

	fmt.Print(groupCharactheristics.Name + " " + userCharacteristics.Username + " ")

}

// func getFileSize(info FileInfo ){

// 	fmt.Println(info.Size());

// }
