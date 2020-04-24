package test

import (
	"os"
	"strings"
	"testing"

	"github.com/helderfarias/superfile"
	"github.com/helderfarias/superfile/layout/ccsitef"
	"github.com/stretchr/testify/assert"
)

func TestRetornoCCSitefExtrato(t *testing.T) {
	source := strings.NewReader(ccsitef.ExtratoEletronico)
	layout, err := superfile.NewLayout(source)

	f, _ := os.Open("fixtures/ccsitef-vendas.rem")
	defer f.Close()
	arquivo, err := superfile.NewRetornoFile(layout, f)
	retorno := arquivo.Read()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, retorno)
	assert.Equal(t, 9, len(retorno.Header))
	assert.Equal(t, 3, len(retorno.Trailer))
	assert.Equal(t, 1, len(retorno.Lotes))
	assert.Equal(t, 4, len(retorno.Lotes[0].Header))
	assert.Equal(t, 3, len(retorno.Lotes[0].Segmentos()))
	assert.Equal(t, 4, len(retorno.Lotes[0].Trailer))
	assert.Equal(t, int64(1), retorno.TotalLotes())
}
