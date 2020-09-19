package test

import (
	"strings"
	"testing"
	"time"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/file"
	"github.com/helloticket/superfile/layout/mte"
	"github.com/stretchr/testify/assert"
)

func TestRemessaAFDT(t *testing.T) {
	source := strings.NewReader(mte.AFDT)
	layout, err := superfile.NewLayout(source)

	remessa := superfile.NewRemessa(layout)
	remessa.Header["seq_registro"] = 1
	remessa.Header["identificador_empregador"] = 1
	remessa.Header["cpf_cnpj_empregador"] = 94281944095
	remessa.Header["cei_empregador"] = 200000000000
	remessa.Header["razao_social_empregador"] = "1"
	remessa.Header["data_inicio"] = time.Now()
	remessa.Header["data_fim"] = time.Now()
	remessa.Header["data_geracao"] = time.Now()
	remessa.Header["hora_geracao"] = time.Now()

	lote := remessa.NovoLote()
	remessa.InserirLote(lote)

	for i := 1; i <= 20; i++ {
		detalhe := lote.NovoDetalhe()
		detalhe["segmento_2"]["seq_registro"] = i
		detalhe["segmento_2"]["data_marcacao_ponto"] = time.Now()
		detalhe["segmento_2"]["hora_marcacao_ponto"] = time.Now()
		detalhe["segmento_2"]["numero_pis"] = 70746928042
		detalhe["segmento_2"]["numero_fabricacao_rep"] = 1
		detalhe["segmento_2"]["tipo_marcacao"] = "1"
		detalhe["segmento_2"]["seq_empregado_jornada_entrada_saida"] = 1
		detalhe["segmento_2"]["tipo_registro_eletro_digi_pre_assina"] = "2"
		detalhe["segmento_2"]["motivo"] = "1"
		lote.InserirDetalhe(detalhe)
	}

	remessa.Trailer["seq_registro"] = 1
	remessaFile := file.NewRemessaFile(remessa, "afdt.rem")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 22, arquivo.Name(), true)
}
