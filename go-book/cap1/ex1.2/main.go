package main

import (
	"fmt"
	"os"
)

func main(){

	for index, args := range os.Args[1:] {
        fmt.Println("Index: ", index, " Value: ", args)
    }
    
}
