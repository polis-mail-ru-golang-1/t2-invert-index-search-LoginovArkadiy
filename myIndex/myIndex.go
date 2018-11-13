package myIndex

import "github.com/polis-mail-ru-golang-1/t2-invert-index-search-LoginovArkadiy/myFile"
import "sort"
import "sync/atomic"
import "time"
import _ "math/rand"
import "sync"
import "strings"

//import "fmt"

var count int64
var indexMap map[string]map[string]int
var files []myFile.MyFile
var mutex = &sync.Mutex{}

func Make() {
	indexMap = make(map[string]map[string]int)
}

func AddFile(file myFile.MyFile) {
	addToIndex(file.Words, file.Name)
	mutex.Lock()
	files = append(files, file)
	mutex.Unlock()

	file.Words = nil
}
func refactorWord(word string) string {
	badSigns := []rune{
		'"',
		',',
		'.',
		'!',
		'?',
	}

	for _, sign := range badSigns {
		word = strings.Trim(word, string(sign))
	}
	word = strings.ToLower(word)
	return word
}

func addToIndex(words []string, name string) {
	for _, word := range words {
		word = refactorWord(word)
		mutex.Lock()
		_, ok := indexMap[word]
		if !ok {
			indexMap[word] = make(map[string]int)
		}
		indexMap[word][name] += 1
		mutex.Unlock()

	}

	time.Sleep(time.Millisecond)
}

func searchWord(word string, wg *sync.WaitGroup) int {
	Summa := 0
	defer wg.Done()
	mutex.Lock()
	fileMap, ok := indexMap[strings.ToLower(word)]
	mutex.Unlock()
	if ok {
		for filename, sum := range fileMap {
			mutex.Lock()
			files[getIndexFileByName(filename, files)].Sum += sum
			Summa += sum
			//fmt.Println(filename, sum, word)
			time.Sleep(time.Millisecond)
			mutex.Unlock()
		}

	}
	//fmt.Println("----------")
	atomic.AddInt64(&count, 1)
	return Summa
}

func Search2(words []string) []myFile.MyFile {
	var wg sync.WaitGroup
	wg.Add(len(words))

	for _, word := range words {
		go searchWord(word, &wg)
	}

	wg.Wait()
	myFile.Count = len(files)
	sort.Sort(myFile.ByIndex(files))
	return files
}

func Clear() {
	for i := range files {
		files[i].Sum = 0
	}
}

func GetSize() int {
	return len(files)
}
func getIndexFileByName(name string, files []myFile.MyFile) int {
	for i := range files {
		if files[i].Name == name {
			return i
		}
	}
	return -1
}
