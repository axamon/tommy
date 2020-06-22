package tommy

import (
	"fmt"
	"os"
	"sync"

	"code.sajari.com/docconv"
)

var lock sync.RWMutex

//Docx2Text restituisce il testo di un documento docx.
func Docx2Text(docx string) (string, error) {
	lock.Lock()
	defer lock.Unlock()

	f, err := os.Open(docx)
	defer f.Close()
	if err != nil {
		return "", fmt.Errorf("Errore, %s non può esssere aperto", docx)
	}
	text, _, err := docconv.ConvertDocx(f)
	// res, err := docconv.ConvertPath(docx)
	if err != nil {
		return "", err
	}

	return text, nil

}
