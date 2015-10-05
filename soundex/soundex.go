// Package soundex defines a Match and Encode method that respectively match
// and index string using the Soundex (https://en.wikipedia.org/wiki/Soundex)
// algorithm
package soundex

import "strings"

// Match checks whether two string match via their soundex values
func Match(val1, val2 string) bool {
	return Encode(val1) == Encode(val2)
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

// Encode encodes a string in to its soundex 4 letter representation.
func Encode(val string) string {
	dex := string(val[0])
	letters := strings.ToLower(stripchars(val[1:], "aeiouyh"))
	var previous string
	for _, letter := range letters {
		var current string
		switch letter {
		case 'b', 'f', 'p', 'v':
			current = "1"
		case 'c', 'q', 'j', 'k', 's', 'x', 'z':
			current = "2"
		case 'd', 't':
			current = "3"
		case 'l':
			current = "4"
		case 'm', 'n':
			current = "5"
		case 'r':
			current = "6"
		}

		if previous != current {
			previous = current
			dex = dex + current
			if len(dex) > 3 {
				break
			}
		}
	}
	return dex + strings.Repeat("0", 4-len(dex))
}
