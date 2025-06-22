package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseOrText(s string) string {
	foundMorse := false

	for _, value := range morse.DefaultMorse {
		if strings.ContainsAny(s, value) {
			foundMorse = true
			break
		}
	}
	if foundMorse {
		// fmt.Print(morse.ToMorse(s))
		return morse.ToText(s)

	}
	// fmt.Print(morse.ToText(s))
	return morse.ToMorse(s)

}
