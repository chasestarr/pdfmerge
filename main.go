package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/chasestarr/fsdft"
)

func readDir(root string) []os.FileInfo {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func getPdfs(root string, files []os.FileInfo) []string {
	pdfs := []string{}
	for _, file := range files {
		if isPdf(file) {
			path := root + "/" + file.Name()
			pdfs = append(pdfs, path)
		}
	}
	return pdfs
}

func merge(root string, file os.FileInfo) {
	path := root + "/" + file.Name()
	if file.IsDir() && isLeafDir(path) {
		files := readDir(path)
		pdfs := getPdfs(path, files)

		if len(pdfs) >= 2 {
			cmd := "java"
			args := []string{"-jar", os.Args[2], "PDFMerger"}
			args = append(args, pdfs...)

			output := path + "/" + os.Args[3]
			args = append(args, output)

			if err := exec.Command(cmd, args...).Run(); err != nil {
				fmt.Println("oh no!")
				log.Fatal(err)
			}
			fmt.Println("merged file at:", output)
		}
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

func checkInput() {
	if len(os.Args) < 4 {
		log.Fatal("must pass 3 arguments to 'pdfmerge', (root directory, path to pdfbox.jar, output file name)")
	}
	if _, err := ioutil.ReadDir(os.Args[1]); err != nil {
		log.Fatal("first argument to 'pdfmerge' is not a directory, ", err)
	}

	if _, err := ioutil.ReadFile(os.Args[2]); err != nil {
		log.Fatal("second argument to 'pdfmerge', path to pdfbox.jar file could not be resolved, ", err)
	}

	if len(os.Args[3]) < 5 {
		log.Fatal("third argument to 'pdfmerge' must be longer than 4 characters")
	}

	if os.Args[3][len(os.Args[3])-4:] != ".pdf" {
		log.Fatal("third argument to 'pdfmerge' must be a .pdf file")
	}
}

// /Users/chasestarr/Dropbox/hr/job-search
func main() {
	checkInput()
	root := os.Args[1]
	fsdft.DFT(root, merge)
}
