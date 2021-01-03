package file

import (
	"fmt"
	"sort"

	"github.com/helloticket/superfile/field"
	"github.com/helloticket/superfile/helper"
	"github.com/helloticket/superfile/model"
)

type Decoder struct {
	picture field.Picture
	layout  model.Layout
}

func NewDecoder(layout model.Layout) *Decoder {
	return &Decoder{
		picture: field.NewPicture(),
		layout:  layout,
	}
}

func (e Decoder) Decode(blockName string, linha model.Linha) (interface{}, error) {
	return linha.Decode(e.picture)
}

func (e Decoder) Parse(blockName string, source string, layout model.FileConfigMap) []model.Linha {
	linhas := []model.Linha{}

	for field, values := range layout {
		if definition, ok := values.(map[interface{}]interface{}); ok {
			pos := definition["pos"].([]interface{})
			start := helper.ToIntFromSlice(pos, 0)
			end := helper.ToIntFromSlice(pos, 1)

			options := map[string]string{}
			if definition["decimal_separator"] != nil {
				options["decimal_separator"] = helper.ToString(definition["decimal_separator"])
			}

			if global := e.layout.GlobalSettings(); len(global) != 0 {
				for k, v := range global {
					options[fmt.Sprintf("global_%v", k)] = v
				}
			}

			linhas = append(linhas, model.Linha{
				Block:        blockName,
				Name:         helper.ToString(field),
				Start:        start,
				End:          end,
				Size:         end - start + 1,
				DefaultValue: definition["default"],
				Picture:      helper.ToString(definition["picture"]),
				Format:       helper.ToString(definition["data_format"]),
				Options:      options,
				Value:        source,
			})
		}
	}

	sort.Sort(model.LinhaSorted(linhas))

	return linhas
}
