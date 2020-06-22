package tommy

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"sync"
)

var mappaparole = make(map[string][]string)

var mapparoleLock sync.RWMutex

// var path = "mappaparole.gob"

// ArchiviaParoleSuFile salva su un unico file .gob tutte le parole dei documenti
// in una mappa.
func ArchiviaParoleSuFile(path, codicetommy string, parole []string) error {

	err := createFile(path)
	if err != nil {
		return err
	}

	// recupera set parole condivise da file
	err = caricaMappa(path, &mappaparole)
	if err != nil {
		log.Println(err)
	}
	// aggiungi al set quelle in input
	mappaparole[codicetommy] = parole

	// salva
	err = salvaMappa(path, mappaparole)
	if err != nil {
		return fmt.Errorf("salvaMappa error: %v", err)
	}

	return nil
}

func salvaMappa(path string, object interface{}) error {
	mapparoleLock.Lock()
	defer mapparoleLock.Unlock()
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewEncoder(file).Encode(object)
}

func caricaMappa(path string, object interface{}) error {
	mapparoleLock.Lock()
	defer mapparoleLock.Unlock()
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("caricaMappa error: %v", err)
	}
	defer file.Close()
	return gob.NewDecoder(file).Decode(object)
}

func createFile(path string) error {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
		fmt.Println("Created file: ", path)
	}

	return nil
}

// PrintParoleArchiviate recupera le parole archiviate e le
// stampa a video.
func PrintParoleArchiviate(codicetommy, gobfile string) ([]string, error) {
	var m = make(map[string][]string)
	err := caricaMappa(gobfile, &m)
	if err != nil {
		return nil, err
	}
	var listaparole []string
	for _, v := range m[codicetommy] {

		listaparole = append(listaparole, v)

	}
	return listaparole, nil
}

// ElencaDocArchiviati restituisce la lista dei codici
// dei documenti le cui parole sono archiavate.
func ElencaDocArchiviati(gobfile string) ([]string, error) {

	var m = make(map[string][]string)
	err := caricaMappa(gobfile, &m)
	if err != nil {
		return nil, err
	}
	var listacodici []string
	for k := range m {
		// fmt.Println(k)
		listacodici = append(listacodici, k)
	}
	return listacodici, nil
}
