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

func TestRemessaAFD(t *testing.T) {
	source := strings.NewReader(mte.AFD)
	layout, err := superfile.NewLayout(source)

	remessa := superfile.NewRemessa(layout)
	remessa.Header["identificador_empregador"] = 1
	remessa.Header["cpf_cnpj_empregador"] = 94281944095
	remessa.Header["razao_social_empregador"] = "HELOO SOLUTIONS"
	remessa.Header["numero_rep"] = 1200
	remessa.Header["data_inicio"] = time.Now()
	remessa.Header["data_fim"] = time.Now()
	remessa.Header["data_geracao"] = time.Now()
	remessa.Header["hora_geracao"] = time.Now()

	lote := remessa.NovoLote()
	remessa.InserirLote(lote)

	for i := 0; i < 10; i++ {
		detalhe := lote.NovoDetalhe()
		detalhe["segmento_2"]["data_gravacao"] = time.Now()
		detalhe["segmento_2"]["hora_gravacao"] = time.Now()
		detalhe["segmento_2"]["identificador_empregador"] = 2
		detalhe["segmento_2"]["cpf_cnpj_empregador"] = 564281944095
		detalhe["segmento_2"]["cei_empregador"] = 100000000000
		detalhe["segmento_2"]["razao_social_ou_nome_empregador"] = "JOSE DA SILVA"
		detalhe["segmento_2"]["local_prestacao_servico"] = "RUA DOS SONHOS"

		detalhe["segmento_3"]["data_gravacao"] = time.Now()
		detalhe["segmento_3"]["hora_gravacao"] = time.Now()
		detalhe["segmento_3"]["numero_pis"] = 70746928042

		detalhe["segmento_4"]["data_ajuste"] = time.Now()
		detalhe["segmento_4"]["hora_ajuste"] = time.Now()
		detalhe["segmento_4"]["data_ajustada"] = time.Now()
		detalhe["segmento_4"]["hora_ajustada"] = time.Now()

		detalhe["segmento_5"]["data_gravacao"] = time.Now()
		detalhe["segmento_5"]["hora_gravacao"] = time.Now()
		detalhe["segmento_5"]["tipo_operacao"] = "I"
		detalhe["segmento_5"]["numero_pis"] = 70746928042
		detalhe["segmento_5"]["nome_empregado"] = "HELOO SOLUTIONS"
		lote.InserirDetalhe(detalhe)
	}

	remessa.Trailer["quantidade_registro_tipo_2"] = 1
	remessa.Trailer["quantidade_registro_tipo_3"] = 1
	remessa.Trailer["quantidade_registro_tipo_4"] = 1
	remessa.Trailer["quantidade_registro_tipo_5"] = 1
	remessaFile := file.NewRemessaFile(remessa, "afd.rem")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 42, arquivo.Name(), true)
}
