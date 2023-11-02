package main

import "fmt"

func main() {

	temp_storage := make(map[int][]float64)

	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, value := range temp {
		// fmt.Println("value", int(value/10)*10)
		diff := int(value/10) * 10
		// if _, ok := temp_storage[diff]; ok {
		// 	temp_storage[diff] = append(temp_storage[diff], value)
		// } else {
		// 	temp_storage[diff] = []float64{value}
		// }
		temp_storage[diff] = append(temp_storage[diff], value)
	}

	for k, v := range temp_storage {
		fmt.Printf("%d : {%v}\n", k, v[:])
	}
}
