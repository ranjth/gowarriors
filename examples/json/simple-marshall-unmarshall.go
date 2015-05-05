package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Dmr struct {
	ID int
	Name string
	Nicknames []string
}

func main() {
	ikshi := Dmr{
		ID:1,
		Name: "Dhiya",
		Nicknames: []string{"Meena ikshi kuttiB", "xyz"},
	}

	b, err := json.Marshal(ikshi)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	os.Stdout.Write(b)
	fmt.Println()

	var kutties Dmr
	err = json.Unmarshal(b, &kutties)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%+v\n", kutties)
	for _, v := range kutties.Nicknames {
		fmt.Println(v)
	}
}
