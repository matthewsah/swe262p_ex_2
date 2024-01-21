package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

type pair struct {
	word  string
	count int
}

func readFile(path string) []rune {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	charData := make([]rune, 0)

	for scanner.Scan() {
		line := scanner.Text() + " "
		for _, char := range line {
			charData = append(charData, char)
		}
	}

	return charData
}

func filterCharsAndNormalize(charData []rune) []string {
	stringData := make([]string, 0)

	for idx, char := range charData {
		if !(rune('a') <= char && char <= rune('z') || rune('A') <= char && char <= rune('Z') || rune('0') <= char && char <= rune('9')) {
			charData[idx] = ' '
		} else {
			charData[idx] = unicode.ToLower(char)
		}
	}

	for _, char := range charData {
		stringData = append(stringData, string(char))
	}

	return stringData
}

func scan(stringData []string) []string {
	text := strings.Join(stringData, "")
	words := strings.Split(text, " ")
	return words
}

func removeStopwords(words []string) []string {
	stopwords, err := os.ReadFile("../stop_words.txt")
	if err != nil {
		log.Fatal(err)
	}
	stopwordsList := strings.Split(string(stopwords), ",")
	asciiVal := 97
	for i := 0; i < 26; i++ {
		char := string(rune(asciiVal + i))
		stopwordsList = append(stopwordsList, char)
	}
	stopwordsList = append(stopwordsList, string(""))

	stopwordIdxs := make([]int, 0)
	for idx, word := range words {
		for _, stopword := range stopwordsList {
			if word == stopword {
				stopwordIdxs = append(stopwordIdxs, idx)
				break
			}
		}
	}

	for i := len(stopwordIdxs) - 1; i >= 0; i-- {
		idx := stopwordIdxs[i]
		left := words[:idx]
		right := words[idx+1:]
		words = append(left, right...)
	}

	return words
}

func frequencies(words []string) []pair {
	wordFreqs := make([]pair, 0)

	for _, word := range words {
		found := false
		for idx, pair := range wordFreqs {
			if word == pair.word {
				wordFreqs[idx].count += 1
				found = true
				break
			}
		}
		if !found {
			wordFreqs = append(wordFreqs, pair{word, 1})
		}
	}

	return wordFreqs
}

func sortFreqs(wordFreqs []pair) []pair {
	sort.SliceStable(wordFreqs, func(p1, p2 int) bool {
		return wordFreqs[p1].count > wordFreqs[p2].count
	})
	return wordFreqs
}

func printFreqs(wordFreqs []pair) {
	for _, pair := range wordFreqs[:25] {
		fmt.Println(pair.word, " - ", pair.count)
	}
}

func main() {
	printFreqs(sortFreqs(frequencies(removeStopwords(scan(filterCharsAndNormalize(readFile(os.Args[1])))))))
}
