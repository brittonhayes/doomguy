package template

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	log "github.com/sirupsen/logrus"
	"text/template"
)

func Usage(command string) string {

	var b bytes.Buffer
	vars := map[string]interface{}{
		"Usage": fmt.Sprintf("```yaml\n+%s```", command),
	}

	tpl := template.Must(
		template.New("base").Funcs(template.FuncMap(sprig.FuncMap())).Parse(`*Example* {{ .Usage }}`),
	)

	if err := tpl.Execute(&b, vars); err != nil {
		log.Error(err)
	}
	return b.String()
}
