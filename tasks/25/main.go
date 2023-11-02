package main

import (
	"fmt"
	"time"
)

func Sleeping(sec int) {
	fmt.Printf("Sleep\n")
	<-time.After(time.Second * time.Duration(sec))
	fmt.Printf("Waking up\n")
}

func main() {
	Sleeping(3)
}
