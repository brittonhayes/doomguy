package template

import (
	"github.com/dimuska139/rawg-sdk-go"
)

func (t *Templates) GameDetails(d *rawg.GameDetailed) string {
	vars := map[string]interface{}{
		"ID":            d.ID,
		"Name":          d.Name,
		"Description":   d.DescriptionRaw,
		"Developer":     d.Developers[0].Name,
		"Rating":        d.Rating,
		"Slug":          d.Slug,
		"Metacritic":    d.Metacritic,
		"MetacriticURL": d.MetacriticUrl,
		"Website":       d.Website,
	}
	return t.ProcessFile("game.tmpl", vars)
}
