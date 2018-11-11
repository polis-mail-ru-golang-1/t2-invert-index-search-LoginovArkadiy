package myFile

import "testing"
import "strings"

func TestCreateIndex(t *testing.T) {
	var fileStrings []string
	fileStrings = append(fileStrings, "мама папа рыба хвост папа мама рыба мама")
	fileStrings = append(fileStrings, "дом деревня дом крыша корова дом деревня")
	fileStrings = append(fileStrings, "go java cordova nodejs java")

	for _, value := range fileStrings {
		file := MyFile{
			strfile: value,
			HashMap: make(map[string]int),
		}

		expected := make(map[string]int)
		for _, subStr := range strings.Split(value, " ") {
			subStr = strings.TrimSpace(subStr)
			if expected[subStr] == 0 {
				expected[subStr] = strings.Count(value, subStr)
			}
		}
		file.createIndex()

		actual := file.HashMap
		if !equals(expected, actual) {
			t.Errorf("%v is not eqal to expected %v", actual, expected)
		}

	}

}

func equals(hm1, hm2 map[string]int) bool {
	if len(hm1) != len(hm2) {
		return false
	}

	if (hm1 == nil) != (hm2 == nil) {
		return false
	}

	check := func(hm1, hm2 map[string]int) bool {
		for key, value := range hm1 {
			if value != hm2[key] {
				return false
			}
		}
		return true
	}

	if !check(hm1, hm2) || !check(hm1, hm2) {
		return false
	}
	return true

}
