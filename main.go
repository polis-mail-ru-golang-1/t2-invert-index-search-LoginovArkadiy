package main

import (
	"fmt"
	"sort"
	"strings"
)

type myFile struct {
	name    string
	strfile string
	sum     int
	count0  int
}

type byIndex []myFile

func (s byIndex) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byIndex) Less(i, j int) bool {
	f1 := s[i]
	f2 := s[j]
	if f1.count0 < f2.count0 {
		return true
	}
	if f1.count0 > f2.count0 {
		return false
	}

	if f1.sum > f2.sum {
		return true
	}
	if f1.sum < f2.sum {
		return false
	}

	return f1.name < f2.name
}
func (s byIndex) Len() int {
	return 1
}

func work(statement string, files []myFile) {
	//считаем

	words := split(statement, ' ')

	for i := 0; i < len(files); i++ {
		sum := 0
		count0 := 0
		for _, word := range words {
			x := strings.Count(files[i].strfile, word)
			if x == 0 {
				count0++
			}
			sum += x
		}
		files[i].sum = sum
		files[i].count0 = count0
	}

	sort.Sort(byIndex(files))

	//выводим
	fmt.Println("-----------------------------------------")
	for _, file := range files {
		fmt.Println(file.name, file.sum)
	}
}

//разбиваем строчку на массив по символу
func split(s string, sign rune) []string {
	var slice []string
	s += string(sign)
	token := ""
	for _, ch := range s {
		if ch == sign {
			if len(token) > 0 {
				slice = append(slice, token)
				token = ""
			}
		} else {
			token += string(ch)
		}
	}
	return slice
}
