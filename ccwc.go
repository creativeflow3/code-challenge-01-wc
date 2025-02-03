package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func generateOutput(switchType string, contents []byte, isStdIn bool) string {
	scanner := bufio.NewScanner(os.Stdin)
	if !isStdIn {
		scanner = bufio.NewScanner(strings.NewReader(string(contents)))
	}
	var counts int

	switch switchType {
	case "bytes":
		scanner.Split(bufio.ScanBytes)
	case "lines":
		scanner.Split(bufio.ScanLines)
	case "words":
		scanner.Split(bufio.ScanWords)
	case "characters":
		scanner.Split(bufio.ScanRunes)
	default:
		scanner.Split(bufio.ScanBytes)
	}

	for scanner.Scan() {
		counts++
	}
	output := "   " + strconv.Itoa(counts)
	return output
}

func main() {
	c := flag.Bool("c", false, "output the number of bytes")
	l := flag.Bool("l", false, "output the number of lines")
	w := flag.Bool("w", false, "output the number of words")
	m := flag.Bool("m", false, "output the number of characters")

	flag.Parse()

	filePath := os.Args[len(os.Args)-1]
	file, err := os.Open(filePath)
	var isStdIn = false

	// no file provided as a parameter, check stdout
	if err != nil {
		file := os.Stdin
		fi, err := file.Stat()

		if err != nil {
			fmt.Println("You must define a <filename> or pipe in file countents", err)
			return
		}
		size := fi.Size()
		if size > 0 {
			isStdIn = true
			filePath = ""
		} else {
			fmt.Println("Stdin is empty")
			return
		}
	}

	var contents []byte

	// if there is a file provided by the parameter
	if !isStdIn {
		cont, err := io.ReadAll(file)

		if err != nil {
			fmt.Println("File cannot be read")
			return
		}
		contents = cont
	}

	output := ""

	if !*c && !*w && !*l && !*m {
		*c = true
		*w = true
		*l = true
	}

	if *c {
		output = output + generateOutput("bytes", contents, isStdIn)
	}

	if *l {
		output = output + generateOutput("lines", contents, isStdIn)
	}

	if *w {
		output = output + generateOutput("words", contents, isStdIn)
	}

	if *m {
		output = output + generateOutput("characters", contents, isStdIn)
	}

	output = output + "   " + filePath
	fmt.Println(output)
}
