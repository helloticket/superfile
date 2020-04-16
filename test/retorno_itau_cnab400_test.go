package test

import (
	"os"
	"strings"
	"testing"

	"github.com/helderfarias/cnab-go"
	"github.com/helderfarias/cnab-go/file"
	"github.com/helderfarias/cnab-go/layout/itau"
	"github.com/stretchr/testify/assert"
)

func TestRetornoItauCnab400Cobranca(t *testing.T) {
	source := strings.NewReader(itau.CNAB400Cobranca)
	layout, err := cnab.NewLayout("400", source)
	assert.Nil(t, err)

	f, err := os.Open("fixtures/cobranca-itau-cnab400.ret")
	assert.Nil(t, err)
	defer f.Close()

	arquivo, err := file.NewRetornoFile(layout, f)
	assert.Nil(t, err)

	retorno := arquivo.Read()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, retorno)
	assert.Equal(t, 20, len(retorno.Header))
	assert.Equal(t, 21, len(retorno.Trailer))
	assert.Equal(t, 4, len(retorno.Segmentos()))
}
