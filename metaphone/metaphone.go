package metaphone

import "strings"

func stripDupes(val string) string {
	for i := 0; i < len(val)-1; i++ {
		if val[i+1] == val[i] {
			splice := val[:i] + val[i+1:]
			val = stripDupes(splice)
		}
	}
	return val
}

func stringAt(start, length int, s string, strings ...string) bool {
	if start < 0 {
		return false
	}

	s = s[start : start+length]
	for _, test := range strings {
		if test == s {
			return true
		}
	}

	return false
}

func transform(val string) string {
	var word string
	for i, char := range val {
		switch char {
		case 'c':
			if val[i:i+2] == "ia" || (val[i+1] == 'h' && (i-1 < 0 || val[i-1] != 's')) {
				char = 'x'
			} else if strings.IndexRune("iey", rune(val[i+1])) > -1 {
				char = 's'
			} else {
				char = 'k'
			}
		case 'd':
			switch val[i : i+2] {
			case "ge", "gy", "gi":
				char = 'j'
			default:
				char = 't'
			}
		case 'g':

		case 'z':
			char = 's'
		}
		word = word + string(char)
	}

	return word
}

func Encode(val string) string {
	dex := strings.ToLower(stripDupes(val))
	switch dex[:2] {
	case "kn", "gn", "pn", "ae", "wr":
		dex = dex[1:]
	}

	if dex[len(dex)-2:] == "mb" {
		dex = dex[:len(dex)-1]
	}

	dex = transform(dex)

	return dex
}
