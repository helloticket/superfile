package test

import (
	"os"
	"strings"
	"testing"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/layout/ccsitef"
	"github.com/stretchr/testify/assert"
)

func TestRetornoCCSitefExtrato(t *testing.T) {
	source := strings.NewReader(ccsitef.ExtratoEletronico)
	layout, err1 := superfile.NewLayout(source)

	f, _ := os.Open("fixtures/ccsitef-vendas.rem")
	defer f.Close()
	arquivo, err2 := superfile.NewRetornoFile(layout, f)
	retorno := arquivo.Read()

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, retorno)
	assert.Equal(t, 18, len(retorno.Header))
	assert.Equal(t, 6, len(retorno.Trailer))
	assert.Equal(t, 1, len(retorno.Lotes))
	assert.Equal(t, 8, len(retorno.Lotes[0].Header))
	assert.Equal(t, 3, len(retorno.Lotes[0].Segmentos()))
	assert.Equal(t, 8, len(retorno.Lotes[0].Trailer))
	assert.Equal(t, int64(1), retorno.TotalLotes())
	assert.Equal(t, 10, len(retorno.Falhas()))
}
