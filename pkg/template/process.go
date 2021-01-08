package template

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"os"
	"text/template"
)

type Templates struct {
	Box *packr.Box
}

// NewTemplates generates a new instance of the
// Templates struct with a packr box
func NewTemplates(box *packr.Box) *Templates {
	return &Templates{Box: box}
}

func (t *Templates) loadFromBox(file string) (string, error) {
	s, err := t.Box.FindString(file)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	return s, nil
}

func (t *Templates) ProcessString(str string, vars interface{}) string {
	tmpl, err := template.New("tmpl").Funcs(template.FuncMap(sprig.FuncMap())).Parse(str)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	return process(tmpl, vars)
}

func (t *Templates) ProcessFile(fileName string, vars interface{}) string {
	file, err := t.loadFromBox(fileName)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	log.Infof("Loaded template from file: '%s'", fileName)
	tmpl, err := template.New("tmpl").Funcs(template.FuncMap(sprig.FuncMap())).Parse(file)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	log.Debugf("Parsed template from file %s", fileName)
	return process(tmpl, vars)
}

// process applies the data structure 'vars' onto an already
// parsed template 't', and returns the resulting string.
func process(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	return tmplBytes.String()
}
