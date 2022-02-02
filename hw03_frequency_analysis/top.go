package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type KV struct {
	k string
	v int
}

type SKV struct {
	k []string
	v int
}

var regCompile = regexp.MustCompile(`[\s!.,-?]+`)

func Top10(text string) []string {
	return getTop10WordsAlphabetically(getCountUniqueWords(text))
}

func getCountUniqueWords(text string) map[string]int {
	words := make(map[string]int)
	for _, elem := range regCompile.Split(text, -1) {
		elem = strings.ToLower(elem)
		_, found := words[elem]
		if found {
			words[elem]++
		}
	}

	return words
}

func getTop10WordsAlphabetically(m map[string]int) []string {
	words := make([]KV, 0, len(m))

	for key, val := range m {
		words = append(words, KV{key, val})
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i].v > words[j].v
	})

	wordsByNumberRepetitions := make(map[int][]string)

	for _, val := range words {
		wordsByNumberRepetitions[val.v] = append(wordsByNumberRepetitions[val.v], val.k)
	}

	for key := range wordsByNumberRepetitions {
		sort.Slice(wordsByNumberRepetitions[key], func(i, j int) bool {
			return wordsByNumberRepetitions[key][i] < wordsByNumberRepetitions[key][j]
		})
	}

	topByValues := make([]SKV, 0, len(wordsByNumberRepetitions))

	for key, val := range wordsByNumberRepetitions {
		topByValues = append(topByValues, SKV{val, key})
	}

	sort.Slice(topByValues, func(i, j int) bool {
		return topByValues[i].v > topByValues[j].v
	})

	top10 := make([]string, 0, len(topByValues))

	for _, elem := range topByValues {
		top10 = append(top10, elem.k...)
	}

	return top10
}
