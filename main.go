package main

import (
	"fmt"

	"github.com/ardatak1992/gator/internal/config"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("fawgag")
	fmt.Println(conf)
	
}
