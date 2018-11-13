package myFile

import "strings"

var Count = 0

type MyFile struct {
	Name    string
	strfile string
	Sum     int
	count0  int
	Words   []string
}

func (file *MyFile) SetData(sum, count0 int) {
	file.Sum = sum
	file.count0 = count0
}

func NewMyFile(Name string, data []byte) MyFile {
	mF := MyFile{
		Name:    Name,
		count0:  0,
		Sum:     0,
		strfile: string(data),
	}

	mF.createSlice()
	return mF
}

func (file *MyFile) createSlice() {
	file.Words = strings.Fields(file.strfile)
}

type ByIndex []MyFile

func (s ByIndex) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByIndex) Less(i, j int) bool {
	f1 := s[i]
	f2 := s[j]
	/*	if f1.count0 < f2.count0 {
			return true
		}
		if f1.count0 > f2.count0 {
			return false
		}
	*/
	if f1.Sum > f2.Sum {
		return true
	}
	if f1.Sum < f2.Sum {
		return false
	}

	return f1.Name < f2.Name
}

func (s ByIndex) Len() int {
	return Count
}
