package myIndex

import "testing"
import "sync"
import "t2-invert-index-search-LoginovArkadiy/myFile"
import "strconv"
import "time"
import _ "fmt"

func initIndexMap(fileStrings []string) {
	var wg sync.WaitGroup
	wg.Add(len(fileStrings))
	indexMap = make(map[string]map[string]int)
	files = nil
	count = 0

	for i, fileString := range fileStrings {

		go func(name string, data []byte) {
			defer wg.Done()
			file := myFile.NewMyFile(name, data)
			AddFile(file)
			time.Sleep(time.Millisecond)
		}("name"+strconv.Itoa(i), []byte(fileString))
	}

	wg.Wait()
}

func TestCreateIndexMap(t *testing.T) {
	var fileStrings []string
	fileStrings = append(fileStrings, "go java go go go хвост java", "мама хвост хвост go мама go хвост java")
	initIndexMap(fileStrings)
	expected := make(map[string]map[string]int)
	expected["мама"] = make(map[string]int)
	expected["java"] = make(map[string]int)
	expected["хвост"] = make(map[string]int)
	expected["go"] = make(map[string]int)
	expected["мама"]["name1"] = 2
	expected["java"]["name0"] = 2
	expected["java"]["name1"] = 1
	expected["хвост"]["name0"] = 1
	expected["хвост"]["name1"] = 3
	expected["go"]["name0"] = 4
	expected["go"]["name1"] = 2

	for key, filesMap := range expected {
		for name, count := range filesMap {
			if indexMap[key][name] != count {
				t.Errorf("%v is not eqal to expected %v", indexMap, expected)
			}
		}
	}

}

//func TestSearchWord(t *testing.T) {

//	var fileStrings []string
//	fileStrings = append(fileStrings, "мама папа рыба хвост папа мама рыба мама")
//	fileStrings = append(fileStrings, "дом деревня дом крыша корова мама дом java деревня")
//	fileStrings = append(fileStrings, "go java cordova nodejs хвост java")

//	initIndexMap(fileStrings)

//	words := []string{"мама", "деревня", "java", "cordova", "хвост"}
//	expected := []int{4, 2, 3, 1, 2}
//	var actual []int

//	var wg sync.WaitGroup
//	wg.Add(len(words))
//	for _, word := range words {
//		actual = append(actual, searchWord(word, &wg))
//	}

//	if !equalsIntSlice(expected, actual) {
//		t.Errorf("%v is not eqal to expected %v", actual, expected)
//	}

//}

func TestSearch2(t *testing.T) {
	var fileStrings []string
	fileStrings = append(fileStrings, "мама папа рыба хвост папа мама рыба мама")
	fileStrings = append(fileStrings, "дом деревня дом крыша корова мама дом java деревня")
	fileStrings = append(fileStrings, "go java cordova nodejs хвост java")

	initIndexMap(fileStrings)

	words := []string{"мама", "деревня", "java", "cordova", "хвост", "go"}
	var expected []myFile.MyFile

	expected = append(expected, myFile.MyFile{
		Name: "name2",
		Sum:  5,
	}, myFile.MyFile{
		Name: "name0",
		Sum:  4,
	}, myFile.MyFile{
		Name: "name1",
		Sum:  4,
	})

	actual := Search2(words)

	if !equalsFileSlice(expected, actual) {
		t.Errorf("%v is not eqal to expected %v", actual, expected)
	}

}

func equalsFileSlice(slice1, slice2 []myFile.MyFile) bool {
	if len(slice1) != len(slice1) {
		return false
	}

	if (slice1 == nil) != (slice2 == nil) {
		return false
	}

	for i := range slice1 {
		if slice1[i].Sum != slice2[i].Sum || slice1[i].Name != slice2[i].Name {
			return false
		}
	}
	return true

}

func equalsIntSlice(slice1, slice2 []int) bool {
	if len(slice1) != len(slice1) {
		return false
	}

	if (slice1 == nil) != (slice2 == nil) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true

}
