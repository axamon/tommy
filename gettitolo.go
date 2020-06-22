package tommy

import (
	"errors"
	"regexp"
	"strings"
)

var title = regexp.MustCompile(`(?m)(.*PUBBLICAZIONE Modello selezionato)(?P<titolo>.*(?P<id>\d{19}).*)(Opzioni.*)`)
var space = regexp.MustCompile(`\s+`)

var title2 = regexp.MustCompile(`(?m).*Nome(?P<title2>.*)\n.*`)
var titleDocx = regexp.MustCompile(`(?m)TITOLO:\s(?P<titolo>.*)`)

// GetTitoloPDF recupera il titolo di un documento tommy.
func GetTitoloPDF(text string) (string, error) {

	var substitution = "${titolo}"

	titolo := title.ReplaceAllString(text, substitution)

	titolo = identificativo.ReplaceAllString(titolo, "")
	titolo = space.ReplaceAllString(titolo, " ")

	return strings.TrimSpace(titolo), nil
}

// GetTitoloDocx recupera il titolo di un documento tommy.
func GetTitoloDocx(text string) (string, error) {

	titololist := titleDocx.FindAllString(text, -1)
	if len(titololist) < 1 {
		return "", errors.New("Titolo non trovato")
	}

	titolo := strings.Split(titololist[0], ":")[1]

	return strings.TrimSpace(titolo), nil
}
