package tommy

import (
	"fmt"
	"strings"
)

// Importanti restituisce le frasi utilizzate più volte nel testo.
func Importanti(testo string) map[string]float64 {

	// indice di significatività relativa dei segmenti.
	var IS = make(map[string]float64)

	periodi := GetPeriodi(testo)

	// calcolo frequenze delle parole singole
	var frequenze = make(map[string]int)

	// calcolo delle corrispondenze di gruppi di parole nel testo
	var importanti = make(map[string]int)
	// var importanti3 = make(map[string]int)
	// var importanti2 = make(map[string]int)

	for i := 0; i < len(periodi); i++ {
		words := strings.Split(periodi[i], " ")
		for _, parola := range words {
			frequenze[parola]++
		}
		for s := 0; s < len(words); s++ {
			if s+3 < len(words) {
				segmento := (words[s] + " " + words[s+1] + " " + words[s+2] + " " + words[s+3])
				for u := i + 1; u < len(periodi); u++ {
					if strings.Contains(periodi[u], segmento) {
						importanti[segmento]++
					}
				}
			}
			if s+2 < len(words) {
				segmento := (words[s] + " " + words[s+1] + " " + words[s+2])
				for u := i + 1; u < len(periodi); u++ {
					if strings.Contains(periodi[u], segmento) {
						importanti[segmento]++
					}
				}
			}
			if s+1 < len(words) {
				segmento := (words[s] + " " + words[s+1])
				for u := i + 1; u < len(periodi); u++ {
					if strings.Contains(periodi[u], segmento) {
						importanti[segmento]++
					}
				}
			}
		}
	}

	ignoramap := make(map[string]struct{})
	ignora := "si no non il lo la le e di a da in con su per tra fra dalla della sulla al allo alla"
	for _, i := range strings.Split(ignora, " ") {
		ignoramap[i] = struct{}{}
	}

	// calcolo indici di significatività delle occorrenze importanti
	for segmento, freqSegmento := range importanti {
		if freqSegmento > 1 {
			elementi := strings.Split(segmento, " ")
			l := len(elementi)
			for _, elemento := range elementi {
				IS[segmento] += float64(freqSegmento / frequenze[elemento])
				if _, ispresent := ignoramap[elemento]; ispresent {
					l--
				}
			}
			IS[segmento] *= float64(l)
			IS[segmento] /= float64(len(elementi) * len(elementi))

		}
	}

	fmt.Println(importanti)
	return IS

}