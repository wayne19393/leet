package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFile := "input.txt"
	outputFile := "output.conf"

	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer func(inFile *os.File) {
		err := inFile.Close()
		if err != nil {
			fmt.Println("Error closing input file:", err)
		}
	}(inFile)

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			fmt.Println("Error closing output file:", err)
		}
	}(outFile)

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		if len(line) < 3 {
			// Skip lines with fewer than 3 fields
			continue
		}

		ip := line[0]
		hostnames := line[2:] // all names after second column

		for _, name := range hostnames {
			_, err := fmt.Fprintf(writer, "local-data: \"%s.megabank. 3600 IN A %s\"\n", name, ip)
			if err != nil {
				fmt.Println("Error writing local-data:", err)
				return
			}

			_, err = fmt.Fprintf(writer, "local-data-ptr: \"%s %s.megabank.\"\n\n", ip, name)
			if err != nil {
				fmt.Println("Error writing local-data-ptr:", err)
				return
			}
		}
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing output:", err)
		return
	}

	fmt.Println("Output written to", outputFile)
}
