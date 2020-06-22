package main

import (
	"fmt"
	"golem/tommy"
	"log"
	"os"
)

// import (
// 	"fmt"
// 	"go/doc"
// 	"golem/tommy"
// 	"os"
// )

// func main() {

// 	tommy.ParoleComuni(os.Args[1], os.Args[2])

// 	package main

func main() {

	text, err := tommy.Docx2Text(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(str)

	titolo, _ := tommy.GetTitoloDocx(text)

	fmt.Println(titolo)

	periodi := tommy.GetPeriodi(text)
	for i, periodo := range periodi {
		fmt.Println(i, periodo)
	}
}
