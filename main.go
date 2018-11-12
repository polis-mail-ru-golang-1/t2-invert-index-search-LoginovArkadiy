package main

import (
	"bufio"
	"fmt"
	"inverseIndex2/myFile"
	"inverseIndex2/myIndex"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
<html>
	<body>
	<form action="/" method="post">
		Введите поисковую фразу: <input type="text" placeholder ="Предложение"  name="phrase"> 
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
	words := strings.Split(phrase, " ")
	files := myIndex.Search2(words)
	s := "\n"
	for _, file := range files {
		fmt.Println(file.Name, file.Sum)
		s = s + file.Name + " " + strconv.Itoa(file.Sum) + "\n"
	}
	return s
}

func readLine(reader *bufio.Reader) string {
	statementBytes, _, _ := reader.ReadLine()
	return string(statementBytes)
}
