package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObterSegmentosOrdenadoPorNumeroSufixo(t *testing.T) {
	layoutMock := &mockLayout{}
	lote := Lote{layout: layoutMock}

	layoutMock.On("GetSegmentoDefinitions").Return(RecordDetailMap{
		"segmento_3": {"sequencial": 0},
		"segmento_1": {"sequencial": 0},
		"segmento_2": {"sequencial": 0},
	})
	layoutMock.On("GlobalSettings").Return(map[string]string{"ordenar_escrita_por": "sufixo_segmento"})

	s2 := lote.NovoDetalhe()
	s2["segmento_2"]["sequencial"] = 2
	lote.InserirDetalhe(s2)

	s1 := lote.NovoDetalhe()
	s1["segmento_1"]["sequencial"] = 1
	lote.InserirDetalhe(s1)

	s3 := lote.NovoDetalhe()
	s3["segmento_3"]["sequencial"] = 3
	lote.InserirDetalhe(s3)

	assert.Equal(t, []Segmento{
		{Nome: "segmento_1", Valores: RecordMap{"sequencial": 1}},
		{Nome: "segmento_1", Valores: RecordMap{"sequencial": 1}},
		{Nome: "segmento_1", Valores: RecordMap{"sequencial": 1}},
		{Nome: "segmento_2", Valores: RecordMap{"sequencial": 2}},
		{Nome: "segmento_2", Valores: RecordMap{"sequencial": 2}},
		{Nome: "segmento_2", Valores: RecordMap{"sequencial": 2}},
		{Nome: "segmento_3", Valores: RecordMap{"sequencial": 3}},
		{Nome: "segmento_3", Valores: RecordMap{"sequencial": 3}},
		{Nome: "segmento_3", Valores: RecordMap{"sequencial": 3}},
	},
		lote.Segmentos())
}

func TestObterSegmentosOrdenadoPorSufixoECampo(t *testing.T) {
	layoutMock := &mockLayout{}
	lote := Lote{layout: layoutMock}

	layoutMock.On("GetSegmentoDefinitions").Return(RecordDetailMap{
		"segmento_3": {"sequencial": 0, "ordem": 0},
		"segmento_1": {"sequencial": 0, "ordem": 0},
		"segmento_2": {"sequencial": 0, "ordem": 0},
	})
	layoutMock.On("GlobalSettings").Return(map[string]string{
		"ordenar_escrita_por":          "sufixo_e_campo_segmento",
		"ordenar_escrita_usando_campo": "ordem",
	})

	s2 := lote.NovoDetalhe()
	s2["segmento_2"]["sequencial"] = 2
	s2["segmento_2"]["ordem"] = 2
	lote.InserirDetalhe(s2)

	s3 := lote.NovoDetalhe()
	s3["segmento_3"]["sequencial"] = 3
	s2["segmento_3"]["ordem"] = 3
	lote.InserirDetalhe(s3)

	s1 := lote.NovoDetalhe()
	s1["segmento_1"]["sequencial"] = 1
	s2["segmento_1"]["ordem"] = 1
	lote.InserirDetalhe(s1)

	assert.Equal(t, []Segmento{
		{Nome: "segmento_1", Valores: RecordMap{"ordem": 1, "sequencial": 1}},
		{Nome: "segmento_1", Valores: RecordMap{"ordem": 1, "sequencial": 1}},
		{Nome: "segmento_1", Valores: RecordMap{"ordem": 1, "sequencial": 1}},
		{Nome: "segmento_2", Valores: RecordMap{"ordem": 2, "sequencial": 2}},
		{Nome: "segmento_2", Valores: RecordMap{"ordem": 2, "sequencial": 2}},
		{Nome: "segmento_2", Valores: RecordMap{"ordem": 2, "sequencial": 2}},
		{Nome: "segmento_3", Valores: RecordMap{"ordem": 3, "sequencial": 3}},
		{Nome: "segmento_3", Valores: RecordMap{"ordem": 3, "sequencial": 3}},
		{Nome: "segmento_3", Valores: RecordMap{"ordem": 3, "sequencial": 3}},
	},
		lote.Segmentos())
}

func TestObterSegmentosOrdenadoPorOutroCampo(t *testing.T) {
	layoutMock := &mockLayout{}
	lote := Lote{layout: layoutMock}

	layoutMock.On("GetSegmentoDefinitions").Return(RecordDetailMap{
		"segmento_3": {"sequencial": 0, "ordem": 0},
		"segmento_1": {"sequencial": 0, "ordem": 0},
		"segmento_2": {"sequencial": 0, "ordem": 0},
	})
	layoutMock.On("GlobalSettings").Return(map[string]string{
		"ordenar_escrita_por":          "campo_segmento",
		"ordenar_escrita_usando_campo": "sequencial",
	})

	s2 := lote.NovoDetalhe()
	s2["segmento_2"]["sequencial"] = 2
	s2["segmento_2"]["ordem"] = 1
	lote.InserirDetalhe(s2)

	s1 := lote.NovoDetalhe()
	s1["segmento_1"]["sequencial"] = 1
	s2["segmento_1"]["ordem"] = 2
	lote.InserirDetalhe(s1)

	s3 := lote.NovoDetalhe()
	s3["segmento_3"]["sequencial"] = 3
	s2["segmento_3"]["ordem"] = 3
	lote.InserirDetalhe(s3)

	assert.Equal(t, []Segmento{
		{Nome: "segmento_1", Valores: RecordMap{"ordem": 2, "sequencial": 1}},
		{Nome: "segmento_1", Valores: RecordMap{"ordem": 2, "sequencial": 1}},
		{Nome: "segmento_1", Valores: RecordMap{"ordem": 2, "sequencial": 1}},
		{Nome: "segmento_2", Valores: RecordMap{"ordem": 1, "sequencial": 2}},
		{Nome: "segmento_2", Valores: RecordMap{"ordem": 1, "sequencial": 2}},
		{Nome: "segmento_2", Valores: RecordMap{"ordem": 1, "sequencial": 2}},
		{Nome: "segmento_3", Valores: RecordMap{"ordem": 3, "sequencial": 3}},
		{Nome: "segmento_3", Valores: RecordMap{"ordem": 3, "sequencial": 3}},
		{Nome: "segmento_3", Valores: RecordMap{"ordem": 3, "sequencial": 3}},
	},
		lote.Segmentos())
}

func TestObterSegmentos(t *testing.T) {
	layoutMock := &mockLayout{}
	lote := Lote{layout: layoutMock}

	layoutMock.On("GetSegmentoDefinitions").Return(RecordDetailMap{
		"segmento_1": {"sequencial": 0, "ordem": 0},
		"segmento_2": {"sequencial": 0, "ordem": 0},
	})
	layoutMock.On("GlobalSettings").Return(map[string]string{})

	s2 := lote.NovoDetalhe()
	s2["segmento_2"]["sequencial"] = 2
	s2["segmento_2"]["ordem"] = 1
	lote.InserirDetalhe(s2)

	s1 := lote.NovoDetalhe()
	s1["segmento_1"]["sequencial"] = 1
	s2["segmento_1"]["ordem"] = 2
	lote.InserirDetalhe(s1)

	assert.Len(t, lote.Segmentos(), 4)
}
