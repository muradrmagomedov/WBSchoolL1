package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100])) // Создаем независимую от v строку
}

func main() {
	someFunc()
}
