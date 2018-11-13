package myIndex

import "t2-invert-index-search-LoginovArkadiy/myFile"
import "sort"
import "sync/atomic"
import "sync"
import "strings"

var count int64
var indexMap map[string]map[string]int
var names []string
var mutex = &sync.Mutex{}

type ByIndex []myFile.MyFile

func (s ByIndex) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByIndex) Less(i, j int) bool {
	f1 := s[i]
	f2 := s[j]

	if f1.Sum > f2.Sum {
		return true
	}
	if f1.Sum < f2.Sum {
		return false
	}

	return f1.Name < f2.Name
}

func (s ByIndex) Len() int {
	return len(names)
}

func Make() {
	indexMap = make(map[string]map[string]int)
}

func AddFile(file myFile.MyFile) {
	addToIndex(file.Words, file.Name)
	mutex.Lock()
	names = append(names, file.Name)
	mutex.Unlock()

	file.Words = nil
}
func refactorWord(word string) string {
	word = strings.Trim(word, ".,?!-\"")
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

}

func searchWord(word string, wg *sync.WaitGroup, files []myFile.MyFile) int {
	Summa := 0
	defer wg.Done()
	mutex.Lock()
	word = refactorWord(word)
	fileMap, ok := indexMap[strings.ToLower(word)]
	mutex.Unlock()
	if ok {
		for filename, sum := range fileMap {
			mutex.Lock()
			files[getIndexFileByName(filename, files)].Sum += sum
			Summa += sum
			mutex.Unlock()
		}

	}
	//fmt.Println("----------")
	atomic.AddInt64(&count, 1)
	return Summa
}

func Search2(words []string) []myFile.MyFile {
	var files []myFile.MyFile
	for _, name := range names {
		files = append(files, myFile.MyFile{
			Name: name,
			Sum:  0,
		})
	}
	var wg sync.WaitGroup
	wg.Add(len(words))

	for _, word := range words {
		go searchWord(word, &wg, files)
	}

	wg.Wait()
	sort.Sort(ByIndex(files))
	return files
}

func GetSize() int {
	return len(names)
}

func getIndexFileByName(name string, files []myFile.MyFile) int {
	for i := range files {
		if files[i].Name == name {
			return i
		}
	}
	return 0
}
