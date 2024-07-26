package main

import (
	"bufio"
	"fmt"
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

	var (
		res float64 = 0
		error1 float64 = 0
		delta float64
	)

	// start := time.Now()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("Failed to convert str to double:", err) // TODO log
			return
		}
		res = sumWithFix(res, number, &delta)
		error1 += delta
	}
	res = res + error1
	// fmt.Println(time.Now().Sub(start).Milliseconds()) // TODO log
	fmt.Println(res)
}

func sumWithFix(a, b float64, error *float64) float64 {
	var (
		x          = a + b
		b_virt     = x - a
		a_virt     = x - b_virt
	)
	*error = (b - b_virt) + (a - a_virt)

	return x
}
