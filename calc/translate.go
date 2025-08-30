package calc

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Operation(scan *bufio.Scanner, text string) (float64 , error) {
	sum := 0.00
	for {
		fmt.Println("current number :",sum)
		fmt.Printf("%s", text)
		scan.Scan()
		input := strings.TrimSpace(scan.Text())
		if input == "" {
			fmt.Println("please write a operator and number")
			continue
		}
		if input == "exit" {
			break
		}
		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("please write a operator or number")
			continue
		}

		number , err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Ops something wrong try again :",err)
			continue
		}
		 operator := parts[0]
		 
		switch operator {
		case "+":
			sum += float64(number)
		case "-":
			sum -= float64(number)
		case "*":
			sum *= float64(number)
		case "/":
			if number == 0 {
				fmt.Println("please write a value defering zero")
				continue
			}

			sum /= float64(number)
		default:
			fmt.Println("please try a operator", operator)
		}
	
	
	}
	fmt.Println("good day")
	return sum, nil

}
