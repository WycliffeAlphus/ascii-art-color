package utils

func IndicesInArr(arr []int, num int) bool {
	for _, x := range arr {
		if x == num {
			return true
		}
	}
	return false
}
