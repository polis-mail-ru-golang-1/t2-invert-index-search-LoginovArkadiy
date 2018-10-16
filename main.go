// inverseIndex2 project main.go
package main

import (
	"fmt"
	"inverseIndex2/myFile"
	"io/ioutil"
	"os"
	"sort"
)

const bye = "bye"

func main() {

	fmt.Println("Hello World!")
	var files []myFile.MyFile
	//var names []string
	//names = append(names, "aa.txt", "bb.txt", "cc.txt")
	names := os.Args

	for _, name := range names {
		data, error := ioutil.ReadFile(name)
		if error != nil {
			fmt.Println("Ошибка в чтении файла")
			return
		}
		files = append(files, myFile.NewMyFile(name, data))
	}
	var statement string

	fmt.Println("Если хотите завершить программу введитe '", bye, "'")
	fmt.Println("Введите поисковую фразу: ")
	fmt.Fscan(os.Stdin, &statement)

	for statement != bye {
		words := split(statement, ' ')

		for i := range files {
			files[i].Analyse(words)
		}
		myFile.Count = len(files)
		sort.Sort(myFile.ByIndex(files))

		for _, file := range files {
			fmt.Println(file.Name, file.Sum)
		}
		////////////////////////////////////////
		fmt.Println("Введите поисковую фразу: ")
		fmt.Fscan(os.Stdin, &statement)
	}
}

//разбиваем строчку на массив по символу
func split(s string, sign rune) []string {
	var slice []string
	s += string(sign)
	token := ""
	for _, ch := range s {
		if ch == sign {
			if len(token) > 0 {
				slice = append(slice, token)
				token = ""
			}
		} else {
			token += string(ch)
		}
	}
	return slice
}
