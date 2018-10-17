// inverseIndex2 project main.go
package main

import (
	"bufio"
	"fmt"
	"inverseIndex2/myFile"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

const bye = "bye"

var indexMap map[string]map[string]int

func main() {
	indexMap = make(map[string]map[string]int)

	fmt.Println("Hello Inverted Index")

	var files []myFile.MyFile

	initFiles(&files)
	processing(files)
}

func initFiles(files *[]myFile.MyFile) {
	//var names []string
	//names = append(names, "aa.txt", "bb.txt", "cc.txt")
	names := os.Args

	for i := range names {
		data, error := ioutil.ReadFile(names[i])
		if error != nil {
			fmt.Println("Ошибка в чтении файла")
			return
		}
		file := myFile.NewMyFile(names[i], data)
		for word, value := range file.HashMap {
			_, ok := indexMap[word]
			if !ok {
				indexMap[word] = make(map[string]int)
			}
			indexMap[word][file.Name] += value
		}

		*files = append(*files, file)
	}
	/*
		for word, fileMap := range indexMap {
			fmt.Print(word + " - ")
			for name, sum := range fileMap {
				fmt.Print(name+": ", sum, " ")
			}
			fmt.Println()
		}*/

}

func processing(files []myFile.MyFile) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------------------------------------")
	fmt.Println("Если хотите завершить программу введитe '" + bye + "'")
	fmt.Println("Введите поисковую фразу: ")
	statement := readLine(reader)
	for statement != bye {
		words := strings.Split(statement, " ")

		for _, word := range words {
			fileMap, ok := indexMap[word]
			fmt.Println("----------" + word)
			if ok {
				for filename, sum := range fileMap {
					fmt.Println(filename, sum)
					files[getIndexFileByName(filename, files)].Sum = sum
				}
			}
		}
		fmt.Println("------------------------------")
		myFile.Count = len(files)
		sort.Sort(myFile.ByIndex(files))

		for _, file := range files {
			fmt.Println(file.Name, file.Sum)
			file.Sum = 0
		}
		////////////////////////////////////////
		fmt.Println("Введите поисковую фразу: ")
		statement = readLine(reader)
	}

}

func readLine(reader *bufio.Reader) string {
	statementBytes, _, _ := reader.ReadLine()
	return string(statementBytes)
}

func getIndexFileByName(name string, files []myFile.MyFile) int {
	for i := range files {
		if files[i].Name == name {
			return i
		}
	}
	return -1
}
