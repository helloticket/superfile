package test

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/layout/fortes"
	"github.com/stretchr/testify/assert"
)

func TestRetornoFortes(t *testing.T) {
	source := strings.NewReader(fortes.ImportacaoMovimento)
	layout, err := superfile.NewLayout(source)
	assert.Nil(t, err)

	f, err := os.Open("fixtures/fortes.ps")
	assert.Nil(t, err)
	defer f.Close()

	arquivo, err := superfile.NewRetornoFile(layout, f)
	assert.Nil(t, err)

	retorno := arquivo.Read()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, retorno)
	assert.Equal(t, 12, len(retorno.Header))
	assert.Equal(t, 10, len(retorno.Trailer))
	assert.Equal(t, 7, len(retorno.Segmentos()))
	assert.Equal(t, 4, len(retorno.Falhas()))

	log.Println(retorno.Falhas())
}
