package main

import (
	"fmt"
	"os"
	"calculator/calculate"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments. Expected one.")
		os.Exit(1)
	}
	str := os.Args[1]
	result, err := calculate.Calculate(str)
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(result)
}
