package myIndex

import "testing"
import "inverseIndex2/myFile"
import "strconv"

var fileStrings []string

func initIndexMap() {
	fileStrings = nil
	indexMap = make(map[string]map[string]int)
	files = nil
	count = 0
	fileStrings = append(fileStrings, "мама папа рыба хвост папа мама рыба мама")
	fileStrings = append(fileStrings, "дом деревня дом крыша корова мама дом java деревня")
	fileStrings = append(fileStrings, "go java cordova nodejs хвост java")

	for i, fileString := range fileStrings {
		file := myFile.NewMyFile("name"+strconv.Itoa(i), []byte(fileString))
		AddFile(file)
	}
}
func TestSearchWord(t *testing.T) {
	initIndexMap()

	words := []string{"мама", "деревня", "java", "cordova", "хвост"}
	expected := []int{4, 2, 3, 1, 2}
	var actual []int

	for _, word := range words {
		actual = append(actual, searchWord(word))
	}

	if !equalsIntSlice(expected, actual) {
		t.Errorf("%v is not eqal to expected %v", actual, expected)
	}

}

func TestSearch2(t *testing.T) {
	initIndexMap()

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
