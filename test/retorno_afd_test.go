package test

import (
	"os"
	"strings"
	"testing"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/layout/mte"
	"github.com/stretchr/testify/assert"
)

func TestRetornoAFD(t *testing.T) {
	source := strings.NewReader(mte.AFD)
	layout, err := superfile.NewLayout(source)
	assert.Nil(t, err)

	f, err := os.Open("fixtures/afd.ret")
	assert.Nil(t, err)
	defer f.Close()

	arquivo, err := superfile.NewRetornoFile(layout, f)
	assert.Nil(t, err)

	retorno := arquivo.Read()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, retorno)
	assert.Equal(t, 22, len(retorno.Header))
	assert.Equal(t, 12, len(retorno.Trailer))
	assert.Equal(t, 40, len(retorno.Segmentos()))
	assert.Equal(t, 0, len(retorno.Falhas()))
}
