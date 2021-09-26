package ledger

import (
	"fmt"
	"os"
	"text/template"
)

// The data used by the template
type Ledger struct {
	Type      string
	Datetime  string
	Account   string
	Amount    string
	Currency  string
	Value     string
	Rate      string
	Fee       string
	SubType   string
	Remainder string
	Comment1  string
	Comment2  string
	Comment3  string
}

type ViewData struct {
	LedgerData []Ledger
}

// Parse the template with site data and output to a new file.
func ParseTemplates(tplpath string, opath string, v ViewData) {
	t, err := template.ParseFiles(tplpath)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create(opath)
	defer f.Close()
	if err != nil {
		fmt.Println("create file: ", err)
		return
	}

	err = t.Execute(f, v)
	if err != nil {
		fmt.Println("execute: ", err)
		return
	}
}
