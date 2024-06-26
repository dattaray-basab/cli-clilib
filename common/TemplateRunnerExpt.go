package common

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/dattaray-basab/cks-clilib/globals"
	// "github.com/dattaray-basab/cks-clilib/templates"
)

func RunTemplateExpt(data map[string]map[string][]string, templateText string, templateMap map[string]string, substitutionTemplate globals.SubstitionTemplateT) (string, error) {
	var buf bytes.Buffer
	template.Must(
		template.New("run").Parse(templateText),
	).Execute(&buf, data)
	fmt.Println(buf.String())
	return buf.String(), nil
}
