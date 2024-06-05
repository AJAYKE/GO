package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welwel := "vachavara brahmi"
	fmt.Println(welwel)

	reader := bufio.NewReader(os.Stdin)

	println("yedhokati rayara: ")

	// comma ok // err ok
	input, err := reader.ReadString('\n')
	fmt.Println("em enter chesadu vadu?", input)
	fmt.Println("error emanna vachindha?? vachindhiiii: ", err)
}