package components

func toUpper(c rune) rune {
	if c >= 'a' && c <= 'z' {
		return c - ('a' - 'A')
	}
	return c
}

func toLower(c rune) rune {
	if c <= 'A' && c >= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func ToLowerRunes(runes []rune) []rune {
	for i, c := range runes {
		runes[i] = toLower(c)
	}
	return runes
}

func ToLower(s string) string {
	runes := []rune(s)
	return string(ToLowerRunes(runes))
}

func ToUpperRunes(runes []rune) []rune {
	for i, c := range runes {
		runes[i] = toUpper(c)
	}
	return runes
}

func ToUpper(s string) string {
	runes := []rune(s)
	return string(ToUpperRunes(runes))
}

func isLowercaseInTitle(s string) bool {
	var smallWords = map[string]bool{
		"a": true, "an": true, "and": true, "as": true, "at": true,
		"but": true, "by": true, "for": true, "in": true, "nor": true,
		"of": true, "on": true, "or": true, "so": true, "the": true,
		"to": true, "up": true, "yet": true, "with": true,
	}
	return smallWords[s]
}

func parseTitle(s string) [][]rune {
	var words [][]rune
	var word []rune
	for _, c := range s {
		if c != ' ' {
			word = append(word, c)
		} else {
			words = append(words, word)
			word = word[:0]
		}
	}
	return append(words, word)
}

func ToTitle(s string) string {
	words := parseTitle(s)
	var runes []rune
	nospace := true
	for _, word := range words {
		if nospace {
			runes = append(runes, ' ')
		} else {
			nospace = false
		}
		if isLowercaseInTitle(string(word)) {
			runes = append(runes, ToLowerRunes(word)...)
		} else {
			runes = append(runes, word...)
		}
	}
	return string(runes)
}
