package test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/layout/febraban"
)

func TestRetornoFebrabanCnab240Cobranca(t *testing.T) {
	source := strings.NewReader(febraban.CNAB240Pagamentos)
	layout, err1 := superfile.NewLayout(source)

	f, _ := os.Open("fixtures/pagamento-febraban.ret")
	defer f.Close()
	arquivo, err2 := superfile.NewRetornoFile(layout, f)
	retorno := arquivo.Read()

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, retorno)
	assert.Equal(t, int64(1), retorno.TotalLotes())

	lote := retorno.Lotes[0]
	assert.Equal(t, int64(123), lote.Header["lote_servico"])
	assert.Equal(t, int64(12234567000186), lote.Header["numero_inscricao_empresa"])
	assert.Equal(t, "NOME EMPRESA LOTE", lote.Header["nome_empresa"])

	segmentos := retorno.Segmentos()
	assert.Len(t, segmentos, 3)
	assert.Equal(t, int64(123), segmentos[0].Valores["lote_servico"])
	assert.Equal(t, int64(123), segmentos[1].Valores["lote_servico"])
	assert.Equal(t, int64(123), segmentos[2].Valores["lote_servico"])
	assert.Equal(t, 0, len(retorno.Falhas()))
}
