package hw03frequencyanalysis

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

type dicItem struct {
	Key   string
	Value int
}

var regCompile = regexp.MustCompile(`(([a-zA-Zа-яА-Я]+\S[a-zA-Zа-яА-Я]+)|[a-zA-Zа-яА-Я]+)`)

func Top10(str string) []string {
	if len(str) == 0 {
		return make([]string, 0)
	}

	words := regCompile.FindAllString(strings.ToLower(str), -1)

	wordsCount := make(map[string]int)
	for _, w := range words {
		wordsCount[w]++
	}

	wordsData := make([]dicItem, 0, len(wordsCount))
	for key, val := range wordsCount {
		wordsData = append(wordsData, dicItem{key, val})
	}

	sort.Slice(wordsData, func(first, second int) bool {
		if wordsData[first].Value == wordsData[second].Value {
			return wordsData[first].Key < wordsData[second].Key
		}
		return wordsData[first].Value > wordsData[second].Value
	})

	resultLen := int(math.Min(10, float64(len(wordsData))))
	result := make([]string, resultLen)
	for i := 0; i < resultLen; i++ {
		result[i] = wordsData[i].Key
	}

	return result
}
