package main

import (
	"bufio"
	"fmt"
	"inverseIndex2/myFile"
	"inverseIndex2/myIndex"
	"io/ioutil"
	"os"
	"strings"
)

const bye = "bye"

var indexMap map[string]map[string]int

func main() {
	myIndex.Make()
	fmt.Println("Hello Inverted Index")
	initFiles()
	processing()
}

func initFiles() {
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

		myIndex.AddFile(file)

	}
}

func processing() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------------------------------------")
	fmt.Println("Если хотите завершить программу введитe '" + bye + "'")
	fmt.Println("Введите поисковую фразу: ")
	statement := readLine(reader)
	for statement != bye {
		words := strings.Split(statement, " ")
		fmt.Println("------------------------------")
		for _, file := range myIndex.Search(words) {
			fmt.Println(file.Name, file.Sum)
		}
		myIndex.Clear()
		////////////////////////////////////////
		fmt.Println("**************************")
		fmt.Println("Введите поисковую фразу: ")
		statement = readLine(reader)
	}

}

func readLine(reader *bufio.Reader) string {
	statementBytes, _, _ := reader.ReadLine()
	return string(statementBytes)
}
