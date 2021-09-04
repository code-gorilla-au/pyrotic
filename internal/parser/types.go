package parser

import "text/template"

type TmplEngine struct {
	root  *template.Template
	funcs template.FuncMap
}

type TemplateData struct {
	Name   string
	To     string
	Append bool
	Output []byte
	Meta   map[string]interface{}
}