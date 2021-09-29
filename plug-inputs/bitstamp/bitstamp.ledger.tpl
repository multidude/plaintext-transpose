{{range .LedgerData}}
{{.Datetime}} * {{.Type}} {{.SubType}}
{{- if eq .Type "Deposit"}}
	Assets:Bitstamp:{{.Account}}		{{.Amount}}
	Income:Deposit
{{- end}}
{{- if eq .Type "Withdrawal"}}
	Assets:Bitstamp:{{.Account}}		-{{.Amount}}
		{{- if .Fee}}
	Expenses:Fees:Exchange			{{.Fee}}
		{{- end}}
	Expenses:Withdrawal
{{- end}}
{{- if eq .Type "Market"}}
	{{- if eq .SubType "Buy"}}
	Assets:Bitstamp:{{.Account}}		{{.Amount}} @@ {{.Value}}
		{{- if .Fee}}
	Expenses:Fees:Exchange			{{.Fee}}
		{{- end}}
	Assets:Bitstamp:{{.Account}}
	{{- end}}
	{{- if eq .SubType "Sell"}}
	Assets:Bitstamp:{{.Account}}		-{{.Amount}} @@ {{.Value}}
		{{- if .Fee}}
	Expenses:Fees:Exchange			{{.Fee}}
		{{- end}}
	Assets:Bitstamp:{{.Account}}
	{{- end}}
{{- end}}
{{end}}
