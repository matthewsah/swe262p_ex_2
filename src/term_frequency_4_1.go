package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type pair struct {
	word  string
	count int
}

func main() {
	// word_freqs := make([]pair, 0)

	stopWords, err := os.ReadFile("../stop_words.txt")
	if err != nil {
		log.Fatal(err)
	}
	stopWordsList := strings.Split(string(stopWords), ",")
	asciiVal := 65
	for i := 0; i < 26; i++ {
		char := string(rune(asciiVal + i))
		stopWordsList = append(stopWordsList, char)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
