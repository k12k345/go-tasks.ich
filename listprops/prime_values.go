package listprops

// PrimeValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der Primzahlen in der Liste.
func PrimeValues(list []int) int {
	count:=0
	for i:=0; i<len(list);i++{
		if list[i]== 0{
			return false
		}else{
			list[i] %i = 0{
				count++
			}

		}
	}
	return count

}
