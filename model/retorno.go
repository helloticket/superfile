package model

import "fmt"

type Retorno struct {
	Header  map[string]interface{}
	Trailer map[string]interface{}
	Lotes   []*Lote
	layout  Layout
	falhas  []FalhaParseDecode
}

func NewRetorno(data Layout) *Retorno {
	return &Retorno{
		Header:  map[string]interface{}{},
		Trailer: map[string]interface{}{},
		falhas:  []FalhaParseDecode{},
		layout:  data,
		Lotes:   []*Lote{},
	}
}

func (r *Retorno) InserirLote(l *Lote) {
	r.Lotes = append(r.Lotes, l)
}

func (r *Retorno) NovoLote(sequencial int64) *Lote {
	return &Lote{
		Header:        map[string]interface{}{},
		Trailer:       map[string]interface{}{},
		Sequencial:    sequencial,
		layout:        r.layout,
		detalhes:      []Detalhe{},
		SegmentoVazio: false,
	}
}

func (r *Retorno) GetRetornoLayout() map[interface{}]interface{} {
	return r.layout.GetRetornoLayout()
}

func (r *Retorno) TamanhoRegistro() int64 {
	return r.layout.GetTamanhoRegistro()
}

func (r *Retorno) TotalLotes() int64 {
	return int64(len(r.Lotes))
}

func (r *Retorno) TotalRegistros() int64 {
	totalHeaderLote := 1
	totalSegmentosLote := 0
	totalTrailerLote := 1

	for _, l := range r.Lotes {
		for _, d := range l.detalhes {
			for range d {
				totalSegmentosLote++
			}
		}
	}

	return int64(totalHeaderLote + totalSegmentosLote + totalTrailerLote)
}

func (l *Retorno) Segmentos() []Segmento {
	segmentos := []Segmento{}

	for _, l := range l.Lotes {
		segmentos = append(segmentos, l.Segmentos()...)
	}

	return segmentos
}

func (l *Retorno) GetSegmentoErro(fieldName string) error {
	for _, l := range l.Segmentos() {
		for k, v := range l.Valores {
			if k == fieldName {
				e, _ := v.(error)
				return e
			}
		}
	}
	return nil
}

func (l *Retorno) RegistrarFalha(posicao int64, err error) {
	if err == nil {
		return
	}

	l.falhas = append(l.falhas, FalhaParseDecode{PosicaoNoArquivo: posicao, Erro: fmt.Errorf("Falha na linha %v. Datalhe: %v", posicao, err)})
}

func (l *Retorno) Falhas() []FalhaParseDecode {
	return l.falhas
}
