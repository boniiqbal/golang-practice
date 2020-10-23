char := make(map[string]int)
	var max int
	var maxChar string

	wordSplit := strings.Split(word, "")

	for _, v := range wordSplit {
		if char[v] == 0 {
			char[v] = 1
		} else {
			char[v]++
		}
	}

	for i, v := range char {
		if v > max {
			maxChar = i
			max = v
		}
	}