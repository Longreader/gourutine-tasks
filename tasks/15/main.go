package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var justString string

func createHugeString(size int) string {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	hugeString := strings.Builder{}
	for i := 0; i < size; i++ {
		hugeString.WriteRune('a' + rune(generator.Intn('z'-'a'+1)))
	}
	return hugeString.String()
}

func someFunc() {

	hugeString := createHugeString(1 << 10)

	// Проблема первого метода связанна с тем,
	// что после определения justString, данный слайс
	// указывал на тот же массив, что и hugeString

	tmp := make([]rune, 100)
	// Функция copy позволяет скопировать элементы массива
	// без привязки к исходному массиву, что исключает ссылочную
	// связь и позволяет отрабоать GC
	copy(tmp, []rune(hugeString)[:100])
	justString = string(tmp)

	fmt.Println("justString = ", justString)
	fmt.Printf("Pointer is %p\n", &justString)
	fmt.Println("hugeString = ", hugeString)
	fmt.Printf("Pointer is %p\n", &hugeString)

}
func main() {
	someFunc()
}
