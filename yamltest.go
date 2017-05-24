package main

import (
        "flag"
        "fmt"
        yaml "gopkg.in/yaml.v2"
        "io/ioutil"
        "os"
        "path/filepath"
        "strings"
        "regexp"
)

var input_dir = flag.String("d", ".", "input the absolute path,the default is the current path")

func main() {
        flag.Parse()
        var allerrnum int
        files, err := walkdir(*input_dir, ".yaml")
        if err != nil {
                fmt.Printf("filepath returned %v\n", err)
        }
        for _, fileName := range files {
                var errnum int
                errnum, err = yamlUnmasrshal(fileName)
                if errnum != 0 {
                        allerrnum += errnum
                        fmt.Printf("%s have %d error. \n", fileName, errnum)
                        fmt.Println("=====================================")
                } else {
                        fmt.Printf("%s is OK. \n", fileName)
                        fmt.Println("-------------------------------------")
                }
        }
        if allerrnum != 0 {
                fmt.Printf("all have %d errors. \n", allerrnum)
                os.Exit(1)
        }

}
func walkdir(dirPath string, suffix string) (files []string, err error) {
        files = make([]string, 0, 50)
        suffix = strings.ToUpper(suffix)
        err = filepath.Walk(dirPath, func(filename string, fi os.FileInfo, err error) error {
                if err != nil {
                        fmt.Printf("filepath.Walk() returned %v\n", err)
                }
                if fi.IsDir() {
                        return nil
                }
                if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
                        files = append(files, filename)
                }
                return nil
        })
        return files, err
}

func yamlUnmasrshal(fileName string) (errnum int, err error) {
        m := make(map[interface{}]interface{})

        contents, err := ioutil.ReadFile(fileName)
        if err != nil {
                fmt.Println("error")
                os.Exit(1)
        }
        //      if strings.HasPrefix(string(contents), "---") {
        //              fmt.Printf("%s is not an ansible yaml.", fileName)
        //              err=fmt.Printf("%s is not an kube yaml.", fileName)
        //      } else {//line := regexp.MustCompile(`[^ +]---`).FindAllString(string(contents), 1)
	line := regexp.MustCompile(` +?---`).FindAllString(string(contents), -1)
	//fmt.Println(len(line), line)
	if len(line) > 0 {
		fmt.Printf("%s have spaces before ---,total is %d.\n", fileName, len(line))
		fmt.Println("=====================================")
		os.Exit(2)
	}
        n := strings.Split(string(contents), "---")
        for _, yamlcon := range n {
                err = yaml.Unmarshal([]byte(yamlcon), &m)
                if err != nil {
                        fmt.Printf("errorun: %v \n", err)
                        fmt.Printf(yamlcon)
                        errnum = errnum + 1
                }
        }
        //      }
        return errnum, err
}
