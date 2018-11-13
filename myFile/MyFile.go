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
