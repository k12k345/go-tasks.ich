package reservations

// ContainsBoth erwartet eine Liste von Strings und zwei Strings s1 und s2.
// Liefert true, falls sowohl s1 als auch s2 in der Liste enthalten sind, sonst false.
func ContainsBoth(list []string, s1 string, s2 string) bool {

	for i := 0; i < len(list); i++ {
		if list[i] == s1 {
			for j := 0; j < len(list); j++ {
				if list[j] == s2 {
					return true
				}
			}

		}
	}
	return false
}
