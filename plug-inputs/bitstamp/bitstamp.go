package bitstamp

import (
	"fmt"
	"os"
	"time"

	"github.com/multidude/plaintext-transpose/plug-outputs/ledger"

	"github.com/gocarina/gocsv"
)

// A data structure to match the CSV from Bitstamp
type CSV struct {
	Type     string `csv:"Type"`
	Datetime string `csv:"Datetime"`
	Account  string `csv:"Account"`
	Amount   string `csv:"Amount"`
	Value    string `csv:"Value"`
	Rate     string `csv:"Rate"`
	Fee      string `csv:"Fee"`
	SubType  string `csv:"Sub Type"`
}

// Unmarshal the CSV into a struct
func UnmarshalCSV(f *os.File, d *[]CSV) {
	if err := gocsv.UnmarshalFile(f, d); err != nil {
		fmt.Println("Unmarshal error")
		fmt.Println(err)
		os.Exit(1)
	}
}

// Convert to Ledger format and pack it up into ledger.ViewData
func Trans(d *[]CSV, v *ledger.ViewData) {
	v.LedgerData = make([]ledger.Ledger, len(*d))
	for n, r := range *d {
		// The easy stuff
		v.LedgerData[n].Type = r.Type
		v.LedgerData[n].Amount = r.Amount
		v.LedgerData[n].Value = r.Value
		v.LedgerData[n].Fee = r.Fee
		v.LedgerData[n].SubType = r.SubType
		// Convert date format
		dt, _ := time.Parse("Jan. 2, 2006, 03:04 PM", r.Datetime)
		v.LedgerData[n].Datetime = dt.Format("2006-01-02")
		// Unused in current tpl:
		v.LedgerData[n].Account = r.Account
		v.LedgerData[n].Rate = r.Rate

	}
}
