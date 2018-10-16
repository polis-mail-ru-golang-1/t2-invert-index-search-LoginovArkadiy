package myFile

var Count = 0

type MyFile struct {
	Name    string
	strfile string
	Sum     int
	count0  int
	hm      map[string]int
}

func NewMyFile(Name string, data []byte) MyFile {
	mF := MyFile{
		Name:    Name,
		count0:  0,
		Sum:     0,
		strfile: string(data),
		hm:      make(map[string]int),
	}

	mF.createIndex()
	return mF
}

func (file *MyFile) createIndex() {
	var words []string
	file.strfile += " "
	var token string
	for _, ch := range file.strfile {
		if ch == ' ' {
			if len(token) > 0 {
				words = append(words, token)
				token = ""
			}
		} else {
			token += string(ch)
		}
	}

	for _, word := range words {
		file.hm[word]++

	}

}

func (file *MyFile) Analyse(words []string) {
	file.Sum = 0
	file.count0 = 0
	for _, word := range words {
		x := file.hm[word]
		if x == 0 {
			file.count0++
		}
		file.Sum += x
	}
}

type ByIndex []MyFile

func (s ByIndex) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByIndex) Less(i, j int) bool {
	f1 := s[i]
	f2 := s[j]
	if f1.count0 < f2.count0 {
		return true
	}
	if f1.count0 > f2.count0 {
		return false
	}

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
