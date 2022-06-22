package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var lines = 0
var errors = 0

func main() {
	fmt.Print("Enter file name: ")
	var input string
	fmt.Scanln(&input)
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error on load file")
		return
	}
	defer file.Close()
	fmt.Println(input)
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	readFiles(data)
}

func readFiles(data [][]string) {
	for _, line := range data {
		lines++
		renameFile(line[0], line[1])
	}

	fmt.Println("Total files:")
	fmt.Println(lines)
	fmt.Println("Total errors:")
	fmt.Println(errors)
}
func renameFile(from string, to string) {
	err := os.Rename(from, to)
	if err != nil {
		errors++
	}
}
