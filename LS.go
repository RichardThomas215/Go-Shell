package main

import (
	"fmt"
//	"io/ioutil"
	"os"
	"syscall"
)

func ls(command []string){

	//fmt.Println(len(command))
//	if len(command) <= 2{

//		if command[1] != "-al"{

			
			displayFileInfo(command[1])

//		}else{
			
//		}

//	}
	
}




func displayFileInfo(fileName string){

	info,_ := os.Stat(fileName)

	fmt.Print(info.Mode(), " ")

	nlink := uint64(0)
    if sys := info.Sys(); sys != nil {
        if stat, ok := sys.(*syscall.Stat_t);  ok {
            nlink = uint64(stat.Nlink) 
        }
	}
	
	fmt.Print(nlink, " ");

	getGroupandUserNames()
	fmt.Println(os.Getgid());
	

}

func getGroupandUserNames() {

	
}
	//fmt.Println(info.Size())
