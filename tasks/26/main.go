package main

import (
	"fmt"
	"strings"
)

func IsUnque(str string) bool {
	dict := make(map[rune]struct{})

	str = strings.ToLower(str)

	for _, symb := range str {
		if _, ok := dict[symb]; ok {
			return false
		} else {
			dict[symb] = struct{}{}
		}
	}
	return true
}

func main() {
	list := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, elem := range list {
		fmt.Printf("Str %s is %t\n", elem, IsUnque(elem))
	}

}
