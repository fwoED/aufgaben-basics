package strings

// Erwartet zwei strings s1 und s2.
// Liefert einen String, der abwechselnd aus den Buchstaben des einen und des anderen
// Strings besteht.
func Zip(s1, s2 string) string {
	result := ""
	k := 0
	t := 0
	if Length(s1) >= Length(s2) {
		k = Length(s1)
		t = Length(s2)
		for i := 0; i < k; i++ {
			result += string(s1[i])
			if t != 0 {
				result += string(s2[i])
				t--
			}
		}
		return result
	} else {
		k = Length(s2)
		t = Length(s1)
		for i := 0; i < k; i++ {
			if t != 0 {
				result += string(s1[i])
				t--
			}
			result += string(s2[i])
		}
		return result
	}

}
