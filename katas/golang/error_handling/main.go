package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var fileError error

func main(){
    file, fileError := os.Open("data.csv")
    defer file.Close()

    if fileError != nil {
        log.Fatal(fileError)
    }
    
    reader := bufio.NewReader(file)
    for {
        line, fileError := reader.ReadString('\n')
        if fileError == io.EOF {
            break
        }
        if fileError != nil {
            log.Output(1, fileError.Error())
            break
        }
    
        fmt.Print(line)
    }

}
