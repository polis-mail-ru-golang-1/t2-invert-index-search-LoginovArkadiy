package myIndex

import "inverseIndex2/myFile"
import "sort"
import "sync/atomic"
import "time"
import "math/rand"
import "fmt"

var count int64
var indexMap map[string]map[string]int
var files []myFile.MyFile

func Make() {
	indexMap = make(map[string]map[string]int)
}

func AddFile(file myFile.MyFile) {
	files = append(files, file)
	addToIndex(file.HashMap, file.Name)
	file.HashMap = nil
}

func addToIndex(hm map[string]int, name string) {
	for word, value := range hm {
		_, ok := indexMap[word]
		if !ok {
			indexMap[word] = make(map[string]int)
		}
		indexMap[word][name] += value
	}
}

func searchWord(word string) int {
	Summa := 0
	fileMap, ok := indexMap[word]
	if ok {
		for filename, sum := range fileMap {
			files[getIndexFileByName(filename, files)].Sum += sum
			Summa += sum
			fmt.Println(filename, sum, word)
			amt := time.Duration(rand.Intn(100))
			time.Sleep(time.Millisecond * amt)
		}

	}
	fmt.Println("----------")
	atomic.AddInt64(&count, 1)
	return Summa
}

func Search2(words []string) []myFile.MyFile {
	count = 0

	for _, word := range words {
		go searchWord(word)
	}

	for count < int64(len(words)) {
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Все файлы проанализированы", count == int64(len(words)))
	myFile.Count = len(files)
	sort.Sort(myFile.ByIndex(files))
	return files
}

func Clear() {
	for i := range files {
		files[i].Sum = 0
	}
}
func getIndexFileByName(name string, files []myFile.MyFile) int {
	for i := range files {
		if files[i].Name == name {
			return i
		}
	}
	return -1
}
