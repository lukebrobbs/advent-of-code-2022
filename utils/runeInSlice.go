package utils

func RuneInSlice(r rune, l []rune) bool {
	for _, v := range l {
		if v == r {
			return true
		}
	}
	return false
}
