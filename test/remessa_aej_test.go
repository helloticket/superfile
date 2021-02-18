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

func TestRemessaAEJ(t *testing.T) {
	source := strings.NewReader(mte.AEJ)
	layout, err := superfile.NewLayout(source)

	remessa := superfile.NewRemessa(layout)
	remessa.Header["identificador_empregador"] = 1
	remessa.Header["cpf_cnpj_empregador"] = 94281944095
	remessa.Header["razao_social_empregador"] = "HELOO SOLUTIONS"
	remessa.Header["data_inicio"] = time.Now()
	remessa.Header["data_fim"] = time.Now()
	remessa.Header["data_geracao"] = time.Now()
	remessa.Header["hora_geracao"] = time.Now()

	lote := remessa.NovoLote()
	remessa.InserirLote(lote)

	for i := 0; i < 10; i++ {
		detalhes := lote.NovoDetalhe()
		detalhes["segmento_2"]["identificador_vinculo"] = 1
		detalhes["segmento_2"]["cpf_empregado"] = 00122344561
		detalhes["segmento_2"]["nome_empregado"] = "Phillipp Probst"

		detalhes["segmento_3"]["codigo_hor_contratual"] = "1"
		detalhes["segmento_3"]["duracao_jornada"] = 300
		detalhes["segmento_3"]["hora_entrada_1"] = time.Now()
		detalhes["segmento_3"]["hora_saida_1"] = time.Now()
		detalhes["segmento_3"]["hora_entrada_2"] = time.Now()
		detalhes["segmento_3"]["hora_saida_2"] = time.Now()

		detalhes["segmento_4"]["identificador_vinculo"] = 1
		detalhes["segmento_4"]["data_marcacao"] = time.Now()
		detalhes["segmento_4"]["hora_marcacao"] = time.Now()
		detalhes["segmento_4"]["tipo_marcacao"] = "E"
		detalhes["segmento_4"]["seq_entrada_saida"] = 1
		detalhes["segmento_4"]["fonte_marcacao"] = "O"
		detalhes["segmento_4"]["codigo_hor_contratual"] = "1"
		detalhes["segmento_4"]["motivo"] = ""

		detalhes["segmento_5"]["identificador_vinculo"] = 1
		detalhes["segmento_5"]["mat_esocial"] = "98765432112345"

		detalhes["segmento_6"]["nome_dev"] = "HELOO PONTO"
		detalhes["segmento_6"]["versao_ptrp"] = "1.0"
		detalhes["segmento_6"]["tipo_id_dev"] = 2
		detalhes["segmento_6"]["razao_social_dev"] = "HELOO SOLUTIONS"
		detalhes["segmento_6"]["email_dev"] = "comercial@heloo.com.br"
		lote.InserirDetalhe(detalhes)
	}

	remessa.Trailer["quantidade_registro_tipo_2"] = 10
	remessa.Trailer["quantidade_registro_tipo_3"] = 10
	remessa.Trailer["quantidade_registro_tipo_4"] = 10
	remessa.Trailer["quantidade_registro_tipo_5"] = 10
	remessa.Trailer["quantidade_registro_tipo_6"] = 10
	remessaFile := file.NewRemessaFile(remessa, "aej.rem")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 52, arquivo.Name(), true)
}
