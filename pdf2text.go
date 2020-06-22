package tommy

import (
	"log"
	"strings"

	"github.com/dcu/pdf"
)

// Pdf2text estae il testo da file pdf.
func Pdf2text(path string) (string, error) {
	var content []string

	f, r, err := pdf.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}

	totalPage := r.NumPage()

	// cicla ogni pagina
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, err := p.GetTextByRow()
		if err != nil {
			log.Print(err)
		}

		for _, row := range rows {
			for _, word := range row.Content {
				content = append(content, word.S)
				//fmt.Println(word.S)
			}
		}
	}
	return strings.Join(content, " "), nil
}
