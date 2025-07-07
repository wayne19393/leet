package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFile := "output.conf"
	outputFile := "deduped_output.conf"

	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer func() {
		if err := inFile.Close(); err != nil {
			fmt.Println("Error closing input file:", err)
		}
	}()

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer func() {
		if err := outFile.Close(); err != nil {
			fmt.Println("Error closing output file:", err)
		}
	}()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)

	seen := make(map[string]bool)

	for {
		if !scanner.Scan() {
			break
		}
		line1 := scanner.Text()

		if !strings.HasPrefix(line1, "local-data:") {
			continue
		}

		if !scanner.Scan() {
			break
		}
		line2 := scanner.Text()

		hostname := extractHostname(line1)
		if hostname == "" {
			continue
		}

		if !seen[hostname] {
			_, err := fmt.Fprintln(writer, line1)
			if err != nil {
				fmt.Println("Error writing line1:", err)
				os.Exit(1)
			}

			_, err = fmt.Fprintln(writer, line2)
			if err != nil {
				fmt.Println("Error writing line2:", err)
				os.Exit(1)
			}

			_, err = fmt.Fprintln(writer)
			if err != nil {
				fmt.Println("Error writing newline:", err)
				os.Exit(1)
			}

			seen[hostname] = true
		}
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing output:", err)
		os.Exit(1)
	}

	fmt.Println("Deduplicated output written to", outputFile)
}

// extractHostname extracts the hostname from a local-data line
func extractHostname(line string) string {
	start := strings.Index(line, "\"")
	end := strings.Index(line, ".megabank.")
	if start == -1 || end == -1 || end <= start {
		return ""
	}
	return line[start+1 : end]
}
