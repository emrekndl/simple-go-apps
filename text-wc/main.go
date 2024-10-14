package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.Open("text.txt")
	check(err)
	defer file.Close()

	word_frequencies := make(map[string]int)
	i := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word_frequencies[word]++
		i += 1
		// fmt.Println(word)
	}
	check(scanner.Err())

	fmt.Printf("%s has %d words.\n", file.Name(), i)

	fmt.Println(word_frequencies)
}
