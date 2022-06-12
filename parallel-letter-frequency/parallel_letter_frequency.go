package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	resChannel := make(chan FreqMap, len(l))

	var wg sync.WaitGroup

	wg.Add(len(l))
	for _, s := range l {
		go func(text string) {
			resChannel <- Frequency(text)
			wg.Done()
		}(s)
	}

	wg.Wait()
	close(resChannel)

	res := make(FreqMap)
	for frequency := range resChannel {
		for letter, freq := range frequency {
			res[letter] += freq
		}
	}

	return res
}
