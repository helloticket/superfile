package model

import (
	"sort"
)

type Lote struct {
	Header        map[string]interface{}
	Trailer       map[string]interface{}
	Sequencial    int64
	layout        Layout
	detalhes      []Detalhe
	SegmentoVazio bool
}

func (l *Lote) NovoDetalhe() Detalhe {
	detalhe := Detalhe{}

	for key, value := range l.layout.GetSegmentoDefinitions() {
		detalhe[key] = value
	}

	return detalhe
}

func (l *Lote) InserirDetalhe(novo Detalhe) {
	l.detalhes = append(l.detalhes, novo)
}

func (l *Lote) Segmentos() []Segmento {
	keys := []string{}
	records := map[string]Segmento{}

	for _, detalhe := range l.detalhes {
		for key, value := range detalhe {
			keys = append(keys, key)
			records[key] = Segmento{
				Nome:    key,
				Valores: value,
			}
		}
	}

	sort.Strings(keys)

	segmentos := []Segmento{}
	for _, key := range keys {
		if l.SegmentoVazio {
			segmentos = append(segmentos, records[key])
		} else {
			if len(records[key].Valores) > 0 {
				segmentos = append(segmentos, records[key])
			}
		}
	}

	return segmentos
}
