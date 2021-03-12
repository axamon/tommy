package tommy

type Processo struct {
	IDProcesso     string
	Titolo         string
	Metadata       meta
	Testo          string
	FlussoImmagine string
	RACI           []Attività
}

type Attività struct {
	IDAttività   string
	NomeAttività string
	UO           string
	RuoloLogico  string
	RuoloRACI    string
	Descrizione  string
}

type meta struct {
}
