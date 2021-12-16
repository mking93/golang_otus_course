package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

func Top10(str string) []string {
	words := strings.Split(str, " ")

	var m map[string]int
	m = getCountUniqueWords(words)

	res := getTop10WordsAlphabetically(m)
	fmt.Println(res)
	return res
}

func getCountUniqueWords(words []string) map[string]int {
	var m map[string]int
	m = make(map[string]int)
	for i := 0; i < len(words); i++ {
		if val, ok := m[words[i]]; ok {
			m[words[i]] = val + 1
			m[words[i]]++
		} else {
			m[words[i]] = 1
		}
	}
	return m
}

func getTop10WordsAlphabetically(m map[string]int) []string {
	type kv struct {
		k string
		v int
	}

	topByValues := make([]kv, 0, len(m))
	for k, v := range m {
		topByValues = append(topByValues, kv{k, v})
	}

	sort.Slice(topByValues, func(i, j int) bool {
		return topByValues[i].v > topByValues[j].v
	})

	count := 0
	if len(topByValues) > 10 {
		count = 10
	} else {
		count = len(topByValues) - 1
	}

	top10 := make([]kv, 0, count)
	for k := 0; k < count; k++ {
		top10 = append(top10, topByValues[k])
	}

	sort.Slice(top10, func(i, j int) bool {
		return top10[i].k < top10[j].k
	})

	top10ByKeys := make([]string, 0, 10)
	for r := 0; r < len(top10); r++ {
		top10ByKeys = append(top10ByKeys, top10[r].k)
	}

	return top10ByKeys
}
