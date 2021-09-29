{{range .LedgerData}}
{{.Datetime}} * {{.Type}} {{.SubType}}
{{- if eq .Type "Deposit"}}
	Assets:Bitstamp			{{.Amount}}
	Income:Deposit
{{- end}}
{{- if eq .Type "Withdrawal"}}
	Assets:Bitstamp			-{{.Amount}}
		{{- if .Fee}}
	Expenses:Fees:Exchange		{{.Fee}}
		{{- end}}
	Expenses:Withdrawal
{{- end}}
{{- if eq .Type "Market"}}
	{{- if eq .SubType "Buy"}}
	Assets:Bitstamp			{{.Amount}} @@ {{.Value}}
		{{- if .Fee}}
	Expenses:Fees:Exchange		{{.Fee}}
		{{- end}}
	Assets:Bitstamp
	{{- end}}
	{{- if eq .SubType "Sell"}}
	Assets:Bitstamp			-{{.Amount}} @@ {{.Value}}
		{{- if .Fee}}
	Expenses:Fees:Exchange		{{.Fee}}
		{{- end}}
	Assets:Bitstamp
	{{- end}}
{{- end}}
{{end}}
