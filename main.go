package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/chasestarr/fsdft"
)

func readDir(root string) []os.FileInfo {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func readFiles(root string, files []os.FileInfo) {
	for _, file := range files {
		if isPdf(file) {
			path := root + "/" + file.Name()
			pdf, _ := ioutil.ReadFile(path)
			fmt.Println(pdf)
		}
	}
}

func print(root string, file os.FileInfo) {
	path := root + "/" + file.Name()
	if file.IsDir() && isLeafDir(path) {
		files := readDir(path)
		readFiles(path, files)
	}
}

func isLeafDir(root string) bool {
	files := readDir(root)

	for _, file := range files {
		if file.Name() == ".DS_Store" {
			continue
		}

		if file.IsDir() {
			return false
		}
	}
	return true
}

func isPdf(file os.FileInfo) bool {
	if len(file.Name()) > 3 {
		ending := file.Name()[len(file.Name())-4:]
		if ending == ".pdf" {
			return true
		}
	}
	return false
}

// /Users/chasestarr/Dropbox/hr/job-search
func main() {
	root := os.Args[1]
	// c := make(chan string)
	fsdft.DFT(root, print)
}
