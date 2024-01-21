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
	wordFreqs := make([]pair, 0)

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

	fmt.Println(stopwordsList)

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	startChar := -1

	for scanner.Scan() {
		line := scanner.Text() + " "
		startChar = -1

		for idx, char := range line {
			if startChar == -1 {
				if rune('a') <= char && char <= rune('z') || rune('A') <= char && char <= rune('Z') || rune('0') <= char && char <= rune('9') {
					startChar = idx
				}
			} else {
				if !(rune('a') <= char && char <= rune('z') || rune('A') <= char && char <= rune('Z') || rune('0') <= char && char <= rune('9')) {
					found := false
					word := strings.ToLower(line[startChar:idx])

					isStopword := false
					for _, stopword := range stopwordsList {
						if word == stopword {
							isStopword = true
							break
						}
					}

					if !isStopword {
						pairIdx := 0
						for idx, curPair := range wordFreqs {
							if word == curPair.word {
								wordFreqs[idx] = pair{curPair.word, curPair.count + 1}
								found = true
								break
							}
							pairIdx += 1
						}

						if !found {
							wordFreqs = append(wordFreqs, pair{word, 1})
						} else {
							for n := pairIdx - 1; n >= 0; n-- {
								if wordFreqs[pairIdx].count > wordFreqs[n].count {
									wordFreqs[n], wordFreqs[pairIdx] = wordFreqs[pairIdx], wordFreqs[n]
									pairIdx--
								} else {
									break
								}
							}
						}
					}
					startChar = -1
				}
			}
		}
	}

	for _, pair := range wordFreqs[:25] {
		fmt.Println(pair.word, "-", pair.count)
	}
}
