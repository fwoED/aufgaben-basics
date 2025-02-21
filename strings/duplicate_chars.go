package strings

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {
	if s == "" {
		return s
	}
	result_string := ""
	for i := 0; i < 1+Length(s[1:]); i++ {
		result_string += string(s[i]) + string(s[i])
	}
	return result_string
}
