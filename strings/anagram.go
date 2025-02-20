package strings

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
func IsAnagram(s1, s2 string) bool {

	gleichzähler := 0
	for i := 0; i < len(s1); i++ {
		for k := 0; k < len(s2); k++ {
			if s1[i] == s2[k] {
				gleichzähler++
				s2 = s2[:k] + "*" + s2[k+1:]
			}

		}

	}
	if gleichzähler == len(s1) {
		return true
	}
	return false
}

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
// Diese Funktion soll keinen Unterschied zwischen Groß- und Kleinschreibung machen.
func IsAnagramIgnoreCase(s1, s2 string) bool {
	// TODO
	return false
}
