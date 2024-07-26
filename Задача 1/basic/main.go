package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Incorrect usage") // TODO log
		return
	}

	var fileName string = os.Args[1]
	if fileName == "" {
		fmt.Println("Incorrect usage") // TODO log
		return
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Failed to open file :", err) // TODO log
		return
	}
	defer file.Close()

	var res float64 = 0
	// start := time.Now()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("Failed to convert str to double:", err) // TODO log
			return
		}
		res += number
	}
	// fmt.Println(time.Now().Sub(start).Milliseconds()) // TODO log

	fmt.Println(res) // TODO log
}
