package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"t2-invert-index-search-LoginovArkadiy/myFile"
	"t2-invert-index-search-LoginovArkadiy/myIndex"
)

type configuration struct {
	Dir      string `json:"dir"`
	Listener string `json:"listener"`
}

const bye = "bye"

var tmp *template.Template
var config = configuration{}

func main() {
	myIndex.Make()
	var err error
	fmt.Println("Hello Inverted Index")

	tmp, err = template.New("indexResponse.html").ParseFiles("indexResponse.html")

	if err != nil {
		fmt.Println("Что то с indexResponse.html")
		return
	}

	initFiles()

	http.HandleFunc("/", mainPage)

	fmt.Println("starting server at:", config.Listener)
	http.ListenAndServe(":"+config.Listener, nil)

}

//работа с браузером
func mainPage(w http.ResponseWriter, r *http.Request) {

	var files = []myFile.MyFile{myFile.MyFile{
		Name: "Здесь мог быть ответ",
		Sum:  0,
	}}
	if r.Method != http.MethodPost {
		tmp.Execute(w, files)
		return
	}

	phrase := r.FormValue("phrase")

	files = searchPhrase(phrase)

	tmp.Execute(w, files)
}

func initFiles() {

	confFile, _ := os.Open("config.json")
	defer confFile.Close()
	decoder := json.NewDecoder(confFile)

	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("Что то не так с config")
		panic(err)
	}

	files, err := ioutil.ReadDir(config.Dir)
	if err != nil {
		fmt.Println("Директория введена неверно dir = " + config.Dir + "listener = " + config.Listener)
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(files))

	for _, file := range files {
		data, error := ioutil.ReadFile(config.Dir + file.Name())
		if error != nil {
			fmt.Println("Ошибка в чтении файла" + file.Name())
			return
		}
		go func(name string, data []byte) {
			defer wg.Done()
			file := myFile.NewMyFile(name, data)
			myIndex.AddFile(file)
		}(file.Name(), data)

	}

	wg.Wait()
}

func searchPhrase(phrase string) []myFile.MyFile {
	if phrase == bye {
		return nil
	}

	words := strings.Fields(phrase)
	files := myIndex.Search2(words)

	return files
}
