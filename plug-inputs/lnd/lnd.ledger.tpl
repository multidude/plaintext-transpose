{{range .LedgerData}}
{{- if .Comment2}}
; {{.Comment2}}
{{- end}}
{{.Datetime}} * {{.Type}}
	{{.Account}}			{{.Amount}}
	{{.Remainder}}
{{end}}
