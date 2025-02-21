package strings

// Erwartet einen String s und einen Buchstaben c.
// Prüft, ob c in s vorkommt.
func Contains(s string, c byte) bool {
	if s == "" {
		return false
	}
	for i := 0; i < Length(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false

}
