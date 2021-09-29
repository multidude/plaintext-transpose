{{range .LedgerData}}
; {{.Comment2 }}
{{.Datetime}} * {{.Comment1 }}
{{- if eq .Type "Deposit"}}
    Assets:Electrum			{{.Amount}} BTC
    Income:Ambiguous
{{- end}}
{{- if eq .Type "Withdrawal"}}
    Assets:Electrum			{{.Amount}} BTC
		{{- if .Fee}}
    Expenses:Fees:OnChain		{{.Fee}} BTC
		{{- end}}
    Expenses:Ambiguous
{{- end}}
{{end}}
