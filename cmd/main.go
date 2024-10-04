package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		log.Println("Usage: consub <vars file> <template file> <output file>")
		return
	}

	if os.Args[2] == os.Args[3] {
		log.Println("Overwriting the template file is not allowed!")
		return
	}

	vars, err := os.Open(os.Args[1])

	if err != nil {
		log.Println("Error opening variables file: ", err)
		return
	}

	defer vars.Close()
	scanner := bufio.NewScanner(vars)
	scanner.Split(bufio.ScanLines)
	substitute := make(map[string]string)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		key, value, found := strings.Cut(line, "=")

		if found {
			substitute[key] = value
		}
	}

	tpl, err := os.ReadFile(os.Args[2])

	if err != nil {
		log.Println("Error reading template file: ", err)
		return
	}

	tplStr := string(tpl)

	for key, value := range substitute {
		tplStr = strings.ReplaceAll(tplStr, fmt.Sprintf("${%s}", key), value)
	}

	if err := os.WriteFile(os.Args[3], []byte(tplStr), 0644); err != nil {
		log.Println("Error writing file: ", err)
		return
	}
}
