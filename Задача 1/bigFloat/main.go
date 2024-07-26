package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Incorrect usage (app.go file.txt)") // TODO log
		return
	}

	var fileName string = os.Args[1]
	if fileName == "" {
		fmt.Println("Incorrect usage (app.go file.txt)") // TODO log
		return
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Failed to open file :", err) // TODO log
		return
	}
	defer file.Close()

	var res big.Float
	res.SetPrec(128)

	// start := time.Now()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("Failed to convert str to double:", err) // TODO log
			return
		}
		var f big.Float
		f.SetFloat64(number)
		res.Add(&res, &f)
	}
	// fmt.Println(time.Now().Sub(start).Milliseconds()) // TODO log
	result, _ := res.Float64()
	fmt.Println(result) // TODO log
}
