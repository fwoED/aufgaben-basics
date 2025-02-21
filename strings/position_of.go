package strings

// Erwartet einen String s und einen Buchstaben c.
// Liefert die Position, an der c in s vorkommt.
// Liefert die Länge von s, falls c nicht in s vorkommt.
// Kommt c mehrfach vor, soll die erste Position geliefert werden.
func PositionOf(s string, c byte) int {
	if s == "" {
		return 0
	}

	for i := 0; i < 1+Length(s[1:]); i++ {
		if c == s[i] {
			result := i
			return result
		}
	}
	return 1 + Length(s[1:])
}
