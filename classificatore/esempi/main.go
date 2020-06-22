package main

import (
	"flag"
	"fmt"
	"golem/tommy"
	"golem/tommy/classificatore"
	"log"
	"strings"

	"github.com/axamon/stringset"
)

var file = flag.String("f", "", "file da verificare")

func main() {

	flag.Parse()

	fileAssur := "assurance.pdf"
	fileFullf := "fullfilment.pdf"

	assurance, err := tommy.Pdf2text(fileAssur)
	if err != nil {
		log.Fatal(err)
	}
	fullfilment, err := tommy.Pdf2text(fileFullf)
	if err != nil {
		log.Fatal(err)
	}

	assPeriodi := tommy.GetPeriodi(assurance)
	fullPeriodi := tommy.GetPeriodi(fullfilment)

	primo := stringset.New()
	primo.AddSlice(assPeriodi)
	secondo := stringset.New()
	secondo.AddSlice(fullPeriodi)
	filtro := primo.Intersect(secondo)
	primo = primo.Difference(filtro)
	secondo = secondo.Difference(filtro)

	classificatore.Update(primo.Strings(), classificatore.ASSUR)
	classificatore.Update(secondo.Strings(), classificatore.FULLF)
	classificatore.DumpClassesToFile("./classes")

	extension := strings.Split(*file, ".")[1]

	var text, titoloDOC string

	switch {
	case extension == "pdf":
		// recupera il testo dal file pdf.
		text, err = tommy.Pdf2text(*file)
		if err != nil {
			log.Fatal(err)
		}

		titoloDOC, _ = tommy.GetTitoloPDF(text)

	case extension == "docx":
		// recupera il testo dal file pdf.
		text, err = tommy.Docx2Text(*file)
		if err != nil {
			log.Fatal(err)
		}
		titoloDOC, _ = tommy.GetTitoloDocx(text)
		fmt.Println(titoloDOC)
		//time.Sleep(10 * time.Second)
	}

	testPeriodi := tommy.GetPeriodi(text)

	_, likely, _ := classificatore.Classifier.LogScores(testPeriodi)

	fmt.Println(classificatore.Classifier.Classes[likely])
}
