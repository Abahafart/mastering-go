package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument")
		return
	}
	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		fileInfo, error := os.Stat(fullPath)
		if error != nil {
			continue
		}
		mode := fileInfo.Mode()
		if !mode.IsRegular() {
			continue
		}
		if mode&0111 != 0 {
			fmt.Println(fullPath)
			return
		}
	}
}
