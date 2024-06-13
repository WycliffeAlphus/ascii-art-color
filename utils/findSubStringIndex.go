package utils

func FindSubStringIndex(mainString, subString string) []int {
	indices := []int{}
	for i := 0; i <= len(mainString)-len(subString); i++ {
		// fmt.Println(mainString[i : i+len(subString)])
		if mainString[i:i+len(subString)] == subString {
			for j := i; j < i+len(subString); j++ {
				indices = append(indices, j)
			}
		}
	}
	// fmt.Println(indices)
	return indices
}
