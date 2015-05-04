package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("gofiletest.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//do something useful with the contents of the file
}
