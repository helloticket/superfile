package file

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/helloticket/superfile/field"
	"github.com/helloticket/superfile/helper"
	"github.com/helloticket/superfile/model"
)

type Encoder struct {
	model   *model.Remessa
	picture field.Picture
}

func NewEncoder(model *model.Remessa) *Encoder {
	return &Encoder{
		model:   model,
		picture: field.NewPicture(),
	}
}

func (e Encoder) ParseAndEncode(blockName string, source map[string]interface{}, layout map[interface{}]interface{}) string {
	linhas := e.Parse(blockName, source, layout)
	encode := e.Encode(blockName, linhas)
	return encode
}

func (e Encoder) Parse(blockName string, source model.RecordMap, layout model.FileConfigMap) []model.Linha {
	linhas := []model.Linha{}

	for field, values := range layout {
		if definition, ok := values.(map[interface{}]interface{}); ok {
			pos := definition["pos"].([]interface{})
			start := helper.ToIntFromSlice(pos, 0)
			end := helper.ToIntFromSlice(pos, 1)

			linhas = append(linhas, model.Linha{
				Block:        blockName,
				Name:         helper.ToString(field),
				Start:        start,
				End:          end,
				Size:         end - start + 1,
				DefaultValue: definition["default"],
				Picture:      helper.ToString(definition["picture"]),
				Format:       helper.ToString(definition["data_format"]),
				Value:        source[helper.ToString(field)],
			})
		}
	}

	sort.Sort(model.LinhaSorted(linhas))

	return linhas
}

func (e Encoder) Encode(blockName string, linhas []model.Linha) string {
	buffer := strings.Builder{}

	var sum int64 = 0
	for _, l := range linhas {
		encoded := l.Encode(e.picture)
		sum += int64(l.Size)

		if l.Size != len(encoded) {
			panic(fmt.Sprintf(`
				Block: %s
				Field: %s
				Picture: %s
				Length: %v <=> %v
			`, blockName, l.Name, l.Picture, l.Size, len(encoded)))
		}

		buffer.WriteString(encoded)
	}

	if sum != int64(buffer.Len()) {
		panic(fmt.Sprintf(`
			Block: %s
			Length: %v <=> %v
		`, blockName, sum, buffer.Len()))
	}

	if e.model.TamanhoRegistro() != -1 {
		if e.model.TamanhoRegistro() != sum {
			log.Println(buffer.String())
			log.Println(int64(buffer.Len()))

			panic(fmt.Sprintf(`
			Block: %s
			Layout: %d
			Result: %v
		`, blockName, e.model.TamanhoRegistro(), sum))
		}
	}

	return buffer.String()
}
