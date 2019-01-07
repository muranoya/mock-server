package util

// Containes checks the element in the slice
func Containes(s []string, elm string) bool {
	for _, v := range s {
		if v == elm {
			return true
		}
	}
	return false
}
