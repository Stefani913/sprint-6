package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) string {
	if checkIfAnyIsCyrillic(input) {
		return morse.DefaultConverter.ToMorse(input)
	} else {
		return morse.DefaultConverter.ToText(input)
	}
}

func checkIfAnyIsCyrillic(input string) bool {
	minCyrillic := rune(0x0410) // 'А'
	maxCyrillic := rune(0x044F) // 'я'
	// U+0030 // '1'
	// U+0039 // '9'
	for _, i := range input {
		if i >= minCyrillic && i <= maxCyrillic {
			return true
		}
	}
	return false
	/*morse := reverseEncodingMap()
	checkFun := func(r rune) bool {
		for _, ch := range morse {
			return ch == r
		}
		return false
	}

	return strings.ContainsFunc(input, checkFun)*/
}

/*func reverseEncodingMap() []rune {
	morse := morse.DefaultMorse
	result := make([]rune, len(morse))

	for r := range morse {
		result = append(result, r)
	}

	return result
}*/
