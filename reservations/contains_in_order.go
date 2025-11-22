package reservations

// ContainsInOrder erwartet eine Liste von Strings und zwei Strings s1 und s2.
// Liefert true, falls sowohl s1 als auch s2 in der Liste enthalten sind, und s1 vor s2 kommt.
func ContainsInOrder(list []string, s1 string, s2 string) bool {

	foundS1 := false
	foundS2 := false
	pos1 := 0
	pos2 := 0
	for i := 0; i < len(list); i++ {
		if list[i] == s1 {
			foundS1 = true
			pos1 = i
		}
		for j := 0; j < len(list); j++ {
			if list[j] == s2 {
				foundS2 = true
				pos2 = j
			}
		}
	}
	return foundS1 && foundS2 && pos1 < pos2

}
