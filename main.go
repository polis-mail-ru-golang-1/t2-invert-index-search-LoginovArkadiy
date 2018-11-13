package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"t2-invert-index-search-LoginovArkadiy/myFile"
	"t2-invert-index-search-LoginovArkadiy/myIndex"
	"time"
)

const bye = "bye"

var indexMap map[string]map[string]int

func main() {
	myIndex.Make()
	fmt.Println("Hello Inverted Index")
	initFiles()

	http.HandleFunc("/", mainPage)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)

	//processing()
}

var loginFormTmpl = []byte(`
	<!DOCTYPE html>
<html>

<head>
	<style>
		.formIn {
			background: gray;
			color: white;
			border-radius: 5%;
			padding: 10px;
			min-width: 30%;
			position: absolute;
			top: 50%;
			left: 50%;
			margin-right: -50%;
			transform: translate(-50%, -50%)
		}

		.Vvedite {
			position: relative;
			width: 90%;
			top: 50%;
			left: 5%;
		}

		.input {
			position: relative;
			width: 90%;
			top: 50%;
			left: 5%;
			height: 300px;
		}
		.button{
			position: relative;
			width: 90%;
			top: 50%;
			left: 5%;

			
		}
	</style>

</head>

<body>
	<form  class="formIn" action="/" method="post">
		<div "Vvedite">Введите поисковую фразу:</div>
		<textarea autofocus class="input" type="submit" name="phrase"></textarea>
		<input  value="Отправить" type="submit" title="Отправить" class="button"/>
	</form>


</body>

</html>
`)

//работа с браузером
func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write(loginFormTmpl)
		return
	}

	phrase := r.FormValue("phrase")

	fmt.Fprintln(w, "you enter: ", phrase)
	time.Sleep(2 * time.Millisecond)
	fmt.Fprintln(w, "Результаты поиска: ", searchPhrase(phrase))
	myIndex.Clear()

}

func initFiles() {
	var names []string
	names = append(names, "noon", "hard", "time", "prisoners")
	//names := os.Args

	for i := range names {
		if i == 0 {
			continue
		}
		data, error := ioutil.ReadFile(names[i])
		if error != nil {
			fmt.Println("Ошибка в чтении файла")
			return
		}
		go func(name string, data []byte) {
			fmt.Println(name, "пошёл на анализ")
			file := myFile.NewMyFile(name, data)
			//time.Sleep(time.Millisecond)
			myIndex.AddFile(file)
			fmt.Println(file.Name, "Вернулся")
		}(names[i], data)

	}

	for myIndex.GetSize()+1 < len(names) {
		//fmt.Println(myIndex.GetSize())
		time.Sleep(10 * time.Millisecond)
	}
}

//работа с конслои
func processing() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------------------------------------")
	fmt.Println("Если хотите завершить программу введитe '" + bye + "'")
	fmt.Println("Введите поисковую фразу: ")
	statement := readLine(reader)
	for statement != bye {
		words := strings.Split(statement, " ")
		fmt.Println("------------------------------")

		for _, file := range myIndex.Search2(words) {
			fmt.Println(file.Name, file.Sum)
		}

		myIndex.Clear()
		////////////////////////////////////////
		fmt.Println("**************************")
		fmt.Println("Введите поисковую фразу: ")
		statement = readLine(reader)
	}

}

func searchPhrase(phrase string) string {
	if phrase == bye {
		return "GOODBYE"
	}
	words := strings.Fields(phrase)
	files := myIndex.Search2(words)
	s := "\n"
	for _, file := range files {
		fmt.Println(file.Name, file.Sum)
		s = s + " - " + file.Name + "; совпадений - " + strconv.Itoa(file.Sum) + "\n"
	}
	return s
}

func readLine(reader *bufio.Reader) string {
	statementBytes, _, _ := reader.ReadLine()
	return string(statementBytes)
}
