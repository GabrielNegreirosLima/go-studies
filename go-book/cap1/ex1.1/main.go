package main

import (
	"fmt"
	"os"
)

func main(){

	var line, token string

	for _, args := range os.Args {
        line += token + args
        token = " "
    }
    
    fmt.Println(line)
}
