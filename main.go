package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

	"github.com/multidude/plaintext-transpose/plug-inputs/electrum"
	"github.com/multidude/plaintext-transpose/plug-inputs/green"
	"github.com/multidude/plaintext-transpose/plug-inputs/lnd"
	"github.com/multidude/plaintext-transpose/plug-outputs/ledger"

	"github.com/multidude/plaintext-transpose/plug-inputs/bitstamp"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	src := os.Args[1]
	tplpath := strings.Join([]string{"plug-inputs", src}, "/")
	csvtemp := strings.Join([]string{"ignore-csv", src}, "/")
	csvpath := strings.Join([]string{csvtemp, "csv"}, ".")

	f := openCSV(csvpath)
	defer f.Close()
	var v ledger.ViewData

	// TODO support available plugins
	// requires dynamically loading mods
	// https://appliedgo.net/plugins/
	switch src {
	case "bitstamp":
		var d []bitstamp.CSV
		bitstamp.UnmarshalCSV(f, &d)
		bitstamp.Trans(&d, &v)
	case "lnd":
		var d []lnd.CSV
		lnd.UnmarshalCSV(f, &d)
		lnd.Trans(&d, &v)
	case "electrum":
		var d []electrum.CSV
		electrum.UnmarshalCSV(f, &d)
		electrum.Trans(&d, &v)
	case "green":
		var d []green.CSV
		green.UnmarshalCSV(f, &d)
		green.Trans(&d, &v)
	default:
		fmt.Println("Invalid src argument")
		os.Exit(1)
	}

	dlist := lsDir(tplpath)
	templates, outputs := getPaths(dlist, tplpath)
	for i, f := range templates {
		ledger.ParseTemplates(f, outputs[i], v)
	}
}

func openCSV(filename string) *os.File {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("File open error")
		os.Exit(1)
	}
	return f
}

// Returns a directory listing
func lsDir(tplpath string) []fs.FileInfo {
	f, err := ioutil.ReadDir(tplpath)
	if err != nil {
		fmt.Println("Directory listing error")
		os.Exit(1)
	}
	return f
}

// Construct complete paths to each file in a directory listing
func getPaths(dlist []fs.FileInfo, tplpath string) ([]string, []string) {
	var fn string
	templates := make([]string, len(dlist))
	outputs := make([]string, len(dlist))
	for i, f := range dlist {
		fn = f.Name()
		if strings.HasSuffix(fn, ".ledger") {
			templates[i] = tplpath + "/" + fn
			outputs[i] = "ignore-output/" + fn
			fmt.Println(fn)
		}
	}
	return templates, outputs
}
