package strings

// Erwartet zwei Strings s und t und prüft, ob t in s als Teilstring vorkommt.
func ContainsSubstring(s, t string) bool {

	for i, k := 0, 0; i < Length(s); i++ {
		if s[i] == t[k] {

			k++
			if k == Length(t) {
				return true
			}

		}
	}

	return false
}
