package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	csvparser := Qwe{}

	for {
		line, err := csvparser.ReadLine(file)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading line:", err)
			return
		}

		fmt.Printf("Read line: %s\n", line)
		fmt.Printf("Number of fields: %d\n", csvparser.GetNumberOfFields())

		for i := 0; i < csvparser.GetNumberOfFields(); i++ {
			field, err := csvparser.GetField(i)
			if err != nil {
				fmt.Printf("Error retrieving field %d: %v\n", i, err)
			} else {
				fmt.Printf("Field %d: %s\n", i, field)
			}
		}
	}
}
