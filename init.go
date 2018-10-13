// init
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	NewMyFile := func(name string, data []byte) myFile {
		mF := myFile{
			name:    name,
			count0:  0,
			sum:     0,
			strfile: string(data),
		}
		return mF
	}

	var files []myFile

	for _, name := range os.Args {
		data, error := ioutil.ReadFile(name)
		if error != nil {
			fmt.Println("Ошибка в чтении файла")
			return
		}
		files = append(files, NewMyFile(name, data))
	}

	var statement string
	fmt.Println("Введите поисковую фразу: ")
	fmt.Fscan(os.Stdin, &statement)
	work(statement, files)
}
