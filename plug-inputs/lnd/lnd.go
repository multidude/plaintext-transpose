package lnd

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

// A data structure to match the CSV from LND
type CSV struct {
	Timestamp    string `csv:"Timestamp"`
	OnChain      string `csv:"OnChain"`
	Type         string `csv:"Type"`
	Category     string `csv:"Category"`
	AmountMsat   string `csv:"Amount(Msat)"`
	AmountUSD    string `csv:"Amount(USD)"`
	TxID         string `csv:"TxID"`
	Reference    string `csv:"Reference"`
	BTCPrice     string `csv:"BTCPrice"`
	BTCTimestamp string `csv:"BTCTimestamp"`
	Note         string `csv:"Note"`
}

// Unmarshal the LND CSV into a struct
func UnmarshalCSV(f *os.File, d *[]CSV) {
	if err := gocsv.UnmarshalFile(f, d); err != nil {
		fmt.Println("Unmarshal error")
		os.Exit(1)
	}
}

func mSatToBTC(r *CSV, v *ledger.ViewData, n *int) {
	// Convert mSat to BTC
	// divide by 1,000 and round down to nearest integer
	// divide result by 100,000,000
	// print with 8 digit accuracy
	// append currency code BTC
	mflo, _ := strconv.ParseFloat(r.AmountMsat, 32)
	satflo := mflo / 1000
	sats := math.Floor(satflo)
	btcflo := sats / 100000000
	v.LedgerData[*n].Amount = fmt.Sprintf("%.8f BTC", btcflo)
}

func chaininess(r *CSV, chain *string) {
	switch r.OnChain {
	case "true":
		*chain = "OnChain"
	case "false":
		*chain = "OffChain"
	default:
		fmt.Println("WARNING OnChain is neither True nor False")
		*chain = "Ambiguous"
	}
}

// Transpose the CSV export from faraday (lnd, bitcoin)
// into LedgerData, in a ledger.ViewData Container
// TODO: return a basic slice, reformat in main
func Trans(d *[]CSV, v *ledger.ViewData) {
	var chain string
	var sats float64
	v.LedgerData = make([]ledger.Ledger, len(*d))
	for n, r := range *d {
		// Convert date format
		dt, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", r.Timestamp)
		v.LedgerData[n].Datetime = dt.Format("2006-01-02")

		mSatToBTC(&r, v, &n)

		v.LedgerData[n].Type = r.Type

		chaininess(&r, &chain)

		asset := []string{"Assets", "LND", chain}
		feesO := []string{"Expenses", "Fees", chain}
		feesI := []string{"Income", "Fees", chain}
		expen := []string{"Expenses", "Misc", chain}
		incom := []string{"Income", "Deposit", chain}

		v.LedgerData[n].Account = strings.Join(asset, ":")

		switch r.Type {
		case "CHANNEL_OPEN_FEE", "CHANNEL_CLOSE_FEE", "FEE", "FORWARD", "FORWARD_FEE":
			if sats <= 0 {
				v.LedgerData[n].Remainder = strings.Join(feesO, ":")
			} else {
				v.LedgerData[n].Remainder = strings.Join(feesI, ":")
			}
		case "LOCAL_CHANNEL_OPEN":
			// Transfer from onchain to offchain
			v.LedgerData[n].Remainder = "Assets:LND:OffChain"
		case "PAYMENT":
			v.LedgerData[n].Remainder = strings.Join(expen, ":")
		case "RECEIPT":
			v.LedgerData[n].Remainder = strings.Join(incom, ":")
		case "CHANNEL_CLOSE":
			v.LedgerData[n].Remainder = strings.Join(incom, ":")
		}
		v.LedgerData[n].Comment1 = r.Note
		v.LedgerData[n].Comment2 = r.TxID
		v.LedgerData[n].Comment3 = r.Reference
	}
}
