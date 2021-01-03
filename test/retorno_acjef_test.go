package test

import (
	"os"
	"strings"
	"testing"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/layout/mte"
	"github.com/stretchr/testify/assert"
)

func TestRetornoACJEFValido(t *testing.T) {
	source := strings.NewReader(mte.ACJEF)
	layout, err := superfile.NewLayout(source)
	assert.Nil(t, err)

	f, err := os.Open("fixtures/acjef.ret")
	assert.Nil(t, err)
	defer f.Close()

	arquivo, err := superfile.NewRetornoFile(layout, f)
	assert.Nil(t, err)

	retorno := arquivo.Read()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, retorno)
	assert.Equal(t, 20, len(retorno.Header))
	assert.Equal(t, 4, len(retorno.Trailer))
	assert.Equal(t, 2, len(retorno.Segmentos()))
}

func TestRetornoACJEFComFalhaEmCampo(t *testing.T) {
	source := strings.NewReader(mte.ACJEF)
	layout, err := superfile.NewLayout(source)
	assert.Nil(t, err)

	f, err := os.Open("fixtures/acjef_erro_campo.ret")
	assert.Nil(t, err)
	defer f.Close()

	arquivo, err := superfile.NewRetornoFile(layout, f)
	assert.Nil(t, err)

	retorno := arquivo.Read()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, retorno)
	assert.Equal(t, 20, len(retorno.Header))
	assert.Equal(t, 4, len(retorno.Trailer))
	assert.Equal(t, 2, len(retorno.Segmentos()))
	assert.EqualError(t, retorno.GetSegmentoErro("horas_extras_1_error"), "Parser error: [Block: detalhes.segmento_3, Name: horas_extras_1, Pos: (47,50), Picture: 9(4), Format: hhmm, Default: <nil>, Value: 00000000332000000000002507202000120001000100021A4300120164300120164300120164300121001210000] - parsing time \"1A43\" as \"1504\": cannot parse \"A43\" as \"04\"")
}

func TestRetornoACJEFComFalhasGerais(t *testing.T) {
	source := strings.NewReader(mte.ACJEF)
	layout, err := superfile.NewLayout(source)
	assert.Nil(t, err)

	f, err := os.Open("fixtures/acjef_falhas.ret")
	assert.Nil(t, err)
	defer f.Close()

	arquivo, err := superfile.NewRetornoFile(layout, f)
	assert.Nil(t, err)

	retorno := arquivo.Read()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, retorno)
	assert.Equal(t, 0, len(retorno.Header))
	assert.Equal(t, 0, len(retorno.Trailer))
	assert.Equal(t, 0, len(retorno.Segmentos()))
	assert.Equal(t, 8, len(retorno.Falhas()))
}
