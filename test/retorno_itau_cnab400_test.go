package test

import (
	"os"
	"strings"
	"testing"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/layout/itau"
	"github.com/stretchr/testify/assert"
)

func TestRetornoItauCnab400Cobranca(t *testing.T) {
	source := strings.NewReader(itau.CNAB400Cobranca)
	layout, err := superfile.NewLayout(source)
	assert.Nil(t, err)

	f, err := os.Open("fixtures/cobranca-itau-cnab400.ret")
	assert.Nil(t, err)
	defer f.Close()

	arquivo, err := superfile.NewRetornoFile(layout, f)
	assert.Nil(t, err)

	retorno := arquivo.Read()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, retorno)
	assert.Equal(t, 40, len(retorno.Header))
	assert.Equal(t, 42, len(retorno.Trailer))
	assert.Equal(t, 4, len(retorno.Segmentos()))
	assert.Equal(t, 0, len(retorno.Falhas()))
}
