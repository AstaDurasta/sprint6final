package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseOrText(s string) string {
	foundtext := false

	for key, _ := range morse.DefaultMorse {
		if strings.ContainsAny(strings.ToUpper(s), string(key)) && (string(key) != "." && string(key) != "-") {
			foundtext = true
			break
		}
	}
	if foundtext {
		// fmt.Print(morse.ToMorse(s))
		return morse.ToMorse(s)

	}
	// fmt.Print(morse.ToText(s))
	fmt.Print(foundtext)
	return morse.ToText(s)

}
