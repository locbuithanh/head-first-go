package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	info, err := os.Stat("file.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info.ModTime(), info.Size())
}
