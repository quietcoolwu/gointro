// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	word2File := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, word2File)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, word2File)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, word2File[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, word2File map[string][]string) {
	input := bufio.NewScanner(f)
	fileName := f.Name()
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if !wordInMap(fileName, word2File[text]) {
			word2File[text] = append(word2File[text], fileName)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

func wordInMap(word string, arr []string) bool {
	for _, ele := range arr {
		if ele == word {
			return true
		}
	}
	return false
}

//!-
