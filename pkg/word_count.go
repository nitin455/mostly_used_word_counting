package pkg

import (
	"fmt"
	"sort"
)

type WordFreq struct {
	word string
	freq int
}

func (p WordFreq) String() string {
	return fmt.Sprintf("%s %d", p.word, p.freq)
}

func CountMostlyUsedWords(inputWords []string) {
	words := make(map[string]int)

	for _, match := range inputWords {
		words[match]++
	}

	var wordFreqs []WordFreq
	for k, v := range words {
		wordFreqs = append(wordFreqs, WordFreq{k, v})
	}

	sort.Slice(wordFreqs, func(i, j int) bool {

		return wordFreqs[i].freq > wordFreqs[j].freq
	})

	for i := 0; i < len(wordFreqs); i++ {
		fmt.Println(wordFreqs[i].String())
	}
}
