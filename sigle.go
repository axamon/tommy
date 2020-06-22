package tommy

import "regexp"

var sigleregex = regexp.MustCompile(`(?m)(?P<sigla>([A-Z-/]+\.)+[A-Z0-9]+\.[A-Z]+)`)

// Sigle recupera tutte le sigle nel testo.
func Sigle(text string) []string {

	
	funzioni := sigleregex.FindAllString(text, -1)

	var m = make(map[string]struct{}, len(funzioni))

	for _, funzione := range funzioni {
		m[funzione] = struct{}{}
	}

	var list []string
	for k := range m {
		list = append(list, k)
	}
	return list
}
