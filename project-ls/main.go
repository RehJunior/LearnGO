package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Your solution goes here. Good luck!
	files := listFiles("testdata")
	fmt.Println(strings.Join(files, "-"))
}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := os.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		result := fmt.Sprintf("%v", f.Name())
		dirs = append(dirs, result )
	}
	

	return dirs
}
