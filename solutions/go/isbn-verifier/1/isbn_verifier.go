package isbnverifier

import (
	"strings"
)

func IsValidISBN(isbn string) bool {
	s := strings.ReplaceAll(isbn, "-", "")

	if len(s) != 10 {
		return false
	}

	n := 0
	c := 10

	for i := 0; i < 10; i++ {
		var val int
		d := s[i]

		switch {
		case d >= '0' && d <= '9':
			val = int(d - '0')
		
		case i == 9 && d == 'X':
			val = 10
		
		default:
			return false
		}

		n += val * c
		c--
	}

	return n%11 == 0
}
