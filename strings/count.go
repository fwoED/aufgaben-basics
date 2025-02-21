package strings

// Erwartet einen string s und zählt, wie oft der Buchstabe 'A' in s vorkommt.

func CountA(s string) int {
	result := 0
	A := "A"
	for i := 0; i < 1+Length(s[1:]); i++ {
		if A[0] == s[i] {
			result++
		}
	}
	if result != 0 {
		return result
	}
	return 0
}

// Erwartet einen string s und einen Buchstaben c.
// Zählt, wie oft c in s vorkommt.
func CountChar(s string, c rune) int {

	result := 0
	for _, char := range s {
		if char == c {
			result++
		}
	}
	return result

}
