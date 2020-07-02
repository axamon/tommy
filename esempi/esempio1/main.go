package main

import (
	"flag"
	"fmt"
	"log"
	"sort"

	"github.com/axamon/tommy"

	"github.com/axamon/stringset"
)

var codiceDOC = flag.String("doc", "", "Codice univoco del documento Tommy 2012-00123")
var gobFile = flag.String("file", "mappaparole.gob", "File .gob con le parole salavate")

func main() {

	flag.Parse()

	listacodici, err := tommy.ElencaDocArchiviati(*gobFile)
	if err != nil {
		log.Fatal(err)
	}
	//var set = make(map[string]*stringset.StringSet)
	var set []*stringset.StringSet
	for _, codice := range listacodici {
		listaparole, err := tommy.PrintParoleArchiviate(codice, *gobFile)
		if err != nil {
			log.Println(err)
		}
		s := stringset.NewStringSet()
		//	fmt.Println(listaparole)
		s.AddSlice(listaparole)
		//set[codice] = s
		set = append(set, s)
	}

	// recupera parole comuni ovunque
	var intersezione *stringset.StringSet
	for i := range set {
		if i == 0 {
			intersezione = set[0].Intersect(set[1])
			continue
		}
		intersezione = intersezione.Intersect(set[i])
	}

	fmt.Println("len: ", intersezione.Len())
	fmt.Println()
	inutili := intersezione.Strings()
	sort.Slice(inutili, func(i, j int) bool { return inutili[i] < inutili[j] })
	fmt.Println(inutili)

	err = tommy.ArchiviaParoleSuFile("intersezione.gob", "intersezione", intersezione.Strings())
	if err != nil {
		log.Println(err)
	}
}
