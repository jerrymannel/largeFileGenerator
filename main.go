package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func usage() {
	s := `
Usage: largeFileGenerator <opFileName> <type> <size>
Parameters:
	opFileName  :	Output file name
	type        :	json/xml/csv/bin
	size        :	Multiples of 1KB
	`
	fmt.Println(s)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generate(fileType, outputFileName, data string, size int) {
	lines := ((size * 1024) / len(data)) + 1

	f, err := os.Create(outputFileName)
	check(err)
	w := bufio.NewWriter(f)

	var rootTag string
	if fileType == "xml" {
		rootTag = strings.SplitN(data, ",", 2)[0]
		data = strings.SplitN(data, ",", 2)[1]
		_, err = w.WriteString("<" + rootTag + ">\n")
		check(err)
		w.Flush()
	}
	for i := 0; i < lines; i++ {
		_, err = w.WriteString(data + "\n")
		check(err)
		w.Flush()
	}
	if fileType == "xml" {
		_, err = w.WriteString("<" + rootTag + "/>")
		check(err)
		w.Flush()
	}
}

func main() {

	if len(os.Args) < 4 {
		fmt.Println("Insufficient parameters")
		usage()
		os.Exit(1)
	}

	opFileName := os.Args[1]
	fileType := os.Args[2]
	size, err := strconv.Atoi(os.Args[3])
	check(err)

	homeDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	stringsFile := path.Join(homeDir, "strings.txt")
	file, err := os.Open(stringsFile)
	if err != nil {
		panic(err)
	}

	strings := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strings = append(strings, line)
	}

	switch fileType {
	case "json":
		generate(fileType, opFileName, strings[0], size)
	case "xml":
		generate(fileType, opFileName, strings[1], size)
	case "csv":
		generate(fileType, opFileName, strings[2], size)
	}
}
