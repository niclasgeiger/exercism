package letter

const testVersion = 1

func ConcurrentFrequency(languages []string) FreqMap {
	out := FreqMap{}
	freqChans := []chan FreqMap{}
	for _, language := range languages {
		c := make(chan FreqMap)
		go func(input string) {
			c <- Frequency(input)
		}(language)
		freqChans = append(freqChans, c)
	}
	for i := 0; i < len(languages); i++ {
		out.Add(<-freqChans[i])
	}
	return out
}

func (f FreqMap) Add(freqMap FreqMap) {
	for r, val := range freqMap {
		f[r] += val
	}
}
