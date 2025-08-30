package main

import (
	"bufio"
	"fmt"
	"os"
	"project/calc"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("------ simple CLI calculator ------")
	result , err := calc.Operation(scan, "Enter operation (e.g., + 5, * 2, / 4, or 'exit'): ")
	if err != nil {
		fmt.Println("Ops try again", err)
		return
	}
	fmt.Println(result)
	fmt.Println("------ End ------")
	
}