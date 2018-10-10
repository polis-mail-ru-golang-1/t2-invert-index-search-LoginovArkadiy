// pr project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello World!")
	bs, err := ioutil.ReadFile("stdin.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	fileout, err := os.Create("stdout.txt")
	if err != nil {
		fmt.Println("Error creating file")
		return
	}
	defer fileout.Close()
	str := string(bs)
	str += " "
	fmt.Println(str)

	//fmt.Println(string(str[0]) + string(str[1]) + string(str[2]) + string(str[len(str)-1]))

	var a []int
	var count string
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			if len(count) > 0 {
				x, err := strconv.Atoi(count)
				if err != nil {
					fmt.Println("Error in conver int to string, count ="+count+" len =", len(count))
					return
				}
				a = append(a, x)
				count = ""
			}
		} else {
			count += string(str[i])
		}

	}

	fmt.Println(a)

	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j-1] > a[j] {
				k := a[j]
				a[j] = a[j-1]
				a[j-1] = k
			} else {
				break
			}
		}
	}

	for _, val := range a {
		fileout.WriteString(strconv.Itoa(val) + " ")
	}

}
