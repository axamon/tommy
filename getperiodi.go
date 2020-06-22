package tommy

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`(?m)[[:punct:]]+`)

// GetPeriodi divide il testo in periodi.
func GetPeriodi(testo string) []string {

	str := re.ReplaceAllString(testo, "\n")

	var periodi []string
	for _, frase := range strings.Split(str, "\n") {
		if len(frase) == 0 {
			continue
		}
		periodi = append(periodi, strings.Trim(frase, " \t\n,.;:\""))
	}

	return periodi
}
