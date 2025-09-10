package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введите строку")
	bf := bufio.NewReader(os.Stdin)
	str, _, _ := bf.ReadLine()
	strArray := strings.Split(string(str), " ")
	l := len(strArray)
	for i := 0; i < l/2; i++ {
		buf := strArray[i]
		strArray[i] = strArray[l-1-i]
		strArray[l-1-i] = buf
	}
	fmt.Println(strings.Join(strArray, " "))
}
