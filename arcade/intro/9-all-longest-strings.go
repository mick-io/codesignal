package codesignal

func allLongestStrings(inputArray []string) (longest []string) {
	maxLength := len(inputArray[0])
	for _, s := range inputArray {
		switch length := len(s); {
		case length == maxLength:
			longest = append(longest, s)
		case length > maxLength:
			maxLength = length
			longest = []string{s}
		}
	}
	return longest
}
