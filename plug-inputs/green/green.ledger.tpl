{{range .LedgerData}}
; {{.Comment2}}
{{.Datetime}} * {{.Comment1 }} 
{{- if eq .Type "Deposit"}}
	Assets:Green				{{.Amount}} {{.Currency}}
	Income:Ambiguous
{{- end}}
{{- if eq .Type "Withdrawal"}}
	Assets:Green				{{.Amount}} {{.Currency}}
		{{- if .Fee}}
	Expenses:Fees:OnChain			{{.Fee}} LBTC
		{{- end}}
	Expenses:Ambiguous
{{- end}}
{{end}}
