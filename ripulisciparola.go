package tommy

import "unicode"

// Ripulisciparola elimina l'ultima runa di una parola se non è alfanumerica.
func Ripulisciparola(s string) string {
	runes := []rune(s)

	lastRune := runes[len(runes)-1]

	if !unicode.IsDigit(lastRune) && !unicode.IsLetter(lastRune) {
		runes = runes[:len(runes)-1]
	}

	return string(runes)
}
