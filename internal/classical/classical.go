package classical

import "unicode"

func RotN(s string, n int) string {
	out := make([]rune, 0, len(s))
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			out = append(out, 'a'+((r-'a')+rune(n))%26)
		case r >= 'A' && r <= 'Z':
			out = append(out, 'A'+((r-'A')+rune(n))%26)
		default:
			out = append(out, r)
		}
	}
	return string(out)
}

func Vigenere(s, key string, decrypt bool) string {
	var ks []rune
	for _, r := range key {
		if unicode.IsLetter(r) {
			kr := unicode.ToLower(r) - 'a'
			ks = append(ks, kr)
		}
	}
	if len(ks) == 0 {
		return s
	}
	out := make([]rune, 0, len(s))
	j := 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			base := rune('A')
			if unicode.IsLower(r) {
				base = 'a'
			}
			k := ks[j%len(ks)]
			shift := k
			if decrypt {
				shift = 26 - k
			}
			out = append(out, base+((r-base)+shift)%26)
			j++
		} else {
			out = append(out, r)
		}
	}
	return string(out)
}
