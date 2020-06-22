package classificatore

import (
	"os"

	"github.com/axamon/bayesian"
)

const (
	// Strategy, Infrastructure & Product
	SIP bayesian.Class = "Strategy Infrastructure & Product"
	SC  bayesian.Class = "Strategy & Commit"
	ILM bayesian.Class = "Infrastructure Lifecycle Management"
	PLM bayesian.Class = "Product Lifecycle Management"
	// Operations
	OP    bayesian.Class = "Operations"
	OSR   bayesian.Class = "Operations Support & Readiness"
	FULLF bayesian.Class = "Fulfillment"
	ASSUR bayesian.Class = "Assurance"
	BRM   bayesian.Class = "Billing & Revenue Management"
)

//Classifier contiene le classi utile per la classificazione dei documenti ETOM.
var Classifier = bayesian.NewClassifier(OSR, FULLF, ASSUR, BRM, SC, ILM, PLM)

//NewClass crea una nuova classe.
func NewClass() *bayesian.Class {
	return new(bayesian.Class)
}

//Update permette di inserire nuove parole nel Classificatore.
func Update(parole []string, class bayesian.Class) {
	Classifier.Learn(parole, class)
	return
}

//DumpClassesToFile salva ogni singola classe in Classifier su file.
func DumpClassesToFile(path string) error {
	_, err := os.Open(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, os.FileMode(777))
	}
	for _, classe := range Classifier.Classes {
		//	fmt.Println(classe)
		err := Classifier.WriteClassToFile(classe, path)
		if err != nil {
			return err
		}
	}
	return nil
}

//SalvaClassificatore salva su fileName i dati del classificatore bayesiano.
func SalvaClassificatore(fileName string) error {

	return Classifier.WriteToFile(fileName)
}
