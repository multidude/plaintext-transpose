package green

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/multidude/plaintext-transpose/plug-outputs/ledger"

	"github.com/gocarina/gocsv"
)

// A data structure to match the CSV from Bitstamp
type CSV struct {
	Datetime    string `csv:"time"`
	Description string `csv:"description"`
	TxID        string `csv:"txhash"`
	Conf        string `csv:"confirmations"`
	Amount      string `csv:"amount"`
	Currency    string `csv:"unit"`
	Fee         string `csv:"fee (L-BTC)"`
	Note        string `csv:"memo"`
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
		v.LedgerData[n].Comment1 = r.Note
		v.LedgerData[n].Comment2 = r.TxID
		v.LedgerData[n].Fee = r.Fee

		// Remove dash from currency unit strings
		v.LedgerData[n].Currency = strings.Replace(r.Currency, "-", "", -1)

		// Convert btc string amount to float
		amtflo, _ := strconv.ParseFloat(r.Amount, 32)
		v.LedgerData[n].Amount = r.Amount

		// Is the amount more or less than zero?
		if math.Signbit(amtflo) {
			v.LedgerData[n].Type = "Withdrawal"
		} else {
			v.LedgerData[n].Type = "Deposit"
		}

		// Convert date format
		// Std: Mon Jan 2 15:04:05 MST 2006
		dt, _ := time.Parse("2006-01-02 15:04:05", r.Datetime)
		v.LedgerData[n].Datetime = dt.Format("2006-01-02")
	}
}
