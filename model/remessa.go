package model

type Remessa struct {
	Header  map[string]interface{}
	Trailer map[string]interface{}
	Lotes   []*Lote
	layout  Layout
}

func NewRemessa(data Layout) *Remessa {
	return &Remessa{
		Header:  map[string]interface{}{},
		Trailer: map[string]interface{}{},
		layout:  data,
		Lotes:   []*Lote{},
	}
}

func (r *Remessa) InserirLote(l *Lote) {
	r.Lotes = append(r.Lotes, l)
}

func (r *Remessa) NovoLote(sequencial int64) *Lote {
	return &Lote{
		Header:        map[string]interface{}{},
		Trailer:       map[string]interface{}{},
		Sequencial:    sequencial,
		layout:        r.layout,
		detalhes:      []Detalhe{},
		SegmentoVazio: false,
	}
}

func (r *Remessa) NovoLoteSegmentoVazio(sequencial int64) *Lote {
	return &Lote{
		Header:        map[string]interface{}{},
		Trailer:       map[string]interface{}{},
		Sequencial:    sequencial,
		layout:        r.layout,
		detalhes:      []Detalhe{},
		SegmentoVazio: true,
	}
}

func (r *Remessa) GetRemessaLayout() map[interface{}]interface{} {
	return r.layout.GetRemessaLayout()
}

func (r *Remessa) TamanhoRegistro() int64 {
	return r.layout.GetTamanhoRegistro()
}

func (r *Remessa) TotalLotes() int64 {
	return int64(len(r.Lotes))
}

func (r *Remessa) TotalRegistros() int64 {
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
