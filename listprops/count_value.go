package listprops

// CountValue gibt erwartet eine Liste von Zahlen und eine Zahl v.
// Liefert die Anzahl der Vorkommen von v in der Liste.
func CountValue(list []int, v int) int {
	count := 0
	// Hinweis:
	// Verwenden Sie eine for-Schleife, um die Liste zu durchlaufen.
	// Prüfen Sie in jeder Iteration, ob das aktuelle Element gleich v ist.
	// Erhöhen Sie den Zähler entsprechend.
	// begin:solution
	for _, el := range list {
		if el == v {
			count++
		}
	}
	// end:solution
	return count
}
