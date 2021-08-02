package main

import (
	"bytes"
	"text/template"
)

type TemplateInfo struct {
	Text string
	data map[string]interface{}
}

func (t *TemplateInfo) GenCodeByTemplate() (*bytes.Buffer, error) {
	tem, err := template.New("gen_code").Parse(t.Text)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err = tem.Execute(buf, t.data); err != nil {
		return nil, err
	}
	return buf, nil
}
