package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"sync"

	"github.com/chasestarr/fsdft"
)

type worker struct {
	cmd  string
	args []string
}

func (w worker) run(wg *sync.WaitGroup) {
	fmt.Println("in run func")
	if err := exec.Command(w.cmd, w.args...).Run(); err != nil {
		fmt.Println("oh no!")
		log.Fatal(err)
		defer wg.Done()
	}
	fmt.Println("merged file at:", w.args[len(w.args)-1])
	defer wg.Done()
}

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
		if file.Name() == os.Args[2] {
			continue
		}
		if isPdf(file) {
			path := root + "/" + file.Name()
			pdfs = append(pdfs, path)
		}
	}
	return pdfs
}

func collect(root string, file os.FileInfo, c chan worker) {
	path := root + "/" + file.Name()
	if file.IsDir() && isLeafDir(path) {
		files := readDir(path)
		pdfs := getPdfs(path, files)

		if len(pdfs) >= 2 {
			cmd := "java"
			args := []string{"-jar", "./jar/pdfbox.jar", "PDFMerger"}
			args = append(args, pdfs...)

			output := path + "/" + os.Args[2]
			args = append(args, output)

			c <- worker{cmd: cmd, args: args}
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

// /Users/chasestarr/Dropbox/hr/job-search
func main() {
	root := os.Args[1]
	c := make(chan worker)
	go func() {
		fsdft.DFT(root, func(root string, file os.FileInfo) {
			collect(root, file, c)
		})
		defer close(c)
	}()

	var wg sync.WaitGroup
	for w := range c {
		wg.Add(1)
		go w.run(&wg)
	}

	wg.Wait()
}
