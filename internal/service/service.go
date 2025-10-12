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

	for _, i := range input {
		if i >= minCyrillic && i <= maxCyrillic {
			return true
		}
	}
	return false
}
