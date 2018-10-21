package myIndex

import "inverseIndex2/myFile"
import "sort"

//import "fmt"

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

func Search(words []string) []myFile.MyFile {
	for _, word := range words {
		fileMap, ok := indexMap[word]
		//fmt.Println("----------" + word)
		if ok {
			for filename, sum := range fileMap {
				//	fmt.Println(filename, sum)
				files[getIndexFileByName(filename, files)].Sum += sum
			}
		}
	}

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
