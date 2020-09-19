package model

import (
	"fmt"
	"sort"
	"strings"

	"github.com/helloticket/superfile/helper"
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
	records, ordenacao := l.ordernarSegmentos()

	preSegmentos := []Segmento{}
	for _, key := range ordenacao {
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

func (l *Lote) ordernarSegmentos() (map[string]Segmento, []string) {
	type entry struct {
		key    string
		sorte1 interface{}
		sorte2 interface{}
	}

	entries := []entry{}
	records := map[string]Segmento{}
	i := 0

	for _, detalhe := range l.detalhes {
		for okey, fields := range detalhe {
			i = i + 1
			key := fmt.Sprintf("%d.%s", i, okey)

			sorte1, sorte2 := l.definirEstrategiaOrdenacao(i, okey, fields)

			entries = append(entries, entry{
				key:    key,
				sorte1: sorte1,
				sorte2: sorte2,
			})

			records[key] = Segmento{
				Nome:    key,
				Valores: fields,
			}
		}
	}

	sort.Slice(entries, func(i, j int) bool {
		return helper.Less(entries[i].sorte1, entries[j].sorte1) &&
			helper.Less(entries[i].sorte2, entries[j].sorte2)
	})

	keys := []string{}
	for _, k := range entries {
		keys = append(keys, k.key)
	}

	return records, keys
}

func (l *Lote) definirEstrategiaOrdenacao(i int, okey string, fields RecordMap) (interface{}, interface{}) {
	val, ok := l.layout.GlobalSettings()["ordenar_escrita_por"]
	if !ok || val == "" {
		return i, i
	}

	if val == GlobalSettingsOrdenarEscritaPorSufixo {
		sorte1 := helper.ToSafeInt(strings.TrimPrefix(okey, "segmento_"))
		sorte2 := helper.ToSafeInt(strings.TrimPrefix(okey, "segmento_"))
		return sorte1, sorte2
	}

	if val == GlobalSettingsOrdenarEscritaPorNomeCampo {
		campo, _ := l.layout.GlobalSettings()["ordenar_escrita_usando_campo"]
		if fields[campo] != "" {
			sorte1 := fields[campo]
			sorte2 := fields[campo]
			return sorte1, sorte2
		}
	}

	if val == GlobalSettingsOrdenarEscritaPorSufixoECampo {
		campo, _ := l.layout.GlobalSettings()["ordenar_escrita_usando_campo"]
		if fields[campo] != "" {
			sorte1 := helper.ToSafeInt(strings.TrimPrefix(okey, "segmento_"))
			sorte2 := fields[campo]
			return sorte1, sorte2
		}
	}

	return i, i
}
