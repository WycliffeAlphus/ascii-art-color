package colorfeature

func IsSubstring(s1, s2 string) bool {

	count := 0
	for i:=0; i < len(s1);{
		for j:=0; j < len(s2);{
			if s1[i] == s2[j]{
				count ++
				j++	
			}
			i ++
			if i == len(s1){
				break
			}
		}
	}
	if count == len(s1) && count==len(s2){
		return true
	}
		return false
	
}

func FindSubStringIndex(mainString, subString string) []int {
	indices := []int{}
	for i:=0; i<len(mainString);i++{
		if  i+len(subString) <= len(mainString) && IsSubstring(subString, mainString[i:i+len(subString)]){
			for j:=i ; j < i+len(subString);j++{
				indices= append(indices, j)
			}
		}
	}
	return indices
}
func IndicesInArr(arr []int, num int) bool {
	for _,x := range arr{
		
			if x == num{
				return true
			}
	}
	return false
}