package tommy

import (
	"fmt"

	"log"
	"time"

	"github.com/axamon/stringset"
)

//ParoleComuni trova le parole comuni nei documenti PDF che non sono utili
//ai fini delle verifiche lessicografiche.
func ParoleComuni(path1, path2 string) {

	text1, err := Pdf2text(path1)
	if err != nil {
		log.Fatal(err)
	}
	text2, err := Pdf2text(path2)
	if err != nil {
		log.Fatal(err)
	}
	periodi1 := GetPeriodi(text1)

	periodi2 := GetPeriodi(text2)

	primo := stringset.New()
	secondo := stringset.New()

	primo.AddSlice(periodi1)
	secondo.AddSlice(periodi2)

	intersect := primo.Intersect(secondo)

	fmt.Println(intersect.Len())
	time.Sleep(10 * time.Second)
	for _, periodo := range intersect.Strings() {
		fmt.Println(periodo)
	}

	return
}
