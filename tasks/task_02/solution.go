package main

func rotateRunes(s string, shift int) string {

	r := []rune(s)
	n := len(r)

	if n == 0 {
		return ""
	}

	shift = shift % n
	if shift == 0 {
		return string(r)
	}

	resultR := make([]rune, n)

	if shift > 0 {
		for i := range r {
			resultR[i] = r[(i+shift)%n]
		}
	} else {
		for i := range r {
			resultR[i] = r[(i+shift+n)%n]
		}
	}

	return string(resultR)
}
