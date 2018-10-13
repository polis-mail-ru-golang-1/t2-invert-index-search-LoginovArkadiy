// inverseindex project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type myFile struct {
	name    string
	strfile string
	sum     int
	count0  int
}

func (file *myFile) setName(name string) {
	file.name = name
}

func (file *myFile) setStrfile(strfile string) {
	file.strfile = strfile
}

func (file myFile) Strfile() string {
	return file.strfile
}

func (file *myFile) setData(sum, count0 int) {
	file.sum = sum
	file.count0 = count0
}

func (file myFile) Count0() int {
	return file.count0
}

func (file myFile) Sum() int {
	return file.sum
}

func (file myFile) Name() string {
	return file.name
}

func main() {

	createFile := func(name string, data []byte) myFile {
		mF := myFile{
			name:    name,
			count0:  0,
			sum:     0,
			strfile: string(data),
		}
		return mF
	}

	var files []myFile

	var countFiles int
	fmt.Println("Введите кол-во файлов")
	fmt.Fscan(os.Stdin, &countFiles)
	fmt.Println("Введите " + strconv.Itoa(countFiles) + " файлов без их расширения (по умолчанию '.txt')")
	for i := 0 + 1; i < countFiles+1; i++ {
		var name string
		fmt.Fscan(os.Stdin, &name)
		data, error := ioutil.ReadFile(name + ".txt")
		if error != nil {
			fmt.Println("Файла с таким именем не существует!")
			return
		}
		files = append(files, createFile(name, data))
	}

	work(getSlice("мама", "мячик", "папа"), files)

}

func work(words []string, files []myFile) {
	//считаем
	for i := 0; i < len(files); i++ {
		sum := 0
		count0 := 0
		for _, word := range words {
			x := strings.Count(files[i].Strfile(), word)
			if x == 0 {
				count0++
			}
			sum += x
		}
		files[i].setData(sum, count0)
	}

	//сортим
	for i := 1; i < len(files); i++ {
		for j := i; j > 0; j-- {
			if compare(files[j], files[j-1]) {
				files[j], files[j-1] = files[j-1], files[j]
			} else {
				break
			}
		}
	}

	//выводим
	fmt.Println("-----------------------------------------")
	for _, file := range files {
		fmt.Println(file.name + " " + strconv.Itoa(file.Sum()))
	}
}

//по массиву стрингов возвращаем слайс
func getSlice(a ...string) []string {
	var slice []string
	for _, val := range a {
		slice = append(slice, val)
	}
	return slice
}

//сравниваем два файла
func compare(f1, f2 myFile) bool {
	if f1.Count0() < f2.Count0() {
		return true
	}
	if f1.Count0() > f2.Count0() {
		return false
	}

	if f1.Sum() > f2.Sum() {
		return true
	}
	if f1.Sum() < f2.Sum() {
		return false
	}

	return f1.name < f2.name
}
