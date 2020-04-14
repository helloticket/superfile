package model

type Retorno struct {
	Header  map[string]interface{}
	Trailer map[string]interface{}
	Lotes   []*Lote
	layout  Layout
}

func NewRetorno(data Layout) *Retorno {
	return &Retorno{
		Header:  map[string]interface{}{},
		Trailer: map[string]interface{}{},
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
