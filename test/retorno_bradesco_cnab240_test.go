package test

import (
	"os"
	"strings"
	"testing"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/layout/bradesco"
	"github.com/stretchr/testify/assert"
)

func TestRetornoBradescoCnab240Pagamento(t *testing.T) {
	source := strings.NewReader(bradesco.CNAB240Pagamentos)
	layout, err1 := superfile.NewLayout(source)

	f, _ := os.Open("fixtures/pagamento-bradesco.ret")
	defer f.Close()
	arquivo, err2 := superfile.NewRetornoFile(layout, f)
	retorno := arquivo.Read()

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, retorno)
	assert.Equal(t, "5SEVENSE SERVICOS DE REDEZINHA", retorno.Header["nome_empresa"])
	assert.Equal(t, 48, len(retorno.Header))
	assert.Equal(t, 16, len(retorno.Trailer))
	assert.Equal(t, 1, len(retorno.Lotes))
	assert.Equal(t, 56, len(retorno.Lotes[0].Header))
	assert.Equal(t, 10, len(retorno.Lotes[0].Segmentos()))
	assert.Equal(t, 20, len(retorno.Lotes[0].Trailer))
	assert.Equal(t, int64(1), retorno.TotalLotes())
	assert.Equal(t, 0, len(retorno.Falhas()))
}
