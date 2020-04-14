package model

import (
	"sort"
	"strings"
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

	preSegmentos := []Segmento{}
	for _, key := range keys {
		if len(records[key].Valores) == 0 && !l.SegmentoVazio {
			continue
		}

		preSegmentos = append(preSegmentos, records[key])
	}

	segmentos := []Segmento{}

	for _, p := range preSegmentos {
		nomes := strings.Split(p.Nome, ".")

		if len(nomes) == 1 {
			segmentos = append(segmentos, p)
		} else {
			segmentos = append(segmentos, Segmento{
				Nome:    nomes[1],
				Valores: p.Valores,
			})
		}
	}

	return segmentos
}
