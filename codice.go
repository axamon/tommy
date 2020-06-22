package tommy

import (
	"errors"
	"regexp"
)

var codex = regexp.MustCompile(`(?m)[0-9]{4}-[0-9]{5}`)

// Codice il codice univoco del documento.
func Codice(text string) (string, error) {

	codice := codex.FindString(text)
	if codice == "" {
		return "", errors.New("Nessun codice univoco trovato")
	}
	return codice, nil
}
