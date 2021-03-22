package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var files []string
var ignoreFile = map[string]struct{}{
	".git":       {},
	".idea":      {},
	"artisan.go": {},
}
var newProjectName string
var wg sync.WaitGroup

func main() {
	flag.StringVar(&newProjectName, "name", "", "new project name")
	flag.Parse()
	if newProjectName == "" {
		panic("input new project name; -name zhangsan")
	}

	files = []string{}
	str, _ := os.Getwd()
	getFiles(str)

	for _, v := range files {
		wg.Add(1)
		go func(file string) {
			renamePath(file)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func renamePath(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	fp, err := os.OpenFile(file, os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	content = bytes.ReplaceAll(content, []byte(`"weather_mgr/`), []byte(`"`+newProjectName+`/`))
	content = bytes.ReplaceAll(content, []byte(`"base"`), []byte(`"`+newProjectName+`"`))
	content = bytes.ReplaceAll(content, []byte(`/ssd/base`), []byte(`/ssd/`+newProjectName))
	content = bytes.ReplaceAll(content, []byte(`module base`), []byte(`module `+newProjectName))
	content = bytes.ReplaceAll(content, []byte(`/go/src/withdraw`), []byte(`/go/src/`+newProjectName))

	_, err = fp.Write(content)
	if err != nil {
		panic(err)
	}
	fmt.Println(file)
}

func getFiles(basePath string) {
	dirs, _ := ioutil.ReadDir(basePath)
	for _, f := range dirs {
		if _, ok := ignoreFile[f.Name()]; ok {
			continue
		}
		fPath := basePath + "/" + f.Name()
		if !f.IsDir() {
			files = append(files, fPath)
		} else {
			getFiles(fPath)
		}
	}
}
