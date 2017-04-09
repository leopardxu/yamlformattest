package main

import (
	"flag"
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)


var input_dir=flag.String("d", ".", "input the absolute path,the default is the current path")

func main() {
	flag.Parse()
	listDirandtestyaml(*input_dir, "yaml")
}
func listDirandtestyaml(dirPath string, suffix string) {
	files := make([]string, 0, 40)
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("the error is %s",err) 
	}
	PathSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, dirPath+PathSep+fi.Name())
		}
	}
	for _, fileName := range files {
              var errnum int
	      errnum,err = yamlUnmasrshal(fileName)
		if errnum != 0 {
			fmt.Printf("%s have %d error. \n", fileName,errnum)
		} else {
			fmt.Printf("%s is OK. \n", fileName)
		}
	}
}

func yamlUnmasrshal(fileName string)(errnum int,err error) {
	m := make(map[interface{}]interface{})
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
	n := strings.Split(string(contents), "---")
	for _, yamlcon := range n {
		err = yaml.Unmarshal([]byte(yamlcon), &m)
		if err != nil {
			fmt.Printf("errorun: %v \n", err)
			fmt.Printf(yamlcon)
                        errnum = errnum +1
		}
	}
	return errnum,err
}
